package service

var (
	WeatherServ *WeatherService
)

func init() {
	WeatherServ = new(WeatherService)
}
