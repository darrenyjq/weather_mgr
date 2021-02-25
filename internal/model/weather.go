package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"math/rand"
	"strconv"
	"time"
	"weather_mgr/cootek/pgd/weather_mgr"
	"weather_mgr/internal/model/key"
	"weather_mgr/pkg/xzap"
)

type weatherModel struct {
}

var WeatherModel *weatherModel

func init() {
	WeatherModel = new(weatherModel)
}

// 获取上一天数据
func (M *weatherModel) GetLastDayData(ctx context.Context, cityCode string) (lastData []*weather_mgr.DailyStyle) {

	currentTime := GetTimeNow(ctx)

	// 可以忽略月初 1 号问题
	var lastDayTime int
	lastDayTime = currentTime.AddDate(0, 0, -1).Day()

	// 取前一天整点
	lastDay, err2 := M.GetLastDay(fmt.Sprintf("%s:%d", cityCode, lastDayTime))
	temp := &weather_mgr.DailyStyle{}
	if err2 == nil && lastDay != "" {
		err := json.Unmarshal([]byte(lastDay), temp)
		if err != nil {
			return
		}
		return []*weather_mgr.DailyStyle{temp}
	}
	return []*weather_mgr.DailyStyle{}
}

// 获取上一小时数据
func (M *weatherModel) GetLastHourlyData(ctx context.Context, cityCode string) (lastData []*weather_mgr.HourlyStyle) {
	currentTime := GetTimeNow(ctx)
	// 可以忽略 0 点
	timestamp := currentTime.Add(-time.Hour).Hour()

	lastHour, err2 := M.GetLastHour(fmt.Sprintf("%s:%d", cityCode, timestamp))
	temp := &weather_mgr.HourlyStyle{}
	if err2 == nil && lastHour != "" {
		err := json.Unmarshal([]byte(lastHour), temp)
		if err != nil {
			return
		}
		return []*weather_mgr.HourlyStyle{temp}
	}
	return []*weather_mgr.HourlyStyle{}
}

// 当天总奖励次数
func (M *weatherModel) TotalAwardTimes(uid uint64, ct time.Time) int {
	awardTimes := M.GetAwardTimes(uid, ct)
	awardTemp := M.GetAwardTemp(uid, ct)
	awardWatchTimes := M.GetAwardWatch(uid, ct)
	return awardTimes + awardTemp + awardWatchTimes
}

func (M *weatherModel) GetAwardTimes(uid uint64, ct time.Time) int {
	k := key.RKWeatherCountdownTimes(uid, ct)
	times, err := redisClient.Get(k).Int()
	if err == nil {
		return times
	}
	return 0
}
func (M *weatherModel) SetAwardTimes(uid uint64, ct time.Time) {
	k := key.RKWeatherCountdownTimes(uid, ct)
	redisClient.Incr(k)
	redisClient.Expire(k, time.Hour*24)
}

func (M *weatherModel) GetAwardWatch(uid uint64, ct time.Time) int {
	k := key.RKAwardWatch(uid, ct)
	times, err := redisClient.Get(k).Int()
	if err == nil {
		return times
	}
	return 0
}
func (M *weatherModel) IncrAwardWatch(uid uint64, ct time.Time) int64 {
	k := key.RKAwardWatch(uid, ct)
	times := redisClient.Incr(k).Val()
	redisClient.Expire(k, time.Hour*24)
	return times
}

func (M *weatherModel) GetAwardTemp(uid uint64, ct time.Time) int {
	k := key.RKWeatherAwardTemp(uid, ct)
	times, err := redisClient.Get(k).Int()
	if err == nil {
		return times
	}
	return 0
}
func (M *weatherModel) SetAwardTemp(uid uint64, ct time.Time) {
	k := key.RKWeatherAwardTemp(uid, ct)
	redisClient.Incr(k)
	redisClient.Expire(k, time.Hour*24)
}

func (M *weatherModel) GetAwardDate(uid uint64, currentTime time.Time) int64 {
	k := key.RKWeatherCountdownDate(uid)
	date, err := redisClient.Get(k).Int64()
	if err == nil && redis.Nil != err {
		return date
	}
	return currentTime.Unix()
}
func (M *weatherModel) SetAwardDate(uid uint64, times int64) {
	k := key.RKWeatherCountdownDate(uid)
	redisClient.Set(k, times, time.Hour*24)
}

