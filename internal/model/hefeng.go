package model

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
	"weather_mgr/cootek/pgd/weather_mgr"
	"weather_mgr/helper"
	"weather_mgr/internal"
	"weather_mgr/pkg/xzap"
)

type hefengModel struct {
}

var HefengModel *hefengModel
var WeatherChan chan helper.WeatherChanData

func init() {
	HefengModel = new(hefengModel)
	WeatherChan = make(chan helper.WeatherChanData, 1)
	go HefengModel.AsyncWeatherData()
}

type HefengData struct {
	Code    string `json:"code"`
	Summary string `json:"summary"` // 降雨文案
	Now     struct {
		ObsTime   string `json:"obsTime"`
		Temp      string `json:"temp"`
		FeelsLike string `json:"feelsLike"`
		Icon      string `json:"icon"`
		Text      string `json:"text"`
		Wind360   string `json:"wind360"`
		WindDir   string `json:"windDir"`
		WindScale string `json:"windScale"`
		WindSpeed string `json:"windSpeed"`
		Humidity  string `json:"humidity"`
		Precip    string `json:"precip"`
		Pressure  string `json:"pressure"`
		Vis       string `json:"vis"`
		Cloud     string `json:"cloud"`
		Dew       string `json:"dew"`
	} `json:"now"`
	Daily []struct {
		FxDate         string `json:"fxDate"`
		Sunrise        string `json:"sunrise"`
		Sunset         string `json:"sunset"`
		Moonrise       string `json:"moonrise"`
		Moonset        string `json:"moonset"`
		MoonPhase      string `json:"moonPhase"`
		TempMax        string `json:"tempMax"`
		TempMin        string `json:"tempMin"`
		IconDay        string `json:"iconDay"`
		TextDay        string `json:"textDay"`
		IconNight      string `json:"iconNight"`
		TextNight      string `json:"textNight"`
		Wind360Day     string `json:"wind360Day"`
		WindDirDay     string `json:"windDirDay"`
		WindScaleDay   string `json:"windScaleDay"`
		WindSpeedDay   string `json:"windSpeedDay"`
		Wind360Night   string `json:"wind360Night"`
		WindDirNight   string `json:"windDirNight"`
		WindScaleNight string `json:"windScaleNight"`
		WindSpeedNight string `json:"windSpeedNight"`
		Humidity       string `json:"humidity"`
		Precip         string `json:"precip"`
		Pressure       string `json:"pressure"`
		Vis            string `json:"vis"`
		Cloud          string `json:"cloud"`
		UvIndex        string `json:"uvIndex"`
		Aqi            string `json:"aqi"`
		Text           string `json:"text"`     // 生活指数预报的详细描述，可能为空
		Level          string `json:"level"`    // 指数预报的等级
		Type           string `json:"type"`     // 指数预报的类型
		Category       string `json:"category"` // 指数类别
	} `json:"daily"`
	Hourly []struct {
		FxTime string `json:"fxTime"`
		Temp   string `json:"temp"`
		Icon   string `json:"icon"`
		Text   string `json:"text"`
		// Wind360   string `json:"wind360"`
		// WindDir   string `json:"windDir"`
		// WindScale string `json:"windScale"`
		// WindSpeed string `json:"windSpeed"`
		// Humidity  string `json:"humidity"`
		// Pop       string `json:"pop"`
		// Precip    string `json:"precip"`
		// Pressure  string `json:"pressure"`
		// Cloud     string `json:"cloud"`
		// Dew       string `json:"dew"`
	} `json:"hourly"`
	Warning []struct {
		// ID        string `json:"id"`
		// Sender    string `json:"sender"`
		// PubTime   string `json:"pubTime"`
		// Title     string `json:"title"`
		// StartTime string `json:"startTime"`
		// EndTime   string `json:"endTime"`
		// Status    string `json:"status"`
		// Level     string `json:"level"`
		// Type      string `json:"type"`
		// TypeName  string `json:"typeName"`
		Text string `json:"text"` // 预警详细文字描述
		// Related   string `json:"related"`
	} `json:"warning"`
}

// https://api.qweather.com/v7/astronomy/sunmoon  日落日出
// https://api.qweather.com/v7/air/now  空气质量
// https://api.qweather.com/v7/indices/1d  生活指数

// https://api.qweather.com/v7/warning/now  灾害预警
// https://api.qweather.com/v7/weather/24h  24 小时
// https://api.qweather.com/v7/weather/7d  近 7 天
// https://api.qweather.com/v7/weather/now   当前温度信息

