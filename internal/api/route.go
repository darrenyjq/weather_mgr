package api

import (
	"net/http"
	"time"
	"weather_mgr/helper"

	"github.com/gin-gonic/gin"
)

func FailedResponse(c *gin.Context, err error) {
	appErr, ok := err.(helper.APP_ERROR)
	if !ok {
		c.Set("error_code", int64(77001))
		resp := gin.H{"error_code": 77001, "msg": "服务出错", "debug_message": err.Error(), "time_unix": time.Now().Unix()}
		c.JSON(http.StatusInternalServerError, resp)
	} else {
		c.Set("error_code", appErr.GetCode())
		resp := gin.H{"error_code": appErr.GetCode(), "msg": appErr.Error(), "time_unix": time.Now().Unix()}
		c.JSON(appErr.GetHttpCode(), resp)
	}
}

func SuccessResponse(c *gin.Context, data interface{}) {
	resp := map[string]interface{}{
		"error_code": 0,
		"msg":        "success",
		"time_unix":  time.Now().Unix(),
		"data":       data,
	}
	c.JSON(http.StatusOK, resp)
}

func StatusNotModifiedResponse(c *gin.Context) {
	c.String(http.StatusNotModified, "")
}
