package service

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"time"
	"weather_mgr/cootek/pgd/weather_mgr"
	"weather_mgr/helper"
	"weather_mgr/internal/model"
	"weather_mgr/pkg/xzap"
)

type WeatherService struct {
}

// 24小时 天气情况
func (w *WeatherService) Hourly(ctx context.Context, params *weather_mgr.WeatherReq) (*weather_mgr.HourlyResp, error) {
	resp := new(weather_mgr.HourlyResp)
	resp.List = []*weather_mgr.HourlyStyle{}
	params.WeatherType = "hourly"

	res, err := model.WeatherModel.GetWeatherHourlyData(fmt.Sprintf("%s", params.CityCode))
	currentTime, _ := model.MsIntToTime(params.SessionBase.SendTsMillisec)
	// temperature := ""
	sign := 0

	if err == nil {
		err1 := json.Unmarshal([]byte(res), &resp.List)
		if err1 == nil {
			// 防止缓存数据过老，循环截去老数据
			for _, v := range resp.List {
				if v.Date < currentTime.Unix()-3600 {
					sign++
					continue
				}
			}
			// 防止超长
			if sign > 0 {
				// 防止超出切片长度
				if sign == len(resp.List) {
					resp.List = resp.List[sign-1:]
				} else {
					resp.List = resp.List[sign:]
				}
			}

			// temperature = resp.List[0].Temperature
			reHour := time.Unix(resp.List[0].Date, 0).Hour()
			// 获取上一时数据 且与当前数据不覆盖前提合并到 data 中
			lastHour := model.WeatherModel.GetLastHourlyData(ctx, params.CityCode, currentTime)
			if len(lastHour) != 0 && lastHour[0].Date != 0 {
				lHour := time.Unix(lastHour[0].Date, 0).Hour()
				if lHour != reHour {
					xzap.Debug("且与当前数据不覆盖前提合并到", zap.Any("lastHour", lastHour), zap.Any("data", resp.List))
					resp.List = append(lastHour, resp.List...)
				}
			}

		}
	}
	if len(resp.List) < 15 {
		res, err := model.HefengModel.GetFormatData(params)
		resp := new(weather_mgr.HourlyResp)
		if err == nil && len(res.Hourly) > 0 {
			sign++
			resp.List = res.Hourly
		}
	}
	// 最近时间和当前时间不符时通知更新缓存
	if sign > 0 {
		go sendWeatherChan(params)
	}
	// // save  提供给温差奖励用
	// temp, err1 := strconv.ParseFloat(temperature, 64)
	// if err1 == nil {
	// 	go model.WeatherModel.SetNowTemp(params.Uid, temp, currentTime)
	// }

	return resp, nil
}

// 近一周天气情况
func (w *WeatherService) Daily(ctx context.Context, params *weather_mgr.WeatherReq) (*weather_mgr.DailyResp, error) {
	resp := new(weather_mgr.DailyResp)
	resp.List = []*weather_mgr.DailyStyle{}
	params.WeatherType = "daily"
	res, err := model.WeatherModel.GetWeatherDailyData(fmt.Sprintf("%s", params.CityCode))
	sign := 0

	if err == nil {
		err1 := json.Unmarshal([]byte(res), &resp.List)
		if err1 == nil {
			currentTime, _ := model.MsIntToTime(params.SessionBase.SendTsMillisec)
			// 防止缓存数据过老，循环截去老数据
			for _, v := range resp.List {
				if v.Date < currentTime.Unix()-86400 {
					sign++
					continue
				}
			}
			// 最近时间和当前时间不符时通知更新缓存
			reDay := time.Unix(resp.List[0].Date, 0).Day()

			// 防止超长
			if sign > 0 {
				resp.List = resp.List[sign-1:]
			}

			// 获取上一天数据 且与当前数据不覆盖前提合并到 data 中
			lastDaily := model.WeatherModel.GetLastDayData(ctx, params.CityCode, currentTime)
			if len(lastDaily) != 0 && lastDaily[0].Date != 0 {
				lDay := time.Unix(lastDaily[0].Date, 0).Day()
				if lDay != reDay {
					resp.List = append(lastDaily, resp.List...)
				}
			}
		}
	}

	if len(resp.List) < 7 {
		res, err := model.HefengModel.GetFormatData(params)
		if err == nil {
			resp.List = res.Daily
			sign++

		}
	}
	if sign > 0 {
		go sendWeatherChan(params)
	}
	return resp, nil
}