func (M *hefengModel) GetFormatData(params *weather_mgr.WeatherReq) (val *helper.ParamWeatherInfoResp, err error) {
	val = new(helper.ParamWeatherInfoResp)
	// currentTime := params.Time
	currentTime := time.Now()
	switch params.WeatherType {
	case "today":
		normalData, err := M.GetApiData(params, "/v7/indices/1d?type=1,3,4,5,6,8,9,14,16&")
		warningData, err := M.GetApiData(params, "/v7/warning/now?")

		if err != nil {
			xzap.Error("和风 API 系统维护：" + err.Error())
			return nil, err
		}
		var alertDesc, forecastKeypoint, warmRemind string
		lifeSuggestion := &weather_mgr.LifeSuggestion{}

		for _, v := range normalData.Daily {
			switch {
			case v.Type == "1":
				// 	运动指数
				lifeSuggestion.Sport = v.Category

			case v.Type == "3":
				// 穿衣指数
				warmRemind = v.Level
				lifeSuggestion.Dressing = v.Category

			case v.Type == "4":
				// 	钓鱼指数
				lifeSuggestion.Fishing = v.Category
			case v.Type == "5":
				// "紫外线指数"
				lifeSuggestion.Uv = v.Category
			case v.Type == "6":
				// 	旅游指数
				lifeSuggestion.Travel = v.Category
			case v.Type == "9":
				// 	感冒指数
				lifeSuggestion.Flu = v.Category
			case v.Type == "14":
				// 	晾晒指数
				lifeSuggestion.Airing = v.Category
			// case v.Type == "16":
			// 	// 	防晒指数
			// 	lifeSuggestion.Sunscreen = v.Category
			case v.Type == "8":
				// 舒适度指数
				forecastKeypoint = v.Text
			default:
				continue
			}
		}

		if normalData.Warning != nil && len(warningData.Warning) > 0 {
			alertDesc = warningData.Warning[0].Text
		}

		val.Realtime = &weather_mgr.TodayResp{
			ForecastKeypoint: forecastKeypoint,
			AlertDesc:        alertDesc,
			WarmRemind:       warmRemind,
			Date:             currentTime.Unix(),
			LifeSuggestion:   lifeSuggestion,
		}
		str, err := json.Marshal(val.Realtime)
		if err == nil {
			WeatherModel.SetWeatherRealTimeData(fmt.Sprintf("%s", params.CityCode), string(str))
		}
	case "daily":
		sevenData, err := M.GetApiData(params, "/v7/weather/15d?")
		if err != nil {
			xzap.Error("和风 API 系统维护(7 day)：" + err.Error())
			return nil, err
		}
		aqiData, err := M.GetApiData(params, "/v7/air/5d?")
		if err != nil {
			xzap.Error("和风 API 系统维护（aqi）：" + err.Error())
		}
		for k, v := range sevenData.Daily {
			// 保存当天地区 最低温度，用于贴心提示
			if k == 0 {
				go WeatherModel.SetLowDayTemp(params.CityCode, v.TempMin, currentTime)
				// 保存当天湿度
				go WeatherModel.SetHumidityDay(params.CityCode, v.Humidity, currentTime)
			}
			if k == 15 {
				break
			}
			aqi := float64(0)
			t, err1 := time.ParseInLocation("2006-01-02", strings.Split(v.FxDate, "+")[0], time.Local)
			if err1 != nil {
				xzap.Error(err1.Error())
				continue
			}

			if aqiData != nil && len(aqiData.Daily) > k {
				aqi, _ = strconv.ParseFloat(aqiData.Daily[k].Aqi, 64)
			}
			val.Daily = append(val.Daily, &weather_mgr.DailyStyle{
				Date:           t.Unix(),
				TemperatureMax: v.TempMax,
				TemperatureMin: v.TempMin,
				Sunset:         v.Sunset,
				Sunrise:        v.Sunrise,
				Aqi:            aqi,
				Skycon:         helper.GetWeatherSky(v.IconDay),
			})
		}

	case "hourly":
		hourlyData, err := M.GetApiData(params, "/v7/weather/24h?")
		if err != nil {
			xzap.Error("和风 API 系统维护：" + err.Error())
			return nil, err
		}
		for k, v := range hourlyData.Hourly {
			// 24 个截止
			if k == 24 {
				break
			}

			path := strings.Split(v.FxTime, "+")[0]
			t, err1 := time.ParseInLocation("2006-01-02T15:04", path, time.Local)
			if err1 != nil {
				xzap.Error(err1.Error())
				continue
			}
			val.Hourly = append(val.Hourly, &weather_mgr.HourlyStyle{
				Date:        t.Unix() - 3600,
				Skycon:      helper.GetWeatherSky(v.Icon),
				Temperature: v.Temp,
			})
		}

	default:
		xzap.Error("不存在该和风 api类似：" + params.WeatherType)
		err = helper.ERROR_NOTICE_API
		return
	}
	return val, nil
}

