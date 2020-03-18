package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"base/pkg/xaes"
)

var (
	apiRouteHandle			*apiRoute
	middlewareHandle		*MiddlewareHandle
	aesHandle				*xaes.Aes
)

func init()  {
	apiRouteHandle = new(apiRoute)
	middlewareHandle = new(MiddlewareHandle)

	aesHandle = xaes.NewAes(viper.GetString("app.aesKey"))
	return
}

func RegisterRoute(g *gin.Engine)  {
	g.Use(middlewareHandle.SupportOptionsMethod(), middlewareHandle.RequestIdMiddleware())

	baseGroup := g.Group("/walkingformoney/withdraw")

	//内网服务
	innerRouteGroup := baseGroup.Group("/inner")
	//没有任何过滤条件
	{
		innerRouteGroup.GET("/welcome",apiRouteHandle.Welcome)
	}
	innerRouteGroup.Use(middlewareHandle.InnerNetworkAuthUserInfo())
	{
		innerRouteGroup.GET("/user-info",apiRouteHandle.Welcome)
	}

	//*****************************外网服务(别把内外网的路由顺序调换)*********************//
	//没有任何过滤条件
	openRouteGroup := baseGroup.Group("/open")
	{
		openRouteGroup.GET("/welcome",apiRouteHandle.Welcome)
	}
	//增加token验证和解密
	openRouteGroup.Use(middlewareHandle.AuthTokenMiddleware(),middlewareHandle.DecryptParamsMiddleware())
	{
		openRouteGroup.POST("/user-info",apiRouteHandle.Welcome)
	}


	return
}