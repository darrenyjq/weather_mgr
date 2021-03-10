package helper

const (
	// 账户状态
	ACCOUNT_STATUS_OK     = 0
	ACCOUNT_STATUS_FORBID = 1

	INVITE_CODE_RULE         = "填写您邀请人的邀请码，邀请人可在我的页面查看到自己的邀请码，不可以自己填自己。每位用户限填一次，确认提交后不能修改。"
	INVITE_REWARD_COIN       = 5000
	INVITE_CODE_ENTER_REWARD = 888

	REWARD_INVITE_FRIEND = "INVITE_FRIEND"
	REWARD_ENTER_CODE    = "ENTER_CODE"

	CashWheelPrizeCoin      = 1 // 现金大转盘奖品 金币
	CashWheelPrizeRedPacket = 2 // 现金大转盘 奖品 红包
	CashWheelPrizePrizeGift = 3 // 现金大转盘奖品互动式广告

	CashWheelMillstoneNotCollectible = "not_collectible"
	CashWheelMillstoneCanCollect     = "collectible"
	CashWheelMillstoneCollected      = "collected"

	TIME_FORMAT_20060102 = "20060102"
)

// 根据 国标AQI 得到 空气质量等级
func GetAqiChn(aqi float64) string {
	switch {
	case aqi > 0 && aqi <= 50:
		return "优"
	case aqi > 50 && aqi <= 100:
		return "良"
	case aqi > 100 && aqi <= 150:
		return "轻度污染"
	case aqi > 150 && aqi <= 200:
		return "中度污染"
	case aqi > 200:
		return "重度污染"
	default:
		return "优"
	}
}

// 1 温差 2 倒计时 3 任务 4签到 5新人红包
func GetAwardType(awardType int64) string {
	switch awardType {
	case 1:
		return "TEMPERATURE_COIN"
	case 2:
		return "TIME_COIN"
	case 3:
		return "FULI_TASK"
	case 4:
		return "SIGN_IN"
	case 5:
		return "RED_PACKET_NEW_USER"
	default:
		return ""
	}
}

// 预警等级
func GetWarningLevel(level string) string {
	switch level {
	case "黄色":
		return "yellow"
	case "蓝色":
		return "blue"
	case "橙色":
		return "orange"
	case "红色":
		return "red"
	case "白色":
		return "white"
	default:
		return "yellow"
	}
}

// 预警标识
func GetWarningTypeName(typeName string) string {
	switch typeName {
	case "台风":
		return "typhoon "
	case "暴雨":
		return "rainstorm "
	case "暴雪":
		return "blizzard "
	case "寒潮":
		return "cold_wave"
	case "大风":
		return "strong_wind"
	case "沙尘暴":
		return "sandstorm"
	case "高温":
		return "high_temperature"
	case "干旱":
		return "drought"
	case "雷电":
		return "thunder_lightning"
	case "冰雹":
		return "hail"
	case "霜冻":
		return "frost"
	case "大雾":
		return "fog"
	case "霾":
		return "haze"
	case "道路结冰":
		return "icy_roads"
	case "寒冷":
		return "cold"
	case "灰霾":
		return "grey_haze"
	case "雷雨大风":
		return "the_thunderstorm_winds"
	case "森林火险":
		return "forest_fire"
	case "降温":
		return "cooling"
	case "道路冰雪":
		return "roads"
	case "干热风":
		return "hot_wind"
	case "低温":
		return "snow_ice"
	case "冰冻":
		return "frozen_at_low_temperature"
	case "空气重污染":
		return "heavy_air_pollution"
	case "海上大雾":
		return "sea_fog"
	case "雷暴大风":
		return "thunderstorm_winds"
	case "持续低温":
		return "continuous_low_temperature"
	case "浓浮尘":
		return "thick_dust"
	case "龙卷风":
		return "tornado"
	case "低温冻害":
		return "low_temperature_cold_injury"
	case "海上大风":
		return "sea_wind"
	case "低温雨雪冰冻":
		return "snow_frozen_at_low_temperature"
	case "强对流":
		return "strong_convection"
	case "臭氧":
		return "ozone"
	case "大雪":
		return "heavy_snow"
	case "强降雨":
		return "heavy_rain_strong_cooling"
	case "强降温":
		return "snow"
	case "雪灾":
		return "forest_fire "
	case "草原火险":
		return "thunderstorms"
	case "雷暴":
		return "strong_wind"
	case "严寒":
		return "cold"
	case "沙尘":
		return "dust"
	case "海上雷雨大风":
		return "sea"
	case "海上雷电":
		return "sea_thunder_lightning"
	case "海上台风":
		return "typhoon_at_sea"
	default:
		return "grey_haze"
	}
}
func GetWarmRemindNotice(t int64, level string) (str string) {
	// 以温度取文案
	switch {
	case level == "1" || t < 6:
		str = "天气寒冷，宜穿羽绒服，佩戴帽子和手套。"
	case level == "2" || t >= 6 && t < 11:
		str = "天气冷，宜穿棉衣、羽绒服、毛衣加大衣。"
	case level == "3" || t >= 11 && t < 15:
		str = "天气较冷，宜穿薄棉衣、羊毛衫加皮夹克。"
	case level == "4" || t >= 15 && t < 21:
		str = "天气微凉，宜穿夹克衫、羊毛衫加西服套装。"
	case level == "5" || t >= 21 && t < 24:
		str = "天气温暖，宜穿长袖衬衫、薄套装。"
	case level == "6" || t >= 24 && t < 28:
		str = "天气较热，宜穿棉麻衬衫、单裤。"
	case level == "7" || 999 > t && t >= 28:
		str = "天气炎热，宜穿薄T恤衫、敞领短袖棉衫。"
	default:
		str = "刷新温度有贴心穿衣提醒，不再穿错衣服"
	}
	return
}

