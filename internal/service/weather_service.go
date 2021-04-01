package service

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
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
	rainDesc := model.WeatherModel.GetCurRain(params.CityCode)

	if err == nil {
		err1 := json.Unmarshal([]byte(res), &data)
		if err1 == nil {
			day := currentTime.Day()
			reDay := time.Unix(data.Date, 0).Day()
			// 最近时间和当前时间不符时通知更新缓存
			if reDay != day {
				go sendWeatherChan(params)
				// 清空数据，重新同步读取
				data = new(weather_mgr.TodayResp)
			}
		}
	}

	// 真实同步读取
	if data.Date == 0 {
		resp, err := model.HefengModel.GetFormatData(params)
		if err != nil {
			return nil, err

		}
		data = resp.Realtime

	}
	// 获取湿度
	humidity := model.WeatherModel.GetHumidityDay(params.CityCode, currentTime)
	if humidity == "" {
		data.Humidity = "20%"
	} else {
		data.Humidity = humidity + "%"
	}

	// 当缓存数据不存在 取温度作为穿衣提醒文案
	temp := model.WeatherModel.GetLowDayTemp(params.CityCode, currentTime)
	var walkRemind string
	data.WarmRemind, walkRemind, data.Comfort = helper.GetWeatherNotices(temp, data.WarmRemind)

	if data.WalkRemind == "" {
		data.WalkRemind = walkRemind
	}
	data.RainDesc = rainDesc
	var v *model.Daily
	// 获取出行数据
	walkOut := model.WeatherModel.GetWalkOut(params.CityCode)
	err = json.Unmarshal([]byte(walkOut), &v)
	if v != nil && err == nil {
		data.WindDirDay = v.WindDirDay + v.WindScaleDay + "级"
	}
	return data, nil

}
func (w *WeatherService) WarningList(ctx context.Context, params *weather_mgr.WeatherReq) (data *weather_mgr.WarningListResp, err error) {
	params.WeatherType = "warning"
	data = new(weather_mgr.WarningListResp)
	res, err := model.WeatherModel.GetWarningList(fmt.Sprintf("%s", params.CityCode))

	if err == nil {
		err1 := json.Unmarshal([]byte(res), &data.List)
		if err1 == nil {
			if len(data.List) > 0 {
				return data, nil
			}
		}
	}
	resp, err := model.HefengModel.GetFormatData(params)
	if err != nil {
		return nil, err
	}
	return resp.WarningList, nil
}

func sendWeatherChan(params *weather_mgr.WeatherReq) {
	model.WeatherChan <- helper.WeatherChanData{params}

}
