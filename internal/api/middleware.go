package api

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
	"base/cootek/pgd/account"
	"base/helper"
	"base/internal/service"
	"base/pkg/xprometheus"
	"base/pkg/xzap"
)

type MiddlewareHandle struct {}

func (MiddlewareHandle) SupportOptionsMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "OPTIONS" {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusOK)
		}
	}
}

func (MiddlewareHandle) RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqId := ""
		idObj := uuid.NewV4()
		if idObj.String() == "" {
			reqId = fmt.Sprintf("%d",time.Now().UnixNano())
		}else{
			reqId = idObj.String()
		}
		c.Set("X-Request-Id",reqId)
		c.Writer.Header().Set("X-Request-Id", reqId)

		c.Next()

		//prometheus上报
		code := strconv.Itoa(c.Writer.Status())
		path := strings.Split(c.Request.RequestURI, "?")[0]
		errorCode := fmt.Sprintf("%d",c.GetInt64("error_code"))
		go xprometheus.AddHttpRequestsCouter(code,path,errorCode,c.GetString("version"),c.GetString("app_name"))
	}
}

func (MiddlewareHandle) InnerNetworkAuthUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			xzap.ErrorContext(c,"InnerNetworkAuthUserInfo",zap.Error(err))
		}else {
			//再赋值回去
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			//获取post json里的account信息
			ant := new(helper.ParamUserInfo)
			err = c.BindJSON(ant)
			if err != nil {
				FailedResponse(c,helper.ERROR_USERINFO)
				return
			}
			//设置用户信息到context
			service.TokenServ.SetUserInfo(c,ant.Account)
			//再赋值回去，用于下面业务获取参数
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
		c.Next()
	}
}

func (MiddlewareHandle) AuthTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := helper.GetAuthToken(c)
		ok,ant := service.TokenServ.AuthToken(c,authToken)
		if !ok {
			FailedResponse(c,helper.ERR_TOKEN)
			c.Abort()
			return
		}
		//未登录
		if ant.GetAccountType() == account.AccountType_TEMP {
			FailedResponse(c,helper.ERROR_CODE_NOT_LOGIN_USER)
			c.Abort()
			return
		}
		c.Next()
	}
}

func (MiddlewareHandle) DecryptParamsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取加密的参数
		encryptParam := new(helper.ParamEncryptData)
		err := c.BindJSON(encryptParam)
		if err != nil {
			FailedResponse(c,helper.ERROR_SECURITY_PARAM)
			c.Abort()
			return
		}

		if encryptParam.Data != "" {
			//获取解密后的json
			originData,err := aesHandle.Decrypt(encryptParam.Data)
			if err != nil {
				FailedResponse(c,helper.ERROR_SECURITY_DECRYPT_ERR)
				c.Abort()
				return
			}
			//把解密后的数据，放到body里，控制器里直接拿来用
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(originData))
		}

		c.Next()
	}
}