const (
	// 签到状态
	// 待签到
	SIGN_IN_STATUS_WAIT_SIGN = 0
	// 待补签
	SIGN_IN_STATUS_WAIT_REMEDY_SIGN = 1
	// 签到成功
	SIGN_IN_STATUS_SUCCESS = 2
	// 禁止签到
	SIGN_IN_STATUS_FORBID = 3
)

const (
	// 任务状态
	// 待完成
	TASK_STATUS_WAIT = 1
	// 待领取
	TASK_STATUS_WAIT_READY = 2
	// 已领取
	TASK_STATUS_SUCCESS = 3
)

// 天气码对应标识
var WEATHER_SKY = map[string]string{
	"100": "CLEAR_DAY",               // 	晴
	"103": "PARTLY_CLOUDY_DAY",       // 	晴间多云
	"150": "CLEAR_NIGHT",             // 	晴
	"153": "PARTLY_CLOUDY_NIGHT",     // 	晴间多云
	"101": "CLOUDY",                  // 	多云
	"305": "LIGHT_RAIN",              // 	小雨
	"306": "MODERATE_RAIN",           // 	中雨
	"307": "HEAVY_RAIN",              // 	大雨
	"310": "STORM_RAIN",              // 	暴雨
	"400": "LIGHT_SNOW",              // 	小雪
	"401": "MODERATE_SNOW",           // 	中雪
	"402": "HEAVY_SNOW",              // 	大雪
	"403": "STORM_SNOW",              // 	暴雪
	"501": "FOG",                     // 	雾
	"502": "LIGHT_HAZE",              // 	霾
	"503": "SAND",                    // 	扬沙
	"504": "DUST",                    // 	浮尘
	"511": "MODERATE_HAZE",           // 	中度霾
	"512": "HEAVY_HAZE",              // 	重度霾
	"104": "OVERCAST",                // 阴
	"300": "SHOWER_RAIN",             // 阵雨
	"301": "HEAVY_SHOWER_RAIN",       // 强阵雨
	"406": "SHOWER_SNOW",             // 阵雨夹雪
	"407": "SNOW_FLURRY",             // 阵雪
	"154": "OVERCAST",                // 阴
	"350": "SHOWER_RAIN",             // 阵雨
	"351": "HEAVY_SHOWER_RAIN",       // 强阵雨
	"456": "SHOWER_SNOW",             // 阵雨夹雪
	"457": "SNOW_FLURRY",             // 阵雪
	"102": "FEW_CLOUDS",              // 少云
	"302": "THUNDERSHOWER",           // 雷阵雨
	"303": "HEAVY_THUNDERSTORM",      // 强雷阵雨
	"304": "THUNDERSHOWER_WITH_HAIL", // 雷阵雨伴有冰雹
	"308": "EXTREME_RAIN",            // 极端降雨
	"309": "DRIZZLE_RAIN",            // 毛毛雨
	"311": "HEAVY_STORM",             // 大暴雨
	"312": "SEVERE_STORM",            // 特大暴雨
	"313": "FREEZING_RAIN",           // 冻雨
	"314": "LIGHT_TO_MODERATE_RAIN",  // 小到中雨
	"315": "MODERATE_TO_HEAVY_RAIN",  // 中到大雨
	"316": "HEAVY_RAIN_TO_STORM",     // 大到暴雨
	"317": "STORM_TO_HEAVY_STORM",    // 暴雨到大暴雨
	"318": "HEAVY_TO_SEVERE_STORM",   // 大暴雨到特大暴雨
	"399": "RAIN",                    // 雨
	"404": "SLEET",                   // 雨夹雪
	"405": "RAIN_AND_SNOW",           // 雨雪天气
	"408": "LIGHT_TO_MODERATE_SNOW",  // 小到中雪
	"409": "MODERATE_TO_HEAVY_SNOW",  // 中到大雪
	"410": "HEAVY_SNOW_TO_SNOWSTORM", // 大到暴雪
	"499": "SNOW",                    // 雪
	"500": "MIST",                    // 薄雾
	"507": "DUSTSTORM",               // 沙尘暴
	"508": "SANDSTORM",               // 强沙尘暴
	"509": "DENSE_FOG",               // 浓雾
	"510": "STRONG_FOG",              // 强浓雾
	"513": "SEVERE_HAZE",             // 严重霾
	"514": "HEAVY_FOG",               // 大雾
	"515": "EXTRA_HEAVY_FOG",         // 特强浓雾
	"900": "HOT",                     // 热
	"901": "COLD",                    // 冷
	"999": "UNKNOWN",                 // 未知

	"0":  "CLEAR_DAY",               //	晴（国内城市白天晴）
	"1":  "CLEAR_DAY",               //	晴（国内城市夜晚晴）
	"2":  "CLEAR_DAY",               //	晴（国外城市白天晴）
	"3":  "CLEAR_DAY",               //	晴（国外城市夜晚晴）
	"4":  "CLOUDY",                  //	多云
	"5":  "PARTLY_CLOUDY_NIGHT",     //	晴间多云
	"6":  "PARTLY_CLOUDY_NIGHT",     //	晴间多云
	"7":  "CLOUDY",                  //	大部多云
	"8":  "CLOUDY",                  //	大部多云
	"9":  "OVERCAST",                //	阴
	"10": "SHOWER_RAIN",             //	阵雨
	"11": "THUNDERSHOWER",           //	雷阵雨
	"12": "THUNDERSHOWER_WITH_HAIL", //	雷阵雨伴有冰雹
	"13": "LIGHT_RAIN",              //	小雨
	"14": "MODERATE_RAIN",           //	中雨
	"15": "HEAVY_RAIN",              //	大雨
	"16": "STORM_RAIN",              //	暴雨
	"17": "HEAVY_STORM",             //	大暴雨
	"18": "SEVERE_STORM",            //	特大暴雨
	"19": "FREEZING_RAIN",           //	冻雨
	"20": "SLEET",                   //	雨夹雪
	"21": "SNOW_FLURRY",             //	阵雪
	"22": "LIGHT_SNOW",              //	小雪
	"23": "MODERATE_SNOW",           //	中雪
	"24": "HEAVY_SNOW",              //	大雪
	"25": "STORM_SNOW",              //	暴雪
	"26": "DUST",                    //	浮尘
	"27": "SAND",                    //	扬沙
	"28": "DUSTSTORM",               //	沙尘暴
	"29": "SANDSTORM",               //	强沙尘暴
	"30": "FOG",                     //	雾
	"31": "LIGHT_HAZE",              //	霾
	"32": "WIND",                    //	风
	"33": "BLUSTERY",                //	大风
	"34": "HURRICANCE",              //	飓风
	"35": "TROPICAL_STORM",          //	热带风暴
	"36": "TORNADO",                 //	龙卷风
	"37": "COLD",                    //	冷
	"38": "HOT",                     //	热
	"99": "UNKNOWN",                 //	未知
}

