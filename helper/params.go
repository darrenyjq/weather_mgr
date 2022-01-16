package helper

import (
	"encoding/json"
	"time"
	"weather_mgr/bbbb/pgd/weather_mgr"
)

type (
	ParamEncryptData struct {
		Data     string `json:"data"`
		SignType string `json:"sign_type"`
	}

	ParamBind struct {
		Code string `json:"code"`
	}

	ParamBingoFlipEvaluate struct {
		Key         string             `json:"key"`
		CollectItem []BingoCollectItem `json:"collect_item"`
	}

	BingoCollectItem struct {
		Type   string `json:"type"`
		Amount int    `json:"amount"`
	}

	BingoGameRewardInfo struct {
		Key        string `json:"key"`
		Ufo        int    `json:"ufo"`
		Battery    int    `json:"battery"`
		NormalCoin int    `json:"normal_coin"`
	}

	BingoRedisSaveInfo struct {
		Reward      *BingoGameRewardInfo `json:"reward"`
		CollectItem []*BingoCollectItem  `json:"collect_item"`
	}

	CashAccountInfoResp struct {
		Data      CashAccountInfo `json:"data"`
		ErrorCode int             `json:"error_code"`
		Msg       string          `json:"msg"`
		TimeUnix  int             `json:"time_unix"`
	}

	CashAccountInfo struct {
		Amount    int `json:"amount"`
		TicketNum int `json:"ticket_num"`
	}
	CashIncrResp struct {
		Data      CashIncrRet `json:"data"`
		ErrorCode int         `json:"error_code"`
		Msg       string      `json:"msg"`
		TimeUnix  int         `json:"time_unix"`
	}
	TicketIncrResp struct {
		Data      TicketIncrRet `json:"data"`
		ErrorCode int           `json:"error_code"`
		Msg       string        `json:"msg"`
		TimeUnix  int           `json:"time_unix"`
	}
	TicketIncrRet struct {
		TicketNum int `json:"ticket_num"`
	}
	CashIncrRet struct {
		Amount int `json:"amount"`
	}

	WithdrawInfoParam struct {
		WithdrawType string `json:"withdrawType"`
	}

	WithdrawInfoResp struct {
		Data      WithdrawInfo `json:"data"`
		ErrorCode int          `json:"error_code"`
		Msg       string       `json:"msg"`
		TimeUnix  int          `json:"time_unix"`
	}

	WithdrawGoods struct {
		Show                  bool   `json:"show"`                     // 是否页面展示
		GoodsID               int64  `json:"goods_id"`                 // 商品 id
		Bank                  string `json:"bank"`                     // 支付方式
		Amount                int64  `json:"amount"`                   // 价格
		TicketNum             int64  `json:"ticket_num"`               // 现金券数量
		GoodsType             string `json:"goods_type"`               // 商品性质 一次性 多次
		CashType              int64  `json:"cash_type"`                // 0 ：48 小时到账 ； 1 实时到账
		MaxWithdrawNum        int64  `json:"max_withdraw_num"`         // 最大提现次数
		CurrentWithdrawNum    int64  `json:"current_withdraw_num"`     // 当前提现次数
		DayMaxWithdrawNum     int64  `json:"day_max_withdraw_num"`     // 单日提现次数
		CurrentDayWithdrawNum int64  `json:"current_day_withdraw_num"` // 当日提现次数
		Group                 int64  `json:"group"`
	}

	WithdrawInfo struct {
		Amount    int64           `json:"amount"`     // 价格
		Goods     []WithdrawGoods `json:"goods"`      //  商品
		Notice    []string        `json:"notice"`     // 提醒
		Rules     string          `json:"rules"`      // 规则
		TicketNum int64           `json:"ticket_num"` // 现金券
		UID       int64           `json:"uid"`        // 用户 id
	}

	MealBenefitInfo struct {
		ID            string        `json:"id"`
		RewardCoin    int           `json:"reward_coin"`
		State         string        `json:"state"`
		Name          string        `json:"name"`
		TimeText      string        `json:"time_text"`
		StartDuration time.Duration `json:"-"`
		EndDuration   time.Duration `json:"-"`
	}

	MealBenefitParam struct {
		Meal     string `json:"meal"`
		MealType int    `json:"meal_type"`
	}

	WorkoutBenefitParam struct {
		Id string `json:"id"`
	}

	InviteEventParam struct {
		EventID  string          `json:"event_id"`
		Extra    json.RawMessage `json:"extra"`
		Category string          `json:"category"`
	}
	CommonEventParam struct {
		Events []CommonEvent `json:"events"`
	}
	CommonEvent struct {
		EventID string `json:"event_id"`
		Extra   string `json:"extra"`
	}
	CommonEventGetParam struct {
		EventIDs []string `json:"event_ids"`
	}

	InviteEventConfig struct {
		EventID      string `json:"event_id"`
		Type         string `json:"type"`
		CashRewardId string `json:"cash_reward_id"`
		DailyLimit   int64  `json:"daily_limit"`
		TotalLimit   int64  `json:"total_limit"`
	}

	InviteInfo struct {
		InviteFriends []InviteFriendStats `json:"invite_friends"`
		RewardInfo    InviteRewardInfo    `json:"invite_info"`
	}

	InviteFriendDaily struct {
		CurrentTime int    `json:"current_time"`
		ID          string `json:"id"`
		MaxTime     int    `json:"max_time"`
		Reward      int    `json:"reward"`
	}

	InviteFriendStats struct {
		AccountType string              `json:"account_type"`
		Avatar      string              `json:"avatar"`
		Daily       []InviteFriendDaily `json:"daily"`
		ID          int64               `json:"id"`
		Name        string              `json:"name"`
		State       int                 `json:"state"`
	}

	InviteRewardInfo struct {
		DailyReward   int `json:"daily_reward"`
		InviteFriends int `json:"invite_friends"`
		TotalReward   int `json:"total_reward"`
	}
	SleepBenefitActionParam struct {
		Action      string `json:"action"`
		SleepID     string `json:"sleep_id"`
		IsIncentive bool   `json:"is_incentive"`
	}

	SleepBenefitRewardParam struct {
		RewardID string `json:"reward_id"`
	}

	SleepBenefitReward struct {
		ID               string `json:"id"`
		RewardCoinAmount int    `json:"reward_coin_amount"`
		ValidTime        int64  `json:"valid_time"`
		TriggerDateYmd   string `json:"trigger_date_ymd"`
	}

	SleepBenefitInfo struct {
		SleepId        string
		ActiveTimeStr  []string
		Name           string
		MaxCoinReward  int
		PreOrderReward int
		StartDuration  time.Duration
		EndDuration    time.Duration
	}

	SleepBenefitCache struct {
		SleepDateID string `json:"sleep_date_id"`
		Action      string `json:"action"`
		ActionTs    int64  `json:"action_ts"`
		SleepId     string `json:"sleep_id"`
		DateYmd     string `json:"date_ymd"`
	}

	WorkoutBenefitInfo struct {
		ID            string        `json:"id"`
		RewardCoin    int           `json:"reward_coin"`
		Duration      int           `json:"duration"`
		State         string        `json:"state"`
		Name          string        `json:"name"`
		Desc          string        `json:"desc"`
		StartDuration time.Duration `json:"-"`
		EndDuration   time.Duration `json:"-"`
	}

	WorkoutStartInfo struct {
		ID        string `json:"id"`
		StartTime int64  `json:"start_time"`
		Duration  int    `json:"duration"`
	}

	StepBenefitInfo struct {
		ID                  string `json:"id"`
		Target              int    `json:"target"`
		Index               int    `json:"index"`
		RewardCoin          int    `json:"reward_coin"`
		IncentiveRewardCoin int    `json:"incentive_reward_coin"`
		State               string `json:"state"`
	}

	StepBenefitParam struct {
		Id         string `json:"id"`
		RewardType string `json:"reward_type"`
	}

	NewcomerIncentiveParam struct {
		RewardType string `json:"reward_type"`
		ActionType string `json:"action_type"`
		Count      int    `json:"count"`
	}

	// 微信code获取openid
	ParamCode2OpenId struct {
		Code   string `json:"code"`   // authcode 码
		Source string `json:"source"` // 可选参数
	}

	AuthCodeResp struct {
		AccessToken  string `json:"access_token"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		Openid       string `json:"openid"`
		Scope        string `json:"scope"`
		Unionid      string `json:"unionid"`
	}

	CheckPreBindParam struct {
		UnionId string `json:"union_id"`
	}

	InviteRewardV3Info struct {
		InviteRange      []int `json:"invite_range"`
		SingleCashAmount int   `json:"single_cash_amount"`
		TotalCashAmount  int   `json:"total_cash_amount"`
		Distribution     []int `json:"distribution"`
	}

	ParamStatisticAction struct {
		Action string `json:"action"` // 登录状态 已登录："login" 未登录 "notlogin"
	}

	ParamStatisticIncentive struct {
		Tu     string `json:"tu"`
		Times  int64  `json:"times"`
		Action string `json:"action"`
	}

	ParamStatisticIncentiveCount struct {
		Tu string `json:"tu"`
	}

	CoinBallConfig struct {
		ID       string  `json:"id"`
		Reward   int     `json:"reward"`
		MaxTimes int     `json:"max_times"`
		Duration []int64 `json:"duration"`
	}

	ParamCoinBallCollect struct {
		ID string `json:"id"`
	}

	WXUserInfo struct {
		Openid     string        `json:"openid"`
		Nickname   string        `json:"nickname"`
		Sex        int           `json:"sex"`
		Language   string        `json:"language"`
		City       string        `json:"city"`
		Province   string        `json:"province"`
		Country    string        `json:"country"`
		HeadImgUrl string        `json:"headimgurl"`
		Privilege  []interface{} `json:"privilege"`
		UnionId    string        `json:"unionid"`
	}

	ParamInviteIncentiveV1 struct {
		Action string `json:"action"`
		Period string `json:"period"`
	}

	IncentiveInviteRewardInfo struct {
		Action    string `json:"action"`
		Reward    int    `json:"reward"`
		Supporter string `json:"supporter"`
		Ts        int64  `json:"ts"`
	}

	InviteIncentiveSupporterInfo struct {
		NickName     string `json:"nick_name"`
		HeadURL      string `json:"head_url"`
		Type         string `json:"type"`
		Contribution int    `json:"contribution"`
	}

	SimpleUserInfo struct {
		UnionID  string `json:"union_id"`
		Nickname string `json:"nickname"`
		HeadUrl  string `json:"headurl"`
	}
	WithdrawResponse struct {
		ErrorCode int64  `json:"error_code"`
		Msg       string `json:"msg"`
		TimeUnix  int64  `json:"time_unix"`
	}

	WithdrawInnerOrderDealResponse struct {
		WithdrawResponse
		Data WithdrawInnerOrderDealData `json:"data"`
	}
	WithdrawInnerOrderDealData struct {
		Amount    int64         `json:"amount"`     // 数额
		IsFirst   int64         `json:"is_first"`   // 是否第一次提现 1 否 2 是
		TicketNum int64         `json:"ticket_num"` // 现金券
		Goods     WithdrawGoods `json:"goods"`      // 商品
		Order     struct {
			OrderNo string `json:"order_no"` // 订单号
			OrderId int64  `json:"order_id"` // 订单 id
		} `json:"order"`
	}

	WithdrawInnerOrderResponse struct {
		WithdrawResponse
		Data []struct {
			ID          int    `json:"id"`           // 订单 id
			OrderNo     string `json:"order_no"`     // 订单号
			UID         int64  `json:"uid"`          // 用户id
			Amount      int    `json:"amount"`       // 金额
			Status      int    `json:"status"`       // 状态
			GoodsType   string `json:"goods_type"`   // 类型
			Bank        string `json:"bank"`         // 提现方式
			BankAccount string `json:"bank_account"` // 提现账户
			Name        string `json:"name"`         // 账户名称
			Phone       string `json:"phone"`        // 手机号
			IDCard      string `json:"id_card"`      // 卡号
			Tickets     string `json:"tickets"`      // 券
			Comment     string `json:"comment"`      // 内容
			AppName     string `json:"app_name"`     // app 名称
			CreatedAt   string `json:"created_at"`
			UpdatedAt   string `json:"updated_at"`
		} `json:"data"`
	}
	ParamWithdrawDeal struct {
		GoodsId     int64  `json:"goods_id"`     // 商品 id
		Name        string `json:"name"`         // 商品名称
		Phone       string `json:"phone"`        // 下单手机号
		IdCard      string `json:"id_card"`      // 银行卡
		Bank        string `json:"bank"`         // 银行
		BankAccount string `json:"bank_account"` // 账户
		Device      struct {
			IsRoot        bool `json:"is_root"`         // 手机设备
			IsXposedExist bool `json:"is_xposed_exist"` // 手机设备
			IsEmulator    bool `json:"is_emulator"`     // 手机设备
		} `json:"device"`
	}

	SpecialWithdrawState struct {
		MaxCount   int    `json:"max_count"`
		DailyLimit int    `json:"daily_limit"`
		Desc       string `json:"desc"`
		Title      string `json:"title"`
		State      string `json:"state"`
	}

	EnergyRewardRecord struct {
		RewardID     string `json:"reward_id"`
		RewardType   string `json:"reward_type"`
		RewardAmount int    `json:"reward_amount"`
	}

	EnergyRewardInfo struct {
		RewardID      string `json:"reward_id"`
		RequireEnergy int    `json:"require_energy"`
		Desc          string `json:"desc"`
		State         string `json:"state"`
	}

	ParamIncentiveEnergyComplete struct {
		PeriodID string `json:"period_id"`
		RewardID string `json:"reward_id"`
	}

	ParamIdiomLevelInfo struct {
		LevelID int `json:"level_id"`
	}
	IdiomConf struct {
		Lvl  int           `json:"lvl"`
		Conf IdiomConfInfo `json:"conf"`
	}

	IdiomConfInfo struct {
		Barrier []interface{} `json:"barrier"`
		ID      int           `json:"id"`
		Word    []string      `json:"word"`
		Idiom   []string      `json:"idiom"`
		Posx    []int         `json:"posx"`
		Posy    []int         `json:"posy"`
		Answer  []int         `json:"answer"`
	}

	IdiomAchievement struct {
		ID           int    `json:"id"`
		RequireLevel int    `json:"require_level"`
		State        string `json:"state"`
		RewardType   string `json:"reward_type"`
		RewardAmount int    `json:"-"`
	}

	ParamIdiomReward struct {
		LevelID      int  `json:"level_id"`
		HasIncentive bool `json:"has_incentive"`
	}
	ParamIdiomAchievement struct {
		ID int `json:"id"`
	}
	ParamWxShareConfig struct {
		Scene string `json:"scene"`
	}

	ParamMealBenefitInfo struct {
		MealType int `json:"meal_type"`
	}

	ParamMillstone struct {
		MillstoneID int `json:"millstone_id"`
	}

	ParamOpenRedPacket struct {
		Id string `json:"id"`
	}
	ParamTimeBenefit struct {
		Type int
	}

	ParamIdiomInfo struct {
		AchiType int64 `json:"achi_type"`
	}

	ParamStepMoneyItem struct {
		ItemID string `json:"item_id"`
	}

	ParamStepMoneyRedeem struct {
		RedeemPoint int64 `json:"redeem_point"`
	}

	ParamWeatherInfo struct {
		Lng         float64   `json:"lng"`
		Lat         float64   `json:"lat"`
		CityCode    string    `json:"city_code"`
		TestGroup   string    // 实验组
		WeatherType string    // 查看天气的类型
		Time        time.Time // 客户端访问时间
		Uid         uint64
		ApiType     string // 华风 api 类型 用于拼path
	}

	ParamSignDo struct {
		Day string `json:"day"`
	}

	ParamTaskGet struct {
		Tag     string   `json:"tag"`
		TaskIDS []string `json:"task_ids"`
	}
	ParamTaskComplete struct {
		TaskId string `json:"task_id"`
	}

	ParamLogin struct {
		LoginType int64  `json:"login_type"` // 1 手机 2 微信
		Account   string `json:"account"`    // 账号名称
		Uid       uint64 `json:"uid"`        // 用户 id
	}

	ParamWeatherInfoResp struct {
		Hourly      []*weather_mgr.HourlyStyle   `json:"hourly"`       // 小时数据
		Daily       []*weather_mgr.DailyStyle    `json:"daily"`        // 日数据
		Realtime    *weather_mgr.TodayResp       `json:"realtime"`     // 当前信息
		WarningList *weather_mgr.WarningListResp `json:"warning_list"` // 当前信息
	}

	DailyStyle struct {
		Date    int64   `json:"date"`
		Max     string  `json:"temperature_max"`
		Min     string  `json:"temperature_min"`
		Skycon  string  `json:"skycon"`
		Sunset  string  `json:"sunset"`
		Sunrise string  `json:"sunrise"`
		Aqi     float64 `json:"aqi"`
	}

	HourlyStyle struct {
		Date        int64  `json:"date"`
		Temperature string `json:"temperature"`
		Skycon      string `json:"skycon"`
	}

	RealTime struct {
		ForecastKeypoint string `json:"forecast_keypoint"` // 生活指数预报的详细描述，可能为空
		RainDesc         string `json:"rain_desc"`         // 降水描述，可能为空
		AlertDesc        string `json:"alert_desc"`        // 预警详细文字描述
		WarmRemind       string `json:"warm_remind"`       // 提示
		Humidity         string `json:"humidity"`          // 湿度
		// Comfort          string `json:"comfort"`           // 舒适指数
		LifeSuggestion LifeSuggestion `json:"life_suggestion"`
		Date           int64          `json:"date"`
	}

	LifeSuggestion struct {
		// Dating    string `json:"dating"`    // 约会指数
		Dressing string `json:"dressing"` // 穿衣指数
		Fishing  string `json:"fishing"`  // 钓鱼指数
		Flu      string `json:"flu"`      // 感冒指数
		Sport    string `json:"sport"`    // 运动指数
		Uv       string `json:"uv"`       // 紫外线指数
		Travel   string `json:"travel"`   // 旅游指数
		Airing   string `json:"airing"`   // 晾晒
	}

	ParamAwardVideo struct {
		BaseCoin  int64  `json:"base_coin"`  // 原来奖励金币数量
		Multiple  int64  `json:"multiple"`   // 倍数
		AwardName string `json:"award_name"` // 奖励类型 (兼容老版本新字段)
		AwardType int64  `json:"award_type"` // 奖励类型 1 温差 2 倒计时 3 任务 4签到 5新人红包
		Uid       uint64 `json:"uid"`        // 用户 id
	}

	ParamAwardVideoResp struct {
		IsAwardWatchTimes bool   `json:"is_award_watch_times"`
		IsGuide           bool   `json:"is_guide"` // 判断是否跳引导页面
		HadCoin           int64  `json:"had_coin"`
		HadCash           string `json:"had_cash"`
		AwardCoin         int64  `json:"award_coins"`
	}

	ParamWarmRemindResp struct {
		Notice string `json:"notice"` // 提醒
	}

	ParamUserGiftInfoResp struct {
		Cash       string `json:"cash"`        // 奖励额度
		CanReward  bool   `json:"can_reward"`  // 是否可以领取
		RewardCash int    `json:"reward_cash"` // 奖励数额
		HadCoin    int64  `json:"had_coin"`
		HadCash    string `json:"had_cash"`
	}
	ParamLoginResp struct {
		OpenId       string      `json:"open_id"`        // 微信 openid
		IsWithdrew03 bool        `json:"is_withdrew_03"` // 是否提现过
		UserInfo     interface{} `json:"user_info"`
	}
)

type ConfigTask struct {
	TaskId        string `json:"task_id"`
	TaskType      string `json:"task_type"`
	MaxNum        int64  `json:"max_num"`
	CurrentNum    int64  `json:"current_num"`
	AwardType     string `json:"award_type"`
	AwardNum      int    `json:"award_num"`
	Status        int64  `json:"status"` // 1待完成 2待领取 3已完成
	Desc          string `json:"desc"`
	Title         string `json:"title"`
	Descv2        string `json:"descv2"`
	Weight        int    `json:"weight"`
	Button        string `json:"button"`
	MinAppVersion int    `json:"min_app_version"`
	CountDown     int64  `json:"count_down"` // 任务领取间隔倒计时
}

type WeatherChanData struct {
	Params *weather_mgr.WeatherReq
}
