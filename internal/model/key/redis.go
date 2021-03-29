package key

import (
	"fmt"
	"time"
)

func RedisCoinCount(appName string, dateYmd string) string {
	return fmt.Sprintf("WALKING_COIN_COUNT:%s:%s", appName, dateYmd)
}
func RedisDeactivateList() string {
	return fmt.Sprintf("DEACTIVATE_UID_LIST")
}

func RedisNewcomerIncentiveStats(uid int64) string {
	return fmt.Sprintf("weather_new_comer_gift_incentive_stats:%d", uid)
}

func RedisNewcomerIncentiveGiftProgress(uid int64, eventId string) string {
	return fmt.Sprintf("weather_new_comer_gift_incentive_process:%s:%d", eventId, uid)
}

func RedisCoinBasicInfo(uid int64, appName string) string {
	return fmt.Sprintf("weather_coin_basic_info:%s:%d", appName, uid)
}

func RedisOpenIdInfo(uid int64) string {
	return fmt.Sprintf("weather_openid_info:%d", uid)
}

func RedisPreBindRecord(unionId string) string {
	return fmt.Sprintf("weather_wecgat_prebind_relation:%s", unionId)
}

func RedisCountIncentiveInfo(uid int64, t time.Time) string {
	return fmt.Sprintf("weather_statistic_incentive:%d:%s", uid, t.Format("2006-01-02"))
}

func RedisCountTotalIncentiveInfo(uid int64) string {
	return fmt.Sprintf("weather_statistic_total_incentive:%d", uid)
}

func RedisCoinBallRecord(uid int64, t time.Time) string {
	return fmt.Sprintf("coin_ball_record:%d:%s", uid, t.Format("2006-01-02"))
}

func RedisCoinBallCompleteTime(uid int64, t time.Time) string {
	return fmt.Sprintf("coin_ball_complete:%d:%s", uid, t.Format("2006-01-02"))
}

func RedisIncentiveInviteeV1Info(uid int64) string {
	return fmt.Sprintf("incentive_invite_v1_info%d", uid)
}

func RedisIncentiveInviteeReward(uid int64, periodId string) string {
	return fmt.Sprintf("incentive_invite_reward:%d:%s", uid, periodId)
}

func RedisKvLockHandleIncentiveAction(uid int64) string {
	return fmt.Sprintf("handle_incentive_action:%d", uid)
}

func RedisKvLockAddReward(uid int64) string {
	return fmt.Sprintf("handle_incentive_action:%d", uid)
}

func RedisKvFriendsInfo(uid int64, tag string) string {
	return fmt.Sprintf("friends_info:%d:%s", uid, tag)
}

func RedisEnergyRecord(uid int64, t time.Time) string {
	return fmt.Sprintf("engery_reward_:%d:%s", uid, t.Format("2006-01-02"))
}
func RedisEventTriggered(uid int64, eventId string) string {
	return fmt.Sprintf("event_triggered:%d:%s", uid, eventId)
}

func RedisIdiomUserPlayInfo(uid int64) string {
	return fmt.Sprintf("idiom_user_play_info:%d", uid)
}

func RedisIdiomUserDailyInfo(uid int64, t time.Time) string {
	return fmt.Sprintf("idiom_user_daily_info:%d:%s", uid, t.Format("20060102"))
}
func RedisWxShareBlackList() string {
	return "wx_share_blacklist"
}

func RedisWxSharePageView() string {
	return "wx_share_url_page_view"
}

func RedisWithdrawBadPeople() string {
	return "BAD_PEOPLE_UID_LIST"
}
func RedisCashWheelTodayDrawNum(uid int64, t time.Time) string {
	return fmt.Sprintf("cw_draw_times:%s:%d", t.Format("20060102"), uid)
}

func RedisCashWheelTotalDrawNum(uid int64) string {
	return fmt.Sprintf("cw_total_draw_times:%d", uid)
}

func RedisCashWheelRedPacket(uid int64, packetKey string) string {
	return fmt.Sprintf("cw_redpacket:%d:%s", uid, packetKey)
}

func RedisCashWheelMillstone(uid int64, t time.Time) string {
	return fmt.Sprintf("cw_millstones:%s:%d", t.Format("20060102"), uid)
}

func RedisCashWheelCoinRecord(uid int64) string {
	return fmt.Sprintf("cw_coin_statistic:%d", uid)
}

func RedisWithdrawAmountList(amount int) string {
	return fmt.Sprintf("withdraw_success_list_%d", amount)
}

func RedisWithdrawSuccessNotice() string {
	return "withdraw_success_notice"
}

func RedisTimeBenefitCollectRecord(uid int64) string {
	return fmt.Sprintf("t_benefit_record:%d", uid)
}

func RedisIdiomRedPacket(uid int64, packetKey string) string {
	return fmt.Sprintf("idiom_packet:%d:%s", uid, packetKey)
}

func RedisIdiomActionPointInfo(uid int64) string {
	return fmt.Sprintf("idiom_ap_info:%d", uid)
}