func GetWeatherSky(code string) string {
	if re, ok := WEATHER_SKY[code]; ok {
		return re
	}
	return "CLEAR_DAY"
}

func GetHuaFengWeatherSky(code string) string {
	if re, ok := Huafeng_SKY[code]; ok {
		return re
	}
	return "CLEAR_DAY"
}

func GetHuaFengHourlySky(code int) string {
	if re, ok := Huafeng_HOURLY_SKY[code]; ok {
		return re
	}
	return "CLEAR_DAY"
}

var Huafeng_SKY = map[string]string{
	"00":  "CLEAR_DAY",               // 晴
	"01":  "CLOUDY",                  // 多云
	"02":  "OVERCAST",                // 阴
	"03":  "SHOWER_RAIN",             // 阵雨
	"04":  "THUNDERSHOWER",           // 雷阵雨
	"05":  "THUNDERSHOWER_WITH_HAIL", // 雷阵雨伴有冰雹
	"06":  "SLEET",                   // 雨夹雪
	"07":  "LIGHT_RAIN",              // 小雨
	"08":  "MODERATE_RAIN",           // 中雨
	"09":  "HEAVY_RAIN",              // 大雨
	"10":  "STORM_RAIN",              // 暴雨
	"11":  "HEAVY_STORM",             // 大暴雨
	"12":  "SEVERE_STORM",            // 特大暴雨
	"13":  "SNOW_FLURRY",             // 阵雪
	"14":  "LIGHT_SNOW",              // 小雪
	"15":  "MODERATE_SNOW",           // 中雪
	"16":  "HEAVY_SNOW",              // 大雪
	"17":  "STORM_SNOW",              // 暴雪
	"18":  "FOG",                     // 雾
	"19":  "FREEZING_RAIN",           // 冻雨
	"20":  "DUSTSTORM",               // 沙尘暴
	"21":  "LIGHT_TO_MODERATE_RAIN",  // 小到中雨
	"22":  "MODERATE_TO_HEAVY_RAIN",  // 中到大雨
	"23":  "HEAVY_RAIN_TO_STORM",     // 大到暴雨
	"24":  "STORM_TO_HEAVY_STORM",    // 暴雨到大暴雨
	"25":  "HEAVY_TO_SEVERE_STORM",   // 大暴雨到特大暴雨
	"26":  "LIGHT_TO_MODERATE_SNOW",  // 小到中雪
	"27":  "MODERATE_TO_HEAVY_SNOW",  // 中到大雪
	"28":  "HEAVY_SNOW_TO_SNOWSTORM", // 大到暴雪
	"29":  "DUST",                    // 浮尘
	"30":  "SAND",                    // 扬沙
	"31":  "SANDSTORM",               // 强沙尘暴
	"32":  "DENSE_FOG",               // 浓雾
	"33":  "SNOW",                    // 雪
	"49":  "STRONG_FOG",              // 强浓雾
	"53":  "LIGHT_HAZE",              // 霾
	"54":  "MODERATE_HAZE",           // 中度霾
	"55":  "HEAVY_HAZE",              // 重度霾
	"56":  "SEVERE_HAZE",             // 严重霾
	"57":  "HEAVY_FOG",               // 大雾
	"58":  "EXTRA_HEAVY_FOG",         // 特强浓雾
	"301": "RAIN",                    // 降雨
	"302": "SNOW",                    // 降雪
}