func (M *weatherModel) GetNowTemp(uid uint64, ct time.Time) float64 {
	k := key.RKWeatherNowTemp(uid, ct)
	count, err := redisClient.Get(k).Float64()
	if err == nil {
		return count
	}
	// 没有缓存取全国平均值：
	return 10
}
func (M *weatherModel) SetNowTemp(uid uint64, temp float64, ct time.Time) {
	k := key.RKWeatherNowTemp(uid, ct)
	redisClient.Set(k, temp, time.Hour*24)
}

func (M *weatherModel) GetLowDayTemp(cityCode string, ct time.Time) int64 {
	k := key.RKLowDayTemp(cityCode, ct)
	count, err := redisClient.Get(k).Int64()
	if err == nil {
		return count
	}
	return 999
}
func (M *weatherModel) SetLowDayTemp(cityCode string, temp string, ct time.Time) {
	k := key.RKLowDayTemp(cityCode, ct)
	t, _ := strconv.ParseInt(temp, 10, 64)
	redisClient.Set(k, t, time.Hour*48)
}

func (M *weatherModel) GetHumidityDay(cityCode string, ct time.Time) string {
	k := key.RKHumidityDay(cityCode, ct)
	humidity := redisClient.Get(k).Val()
	return humidity
}
func (M *weatherModel) SetHumidityDay(cityCode string, humidity string, ct time.Time) {
	k := key.RKHumidityDay(cityCode, ct)
	redisClient.Set(k, humidity, time.Hour*48)
}

func (M *weatherModel) GetTodayAwardCount(uid uint64, ct time.Time) int {
	k := key.RKWeatherCountdownAwardCount(uid, ct)
	count, err := redisClient.Get(k).Int()
	if err == nil {
		return count
	}
	return 0
}
func (M *weatherModel) SetTodayAwardCount(uid uint64, coins int, ct time.Time) {
	k := key.RKWeatherCountdownAwardCount(uid, ct)
	redisClient.Set(k, coins, time.Hour*24)
}

func (M *weatherModel) GetTempTodayAwardCount(uid uint64, ct time.Time) int {
	k := key.RKWeatherTempAwardCount(uid, ct)
	count, err := redisClient.Get(k).Int()
	if err == nil {
		return count
	}
	return 0
}

// 今日已经奖励金币数目
func (M *weatherModel) SetTempTodayAwardCount(uid uint64, coins int, ct time.Time) {
	k := key.RKWeatherTempAwardCount(uid, ct)
	redisClient.Set(k, coins, time.Hour*24)
}

func (M *weatherModel) GetCountdownTimeCoins(currentTime time.Time, times int) (coins int, nextTime int64) {
	// min := currentTime.Sub(time.Unix(lastTime, 0)).Minutes()
	rand.Seed(currentTime.UnixNano())
	lastTime := currentTime.Unix()
	switch {
	case 1 == times:
		return 18, lastTime + 30
	case 2 == times:
		// case 2 == times && min > 30:
		return rand.Intn(19) + 1, lastTime + 60
	case 3 == times:
		// case 3 == times && min > 60:
		return rand.Intn(4) + 21, lastTime + 30
	case 4 == times:
		// case 4 == times && min > 30:
		return rand.Intn(4) + 11, lastTime + 60
	case 5 == times:
		// case 5 == times && min > 60:
		return rand.Intn(4) + 16, lastTime + 120
	case 6 == times:
		// case 6 == times && min > 120:
		return rand.Intn(4) + 21, lastTime + 300
	case 7 == times:
		// case 7 == times && min > 300:
		return rand.Intn(4) + 26, lastTime + 30
	case 8 == times:
		// case 8 == times && min > 30:
		return rand.Intn(2) + 3, lastTime + 60
	case 9 == times:
		// case 9 == times && min > 60:
		return rand.Intn(1) + 6, lastTime + 30
	case 10 == times:
		// case 10 == times && min > 30:
		return rand.Intn(4) + 2, lastTime + 180
	case 11 == times:
		// case 11 == times && min > 180:
		return rand.Intn(4) + 6, lastTime + 600
	case 12 == times:
		// case 12 == times && min > 600:
		return rand.Intn(9) + 11, lastTime + 18600
	}
	xzap.Debug("异常参数！")
	return 0, lastTime
}