// 获取分钟级降水文案
func (M *hefengModel) GetRainData(params *weather_mgr.WeatherReq) (desc string) {
	// 分钟级降水   内层已打印返回结果
	rainData, _ := M.GetApiData(params, "/v7/minutely/5m?")
	if rainData != nil {
		return rainData.Summary
	}
	return
}

// 获取和风 通用数据
func (M *hefengModel) GetApiData(params *weather_mgr.WeatherReq, api string) (val *HefengData, err error) {
	hefengHost := internal.GetApolloCli().GetStringValue("hefeng.host", "application", "")
	hefengToken := internal.GetApolloCli().GetStringValue("hefeng.token", "application", "")
	if hefengHost == "" || hefengToken == "" {
		xzap.Error("和风 api 无效,联系开发者", zap.String("hefengHost", hefengHost), zap.String("hefengToken", hefengToken))
		return nil, fmt.Errorf("和风 api 无效 ")
	}
	// go xprometheus.AddPrometheusCityCode("hefeng", params.CityCode, api)

	path := fmt.Sprintf("%s%slocation=%.4f,%.4f&key=%s", hefengHost, api, params.Lng, params.Lat, hefengToken)

	resp, err := http.Get(path)
	if err != nil {
		xzap.Error("readFromHttp", zap.Error(err), zap.Any("path:", path))
		return
	}
	defer func() {
		resp.Body.Close()
	}()
	xzap.Debug("调用和风API ：", zap.Any("path:", path), zap.Any("params", params))

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		xzap.Error("readFromHttp", zap.Error(err), zap.Any("path:", path))
		return
	}

	response := new(HefengData)
	err = json.Unmarshal(body, response)
	if err != nil {
		xzap.Error("解析api 数据失败，body 为空，检查路由参数", zap.Error(err),
			zap.Any("resp", body), zap.Any("path:", path))
		return
	}

	if response.Code != "200" {
		err = fmt.Errorf("err msg:%s", response.Code)
		xzap.Error("readFromHttp", zap.Error(err), zap.Any("path:", path))
		return
	}

	return response, nil
}

// 异步更新缓存
func (M *hefengModel) AsyncWeatherData() {
	for {
		select {
		case data, ok := <-WeatherChan:
			if ok {
				if data.Params == nil {
					continue
				}
				params := data.Params
				// currentTime := params.Time
				currentTime := time.Now()
				res, err := HefengModel.GetFormatData(params)
				if err != nil {
					xzap.Error("和风 API 系统维护：" + err.Error())
					continue
				}
				switch params.WeatherType {
				// case "today":
				// 	str, err := json.Marshal(res.Realtime)
				// 	if err == nil {
				// 		WeatherModel.SetWeatherRealTimeData(fmt.Sprintf("%s", params.CityCode), string(str))
				// 	}
				case "daily":
					str, err := json.Marshal(res.Daily)
					if err == nil {
						WeatherModel.SetWeatherDailyData(fmt.Sprintf("%s", params.CityCode), string(str))
						// 保存今天数据，提供给明天使用
						lastDaily, err1 := json.Marshal(res.Daily[0])
						if err1 == nil {
							WeatherModel.SetLastDay(fmt.Sprintf("%s:%d", params.CityCode, currentTime.Day()), string(lastDaily))
						}
					}
				case "hourly":
					str, err := json.Marshal(res.Hourly)
					if err == nil {
						WeatherModel.SetWeatherHourlyData(fmt.Sprintf("%s", params.CityCode), string(str), currentTime)
						// 保存当前数据，提供给下个小时使用
						lastHourly, err1 := json.Marshal(res.Hourly[0])
						if err1 == nil {
							WeatherModel.SetLastHour(fmt.Sprintf("%s:%d", params.CityCode, currentTime.Hour()), string(lastHourly), currentTime)
						}
					}
				default:
					xzap.Error("不存在该和风 api 类型：" + params.WeatherType)
					// TODO：通知飞书 和风 api  ERR
				}
			} else {
				// TODO：通知飞书 异步缓存异常关闭
				break
			}
		}
	}

}
