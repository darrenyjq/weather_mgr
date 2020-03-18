package helper

import "net/http"

type APP_ERROR struct {
	Code int64
	Text string
	HttpCode int
}

func (this APP_ERROR) GetCode() (int64) {
	return this.Code
}

func (this APP_ERROR) Error() (string) {
	return this.Text
}

func (this APP_ERROR) GetHttpCode() (int) {
	return this.HttpCode
}

func NewAppError(code int64, text string, httpCode int) error {
	return APP_ERROR{Code: code, Text: text, HttpCode: httpCode}
}

var (
	ERR_INPUT	=	APP_ERROR{1000,"输入错误", http.StatusOK}
	ERR_TOKEN	=	APP_ERROR{1001,"token失效",http.StatusOK}
	ERROR_CODE_NOT_LOGIN_USER	=	APP_ERROR{1108,"用户未登录",http.StatusOK}
	ERROR_CODE_NOT_EXIST_USER	=	APP_ERROR{1109,"用户不存在",http.StatusOK}

	ERROR_SECURITY_REQ_REPLAY	=	APP_ERROR{2001,"接口重放",http.StatusOK}
	ERROR_SECURITY_SIGN_ERR	=	APP_ERROR{2002,"签名验证失败",http.StatusOK}
	ERROR_SECURITY_DECRYPT_ERR	=	APP_ERROR{2003,"数据解密失败",http.StatusOK}
	ERROR_SECURITY_TIME_DIFF	=	APP_ERROR{2004,"客户端和服务器时间不一致",http.StatusOK}
	ERROR_SECURITY_PARAM	=	APP_ERROR{2004,"加密参数格式解析失败",http.StatusOK}

	ERROR_USERINFO	=	APP_ERROR{3000,"用户信息错误",http.StatusOK}

	FAILED	=	APP_ERROR{77001,"服务器异常",http.StatusInternalServerError}
)