// 获取温差奖励数目
func (M *weatherModel) GetTempCoins(uid uint64, todayAwardCoins int, ct time.Time) (coins int) {
	// var nowTemp int
	lastTemp := M.GetLastTemps(uid, ct)
	lens := len(lastTemp)
	// 当日第一次 给 88
	if todayAwardCoins == 0 {
		return 88
	}
	// 不满足两次 返回 0个 coin
	if lens < 2 {
		return 0
	}
	var one, two float64
	var err error
	one, err = strconv.ParseFloat(lastTemp[lens-1], 64)
	if err != nil {
		one = 0
	}

	two, err = strconv.ParseFloat(lastTemp[lens-2], 64)
	if err != nil {
		two = 0
	}
	res := (one - two) * 10
	// 取温度绝对值
	if res < 0 {
		res = res * -1
	}

	if res > 100 {
		// TODO ****

		res = 100
	}

	return int(res)
}

func (M *weatherModel) GetLastTemps(uid uint64, ct time.Time) (res []string) {
	k := key.RKWeatherLastTemps(uid, ct)
	var err error
	res, err = redisClient.ZRange(k, 0, -1).Result()
	if err != nil || err == redis.Nil {
		return []string{}
	}
	return res
}
func (M *weatherModel) SetLastTemps(uid uint64, ct time.Time) {
	var temp float64
	// 获取当前时间温度
	temp = M.GetNowTemp(uid, ct)
	// 此处同一个小时内温度不变化 插入无效
	k := key.RKWeatherLastTemps(uid, ct)
	redisClient.ZAdd(k, redis.Z{float64(ct.Hour()), temp})
	redisClient.Expire(k, 24*time.Hour)
}

// 重置温差
func (M *weatherModel) DelLastTemps(uid uint64, ct time.Time) {
	k := key.RKWeatherLastTemps(uid, ct)
	redisClient.Del(k)

}

// 当前时间段 天气状况
func (M *weatherModel) SetWeatherHourlyData(cityCode, data string, currentTime time.Time) {
	k := key.RedisKvWeatherHourlyData(cityCode)
	// 取距离下个小时的剩余分钟数
	// t := uint32(59 - currentTime.Minute())
	redisClient.Set(k, data, 0)
}

func (M *weatherModel) GetWeatherHourlyData(cityCode string) (res string, err error) {
	k := key.RedisKvWeatherHourlyData(cityCode)
	res, err = redisClient.Get(k).Result()
	return
}

// 当前天数据
func (M *weatherModel) SetWeatherDailyData(cityCode, data string) {
	k := key.RedisKvWeatherDailyData(cityCode)
	redisClient.Set(k, data, 0)
}

func (M *weatherModel) GetWeatherDailyData(cityCode string) (res string, err error) {
	k := key.RedisKvWeatherDailyData(cityCode)
	res, err = redisClient.Get(k).Result()
	return
}

// 当前天 天气个性化数据
func (M *weatherModel) SetWeatherRealTimeData(cityCode, data string) {
	k := key.RedisKvWeatherRealTimeData(cityCode)
	redisClient.Set(k, data, 0)
}

func (M *weatherModel) GetWeatherRealTimeData(cityCode string) (res string, err error) {
	k := key.RedisKvWeatherRealTimeData(cityCode)
	res, err = redisClient.Get(k).Result()
	return
}

func (M *weatherModel) SetLastDay(cityCode, data string) {
	k := key.RedisKvWeatherLastDay(cityCode)
	redisClient.Set(k, data, time.Hour*24)
}

func (M *weatherModel) GetLastDay(cityCode string) (res string, err error) {
	k := key.RedisKvWeatherLastDay(cityCode)
	res, err = redisClient.Get(k).Result()
	return
}

func (M *weatherModel) SetLastHour(cityCode, data string, currentTime time.Time) {
	k := key.RedisKvWeatherLastHour(cityCode)
	// 取距离下个小时的剩余分钟数
	t := 119 - currentTime.Minute()
	redisClient.Set(k, data, time.Minute*time.Duration(t))
}

func (M *weatherModel) GetLastHour(cityCode string) (res string, err error) {
	k := key.RedisKvWeatherLastHour(cityCode)
	res, err = redisClient.Get(k).Result()
	return
}