func RedisIdiomActionPointRestoreTime(uid int64) string {
	return fmt.Sprintf("idiom_ap_restore:%d", uid)
}

func RedisIdiomAchievementInfo(uid int64) string {
	return fmt.Sprintf("idiom_achi_info:%d", uid)
}

func RedisIdiomAchiGroup(uid int64) string {
	return fmt.Sprintf("idiom_achi_group:%d", uid)
}

func RedisCountBurstLimiter(key string, date string) string {
	return fmt.Sprintf("weather_count_burst_limiter:%s:%s", key, date)
}

func RedisStepMoneyStashInfo(uid int64) string {
	return fmt.Sprintf("weather_sm_stash_info:%d", uid)
}
func RedisStepMoneyItemInfo(uid int64) string {
	return fmt.Sprintf("weather_sm_item_info:%d", uid)
}

func RedisStepMoneyPathLock(uid int64) string {
	return fmt.Sprintf("weather_sm_lock:%d", uid)
}

func RedisDigPlayTime(uid int64, t time.Time) string {
	return fmt.Sprintf("weather_dig_info:%s:%d", t.Format("2006-01-02"), uid)
}

func RedisDigRedPacket(uid int64, packetKey string) string {
	return fmt.Sprintf("dig_redpacket:%d:%s", uid, packetKey)
}

func RedisIncentivePeriodData(uid int64, period string) string {
	return fmt.Sprintf("incentive_100:%d:%s", uid, period)
}

func RedisIncentive100Reward(uid int64, period string) string {
	return fmt.Sprintf("incentive_100_reward:%d:%s", uid, period)
}

// //////////累计签到///////////////
// 用户签到周期数
func RedisKvSignIn2Period(uid uint64) string {
	return fmt.Sprintf("sgin:period:%d", uid)
}

// 用户周期下签到的天数
func RedisSetSignIn2PeriodDay(uid uint64, period int64) string {
	return fmt.Sprintf("sgin:period:day:%d:%d", uid, period)
}

// 首次签到日期
func RKWeatherFirstSignDay(uid uint64) string {
	return fmt.Sprintf("weather:first:sign:day:%d", uid)
}

// 第七天是否送礼
func RKWeatherSevenGift(uid uint64) string {
	return fmt.Sprintf("weather:seven:gift:%d", uid)
}

//
// // /////////////连续签到///////////////
// // 锁
// func RedisKvLocker(uid uint64, event string) string {
// 	return fmt.Sprintf("bd_locker_%d_%s", uid, event)
// }
//
// // 签到开始天数
// func RedisKvSignInStartDay(uid uint64) string {
// 	return fmt.Sprintf("bd_signin_start_%d", uid)
// }
//
// // 签到周期内签到的状态
// func RedisSetSignInDays(uid uint64, startDay string) string {
// 	return fmt.Sprintf("bd_signin_days_%d_%s", uid, startDay)
// }
//
// // 签到周期次数
// func RedisKvSignInPeriodNum(uid uint64) string {
// 	return fmt.Sprintf("bd_signin_period_%d", uid)
// }
//
// // 签到总次数
// func RedisKvSignInTotalNum(uid uint64) string {
// 	return fmt.Sprintf("bd_signin_total_%d", uid)
// }

// 最后的签到日期
func RedisKvSignInLastDay(uid uint64) string {
	return fmt.Sprintf("bd_signin_last_day_%d", uid)
}

// 连续签到天数
func RedisKvSignInContinuousNum(uid uint64) string {
	return fmt.Sprintf("bd_signin_cont_num_%d", uid)
}

// 天气热缓存/小时
func RedisKvWeatherHourlyData(cityCode string) string {
	return fmt.Sprintf("weather:hourly:data:%s", cityCode)
}

// 天气个性数据/天
func RedisKvWeatherRealTimeData(cityCode string) string {
	return fmt.Sprintf("weather:real_time:data:%s", cityCode)
}

// 天气预警
func RedisWarningList(cityCode string) string {
	return fmt.Sprintf("weather:warning_list:%s", cityCode)
}

// 天气数据/天
func RedisKvWeatherDailyData(cityCode string) string {
	return fmt.Sprintf("weather:daily:data:%s", cityCode)
}

// 上一天天气热缓存
func RedisKvWeatherLastDay(cityCode string) string {
	return fmt.Sprintf("weather:last:day:%s", cityCode)
}

// 上一小时 缓存
func RedisKvWeatherLastHour(cityCode string) string {
	return fmt.Sprintf("weather:last:hour:%s", cityCode)
}

// 保存该地区降雨情况
func RedisCurRain(cityCode string) string {
	return fmt.Sprintf("weather:cur:rain:%s", cityCode)
}

// 保存当天出行信息
func RedisWalkOut(cityCode string) string {
	return fmt.Sprintf("weather:walk:out:%s", cityCode)
}

// 倒计时下一次领取金币时间
func RKWeatherCountdownDate(uid uint64) string {
	return fmt.Sprintf("weather:countdown:date:uid:%d", uid)
}