var Huafeng_HOURLY_SKY = map[int]string{
	1:  "CLEAR_DAY",           // 晴
	2:  "CLEAR_DAY",           // 大部分晴
	3:  "CLEAR_DAY",           // 部分晴
	4:  "PARTLY_CLOUDY_DAY",   // 间歇性多云
	5:  "CLEAR_DAY",           // 晴，空气质量差
	6:  "PARTLY_CLOUDY_DAY",   // 大部分多云
	7:  "CLOUDY",              // 多云
	8:  "OVERCAST",            // , Overcast阴
	11: "FOG",                 // 雾
	12: "SHOWER_RAIN",         // 阵雨
	13: "SHOWER_RAIN",         // 多云，有时有阵雨
	14: "SHOWER_RAIN",         // Showers部分晴，有时有阵雨
	15: "THUNDERSHOWER",       // 雷雨
	16: "THUNDERSHOWER",       // 多云，有时有雷雨
	17: "THUNDERSHOWER",       // 部分晴，有时有雷雨
	18: "RAIN",                // 雨
	19: "LIGHT_SNOW",          // 轻雪
	20: "LIGHT_SNOW",          // 多云，有时有轻雪
	21: "LIGHT_SNOW",          // 部分晴，有时有轻雪
	22: "SNOW",                // 雪
	23: "SNOW",                // 多云，有时有雪
	24: "SNOW",                // 冻雪
	25: "SLEET",               // 冰霰
	26: "FREEZING_RAIN",       // 冻雨
	29: "SLEET",               // Snow雨和雪
	30: "HOT",                 // 热
	31: "COLD",                // 冷
	32: "WIND",                // 风
	33: "CLEAR_NIGHT",         // 晴
	34: "CLEAR_NIGHT",         // 大部分晴朗
	35: "PARTLY_CLOUDY_NIGHT", // Cloudy部分多云
	36: "PARTLY_CLOUDY_NIGHT", // Clouds间歇性多云
	37: "CLEAR_NIGHT",         // 晴，空气质量差
	38: "PARTLY_CLOUDY_NIGHT", // 大部分多云
	39: "SHOWER_RAIN",         // 部分多云，有时有阵雨
	40: "SHOWER_RAIN",         // 大部分多云，有时有阵雨
	41: "THUNDERSHOWER",       // 部分多云，有时有雷雨
	42: "THUNDERSHOWER",       // 大部分多云，有时有雷雨
	43: "SNOW",                // 大部分多云，有时有阵雪
	44: "SNOW",                // 大部分多云，有雪
}
