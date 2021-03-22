package helper

import "net/http"

type APP_ERROR struct {
	Code     int64
	Text     string
	HttpCode int
}

func (this APP_ERROR) GetCode() int64 {
	return this.Code
}

func (this APP_ERROR) Error() string {
	return this.Text
}

func (this APP_ERROR) GetHttpCode() int {
	return this.HttpCode
}

func GetErrCode(err error) (errCode int64) {
	if err != nil {
		appErr, ok := err.(APP_ERROR)
		if ok {
			errCode = appErr.GetCode()
		} else {
			errCode = FAILED.GetCode()
		}
	}
	return
}
func NewAppError(code int64, text string, httpCode int) error {
	return APP_ERROR{Code: code, Text: text, HttpCode: httpCode}
}

var (
	ERR_INPUT                 = APP_ERROR{1000, "输入错误", http.StatusOK}
	ERR_TOKEN                 = APP_ERROR{1001, "token失效", http.StatusOK}
	ERROR_CODE_NOT_LOGIN_USER = APP_ERROR{1108, "用户未登录", http.StatusOK}
	ERROR_CODE_NOT_EXIST_USER = APP_ERROR{1109, "用户不存在", http.StatusOK}

	ERROR_SECURITY_REQ_REPLAY  = APP_ERROR{2001, "接口重放", http.StatusOK}
	ERROR_SECURITY_SIGN_ERR    = APP_ERROR{2002, "签名验证失败", http.StatusOK}
	ERROR_SECURITY_DECRYPT_ERR = APP_ERROR{2003, "数据解密失败", http.StatusOK}
	ERROR_SECURITY_TIME_DIFF   = APP_ERROR{2004, "客户端和服务器时间不一致", http.StatusOK}
	ERROR_SECURITY_PARAM       = APP_ERROR{2004, "加密参数格式解析失败", http.StatusOK}
	PANIC                      = APP_ERROR{1200, "panic错误", http.StatusInternalServerError}

	ERROR_USERINFO = APP_ERROR{3000, "用户信息错误", http.StatusOK}

	FAILED = APP_ERROR{77001, "服务器异常", http.StatusInternalServerError}

	ERROR_DATA                   = APP_ERROR{78001, "数据错误", http.StatusInternalServerError}
	ERROR_OVER_LIMIT             = APP_ERROR{74006, "超过限制", http.StatusOK}
	ERROR_NOT_FOUND              = APP_ERROR{74002, "未找到记录", http.StatusOK}
	ERROR_CONFLICT               = APP_ERROR{74003, "记录冲突", http.StatusOK}
	ERROR_ILEAGAL                = APP_ERROR{74004, "非法数据", http.StatusOK}
	ERROR_CONCURRENT_CONFLICT    = APP_ERROR{78002, "有正在进行的操作", http.StatusOK}
	ERROR_TOUCH_LIMIT            = APP_ERROR{78100, "请勿频繁请求", http.StatusOK}
	ERROR_NOT_CURRECT_TIME       = APP_ERROR{78003, "未在指定时间", http.StatusOK}
	ERROR_ALREADY_EXIST          = APP_ERROR{78004, "已经存在", http.StatusOK}
	ERROR_REPEATED               = APP_ERROR{78005, "重复行为", http.StatusOK}
	ERROR_TOO_FREQUENT           = APP_ERROR{78006, "频率过高", http.StatusOK}
	ERROR_NOT_TARGET_USER        = APP_ERROR{78007, "未符合条件", http.StatusOK}
	ERROR_EXCEED_LIMIT           = APP_ERROR{78008, "超过最大限制", http.StatusOK}
	ERROR_CASH                   = APP_ERROR{78009, "现金操作失败", http.StatusOK}
	ERROR_INVITE                 = APP_ERROR{78010, "邀请失败", http.StatusOK}
	ERROR_ALREADY_INVITED        = APP_ERROR{78011, "已经被邀请过了", http.StatusOK}
	ERROR_NOT_MATCH              = APP_ERROR{78012, "未满足条件", http.StatusOK}
	ERROR_ALREDAY_COLLECT        = APP_ERROR{78013, "已经领取过了", http.StatusOK}
	ERROR_REACH_MAX_LEVEL        = APP_ERROR{78014, "已到达最大关卡", http.StatusOK}
	ERROR_TIME                   = APP_ERROR{78015, "时间异常", http.StatusInternalServerError}
	ERROR_PRIZE_NOT_FOUND        = APP_ERROR{78016, "未找到对应奖励", http.StatusInternalServerError}
	ERROR_ACTION_INSUFFICIENT    = APP_ERROR{78017, "体力不足", http.StatusOK}
	ERROR_RATE_LIMIT             = APP_ERROR{78018, "提现用户太多，请稍后重试", http.StatusOK}
	ERROR_COIN_ERROR             = APP_ERROR{78019, "内部错误，请稍后重试", http.StatusOK}
	ERROR_INNER_NET_ERROR        = APP_ERROR{78020, "内网错误", http.StatusOK}
	ERROR_NOT_ENOUGH             = APP_ERROR{78021, "数量不足", http.StatusOK}
	ERROR_ALREADY_END            = APP_ERROR{78022, "活动已经结束", http.StatusOK}
	ERROR_EXPIRE                 = APP_ERROR{78023, "已过期", http.StatusOK}
	ERROR_GIFT_NOT_FOUND         = APP_ERROR{78032, "未找到新手礼记录", http.StatusInternalServerError}
	ERROR_GOODS_WITHDRAW_LIMIT   = APP_ERROR{3009, "提现次数超额", http.StatusOK}
	ERROR_AMOUNT_NOT_ENOUGH      = APP_ERROR{3005, "金额额度不够", http.StatusOK}
	ERROR_AMOUNT_EVENT_NOT_EXIST = APP_ERROR{3002, "金额操作事件不存在", http.StatusOK}
	ERROR_AMOUNT_EVENT_LIMIT     = APP_ERROR{3003, "金额操作事件达最大额度", http.StatusOK}
	ERROR_TODAY_SIGNIN           = APP_ERROR{3000, "今日已签到", http.StatusOK}
	ERROR_TASK_NUM               = APP_ERROR{3001, "次数已达到最大值", http.StatusOK}
	ERROR_NOTICE_API             = APP_ERROR{3011, "API系统维护升级", http.StatusOK}

	ERROR_STATUS_ERR    = APP_ERROR{1110, "状态错误", http.StatusOK}
	ERR_CODE_NOT_FOUND  = APP_ERROR{74002, "资源找不到", http.StatusOK}
	ERR_NOT_FOUND       = APP_ERROR{74002, "资源不存在", http.StatusOK}
	ERROR_INVALID       = APP_ERROR{74107, "请求无合法", http.StatusOK}
	ERROR_REPEATE_PHONE = APP_ERROR{3010, "手机号已被其他账户绑定", http.StatusOK}
	ERROR_REPEATE_WX    = APP_ERROR{3010, "微信号已被其他账户绑定", http.StatusOK}
)

func NewMsgError(err APP_ERROR, msg string) APP_ERROR {
	return APP_ERROR{err.Code, msg, err.HttpCode}
}