// 倒计时 今日已经领取次数
func RKWeatherCountdownTimes(uid uint64, ct time.Time) string {
	return fmt.Sprintf("weather:countdown:times:uid:%d:day:%d", uid, ct.Day())
}

// 温差奖励 今日已经领取次数
func RKWeatherAwardTemp(uid uint64, ct time.Time) string {
	return fmt.Sprintf("weather:award:temp:uid:%d:day:%d", uid, ct.Day())
}

// 【商业化模块】看视频领金币
func RKAwardWatch(uid uint64, ct time.Time) string {
	return fmt.Sprintf("weather:award:watch:uid:%d:day:%d", uid, ct.Day())
}

// 离线奖励
func RkOffLineTime(uid uint64) string {
	return fmt.Sprintf("offline:time:uid:%d", uid)
}

// 倒计时 今日已经领取数量
func RKWeatherCountdownAwardCount(uid uint64, ct time.Time) string {
	return fmt.Sprintf("countdown:award_count:uid:%d:day:%d", uid, ct.Day())
}

// 温差 今日已经领取数量
func RKWeatherTempAwardCount(uid uint64, ct time.Time) string {
	return fmt.Sprintf("temp:award_count:uid:%d:day:%d", uid, ct.Day())
}

// 离线奖励 已经领取数量
func RKOffLineAwardCount(uid uint64) string {
	return fmt.Sprintf("offline:award:count:uid:%d", uid)
}

// 温差 进两次的温度
func RKWeatherLastTemps(uid uint64, ct time.Time) string {
	return fmt.Sprintf("weather:last:temps:uid:%d:day:%d", uid, ct.Day())
}

// 温差 每次进入温度记录
func RKWeatherNowTemp(uid uint64, ct time.Time) string {
	return fmt.Sprintf("now:temp:uid:%d:day:%d", uid, ct.Day())
}

// 温差 每次进入温度记录
func RKLowDayTemp(cityCode string, ct time.Time) string {
	return fmt.Sprintf("low:day:temp:uid:%s:day:%d", cityCode, ct.Day())
}

// 湿度
func RKHumidityDay(cityCode string, ct time.Time) string {
	return fmt.Sprintf("humidity:day:%d:uid:%s", ct.Day(), cityCode)
}

func RedisDailyTask(uid int64, taskId string, t time.Time) string {
	return fmt.Sprintf("DailyTask:%d:%s:%s", uid, t.Format("2006-01-02"), taskId)
}
func RKAwardWatchCountDown(uid int64, taskId string, t time.Time) string {
	return fmt.Sprintf("award:watch:count_down:%d:%d:%s", uid, t.Day(), taskId)
}

func RedisHashDailyTaskProgress(uid int64, today string) string {
	return fmt.Sprintf("DailyTaskProg:%d:%s", uid, today)
}

func RedisTaskCompleteTimesTotal(uid int64, taskId string) string {
	return fmt.Sprintf("LimitTask:%d:%s", uid, taskId)
}

// //////////任务///////////////
// 一次性任务
func RedisHashTaskTypeTotalNum(uid uint64) string {
	return fmt.Sprintf("task:total:num:%d", uid)
}

// 每日任务
func RedisHashTaskTypeDailyNum(uid uint64, day string) string {
	return fmt.Sprintf("task:daily:num:%d:%s", uid, day)
}

// //////////微信/手机 登录 提现///////////////
func RKUserInfoType(uid uint64, loginType int64) string {
	return fmt.Sprintf("weather:user:info:%d:type:%d", uid, loginType)
}

func RKAccountInfo(account string, accountType int64) string {
	return fmt.Sprintf("weather:account:info:%s:type:%d", account, accountType)
}

func RKFirstWithdrew(uid int64) string {
	return fmt.Sprintf("weather:first:withdrew:%d", uid)
}

func RKWeatherCashCache(uid uint64) string {
	return fmt.Sprintf("weather:cash:cache:%d", uid)
}

func RKAwardTypeCoin(key, day string, uid uint64) string {
	return fmt.Sprintf("award:type:coin:uid:%s:%d:%s", key, uid, day)
}

func RKIsWatchedVideo(uid uint64, day int) string {
	return fmt.Sprintf("is:watched:video:%d:day:%d", uid, day)
}

func RedisKvCityCode(uid uint64) string {
	return fmt.Sprintf("city_code:%d", uid)
}

// //////////新人礼///////////////

// 记录用户是否可以重复领取30
func RedisWithdrawV2Reset30(uid int64) string {
	return fmt.Sprintf("wv2r_30:%d", uid)
}

// 记录用户是否可以重复领取30
func RedisNewUserTaskReset(uid int64) string {
	return fmt.Sprintf("nutr_:%d", uid)
}

// 新人红包白名单
func RKNewUserWhiteList() string {
	return fmt.Sprintf("new:user:red:whitelist")
}

// 新人红包白名单已消费
func RKNewUserWhiteConsumedWhiteList() string {
	return fmt.Sprintf("new:user:red:consumed_whitelist")
}