// 当天生活指数情况
func (w *WeatherService) Today(ctx context.Context, params *weather_mgr.WeatherReq) (data *weather_mgr.TodayResp, err error) {
	params.WeatherType = "today"
	data = new(weather_mgr.TodayResp)
	res, err := model.WeatherModel.GetWeatherRealTimeData(fmt.Sprintf("%s", params.CityCode))
	currentTime, _ := model.MsIntToTime(params.SessionBase.SendTsMillisec)
	var rainDesc string

	if err == nil {
		err1 := json.Unmarshal([]byte(res), &data)
		if err1 == nil {
			day := currentTime.Day()
			reDay := time.Unix(data.Date, 0).Day()
			// 最近时间和当前时间不符时通知更新缓存
			if reDay != day {
				go sendWeatherChan(params)
			}
			if data.WarmRemind != "" {
				level, _ := strconv.ParseInt(data.WarmRemind, 10, 64)
				if level > 0 {
					data.WarmRemind = helper.GetWarmRemindNotice(999, data.WarmRemind)
				} else {
					// 当缓存数据不存在 取温度作为穿衣提醒标识
					temp := model.WeatherModel.GetLowDayTemp(params.CityCode, currentTime)
					data.WarmRemind = helper.GetWarmRemindNotice(temp, "")
				}
			}
			if data.Humidity == "" {
				humidity := model.WeatherModel.GetHumidityDay(params.CityCode, currentTime)
				if humidity == "" {
					data.Humidity = "20%"
				} else {
					data.Humidity = humidity + "%"
				}
			}
			// 穿衣：如无缓存，显示“舒适”
			// 感冒：如无缓存，显示“较易发”
			// 紫外线：如无缓存，显示“中等”
			// 运动：如无缓存，显示“较不宜”
			// 钓鱼：如无缓存，显示“一般”
			// 晾晒：如无缓存，显示“一般”
			if data.LifeSuggestion.Dressing == "" {
				data.LifeSuggestion.Dressing = "舒适"
			}

			if data.LifeSuggestion.Fishing == "" {
				data.LifeSuggestion.Fishing = "一般"
			}
			if data.LifeSuggestion.Flu == "" {
				data.LifeSuggestion.Flu = "较易发"
			}
			if data.LifeSuggestion.Sport == "" {
				data.LifeSuggestion.Sport = "较不宜"
			}
			if data.LifeSuggestion.Uv == "" {
				data.LifeSuggestion.Uv = "中等"
			}
			if data.LifeSuggestion.Airing == "" {
				data.LifeSuggestion.Airing = "一般"
			}
			data.RainDesc = rainDesc
			return data, nil
		}
	}
	resp, err := model.HefengModel.GetFormatData(params)
	if err != nil {
		return nil, err

	}
	if params.TestGroup == "1" {
		// 分钟级降水文案
		resp.Realtime.RainDesc = rainDesc
	}

	if resp.Realtime.WarmRemind != "" {
		warmRemind := helper.GetWarmRemindNotice(999, resp.Realtime.WarmRemind)
		resp.Realtime.WarmRemind = warmRemind
	}
	if resp.Realtime.Humidity == "" {
		humidity := model.WeatherModel.GetHumidityDay(params.CityCode, currentTime)
		if humidity == "" {
			resp.Realtime.Humidity = "20%"
		} else {
			resp.Realtime.Humidity = humidity + "%"
		}
	}

	// 穿衣：如无缓存，显示“舒适”
	// 感冒：如无缓存，显示“较易发”
	// 紫外线：如无缓存，显示“中等”
	// 运动：如无缓存，显示“较不宜”
	// 钓鱼：如无缓存，显示“一般”
	// 晾晒：如无缓存，显示“一般”
	if resp.Realtime.LifeSuggestion.Dressing == "" {
		resp.Realtime.LifeSuggestion.Dressing = "舒适"
	}

	if resp.Realtime.LifeSuggestion.Fishing == "" {
		resp.Realtime.LifeSuggestion.Fishing = "一般"
	}
	if resp.Realtime.LifeSuggestion.Flu == "" {
		resp.Realtime.LifeSuggestion.Flu = "较易发"
	}
	if resp.Realtime.LifeSuggestion.Sport == "" {
		resp.Realtime.LifeSuggestion.Sport = "较不宜"
	}
	if resp.Realtime.LifeSuggestion.Uv == "" {
		resp.Realtime.LifeSuggestion.Uv = "中等"
	}
	if resp.Realtime.LifeSuggestion.Airing == "" {
		resp.Realtime.LifeSuggestion.Airing = "一般"
	}

	return resp.Realtime, nil

}

func sendWeatherChan(params *weather_mgr.WeatherReq) {
	model.WeatherChan <- helper.WeatherChanData{params}

}
