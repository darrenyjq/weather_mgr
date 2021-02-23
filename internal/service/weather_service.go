package service

import (
	"context"
	"weather_mgr/cootek/pgd/weather_mgr"
	"weather_mgr/helper"
	"weather_mgr/internal/model"
)

type WeatherService struct {
}

// 24小时 天气情况
func (w *WeatherService) Hourly(c context.Context, params *weather_mgr.HourlyReq) (*weather_mgr.HourlyResp, error) {
	param := &helper.ParamWeatherInfo{
		Lat:      params.Lat,
		Lng:      params.Lng,
		CityCode: params.CityCode,
	}
	res, err := model.HefengModel.GetFormatData(param)
	resp := new(weather_mgr.HourlyResp)
	if err == nil {
		resp.List = []*weather_mgr.HourlyStyle{}
		resp.List = res.Hourly
	}

	return resp, err
}

// 近一周天气情况
func (w *WeatherService) Daily(c context.Context, params *weather_mgr.DailyReq) (*weather_mgr.DailyResp, error) {
	println(3333333)

	return nil, nil
}

// 当天生活指数情况
func (w *WeatherService) Today(c context.Context, params *weather_mgr.TodayReq) (*weather_mgr.TodayResp, error) {
	println(4444444)

	return nil, nil
}
