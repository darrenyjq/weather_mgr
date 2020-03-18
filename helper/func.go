package helper

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"base/pkg/xzap"
)

func Decimal(value float64) float64 {
	value, err := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	if err != nil {
		xzap.Error("Decimal", zap.Error(err))
	}
	return value
}

func GetInstNo() string {
	instNo := os.Getenv("INST_NO")
	if instNo == "" {
		instNo = "1"
	}
	return instNo
}

func GetAppVersion(ctx context.Context) (version string) {
	val := ctx.Value("version")
	if val != nil {
		version, _ = val.(string)
	}
	return
}


func GetAppName(ctx context.Context) (appName string) {
	val := ctx.Value("app_name")
	if val != nil {
		appName, _ = val.(string)
	}
	return
}

func GetAccountUid(ctx context.Context) (uid int64)  {
	val := ctx.Value("uid")
	if val != nil {
		uid, _ = val.(int64)
	}
	return
}

func GetRequestId(ctx context.Context) (requestId string) {
	val := ctx.Value("X-Request-Id")
	if val != nil {
		requestId,_ = val.(string)
	}
	return
}

func GetAuthToken(c *gin.Context) (authToken string)  {
	authToken,_ = c.Cookie("auth_token")
	if authToken == "" {
		authToken = c.GetHeader("Auth-Token")
		if authToken == "" {
			authToken = c.Query("auth_token")
		}
	}
	return
}

func GetToday(ctx context.Context) (string) {
	return time.Now().Format("20060102")
}

func GetToday2(ctx context.Context) (string) {
	return time.Now().Format("2006-01-02")
}

func Abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

func JsonStrConnect(data map[string]interface{}) (s string) {
	keys := make([]string, 0, len(data))
	for k,_ := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	//fmt.Println(keys)

	for _,k := range keys {
		v := data[k]
		switch v.(type) {
		case map[string]interface{}:
			//对象值
			s += fmt.Sprintf("%s={%s}&",k,JsonStrConnect(v.(map[string]interface{})))
		case []interface{}:
			//数组值
			arrStr := ""
			for _,sonV := range v.([]interface{}) {
				switch sonV.(type) {
				case map[string]interface{}:
					//数组里的值是对象值
					arrStr += fmt.Sprintf("{%s},",JsonStrConnect(sonV.(map[string]interface{})))
				default:
					//数组里的值是单元素值.[1,2]
					arrStr += fmt.Sprintf("%v,",sonV)
				}
			}
			arrStr = strings.TrimRight(arrStr,",")
			s += fmt.Sprintf("%s=[%s]&",k,arrStr)
		case nil:
			//排除掉空值 null
		case string:
			sv := v.(string)
			//排除掉空字符串
			if sv != "" {
				s += fmt.Sprintf("%s=%s&",k,sv)
			}
		default:
			//其他的直接相连接
			s += fmt.Sprintf("%s=%v&",k,v)
		}
	}
	s = strings.TrimRight(s,"&")
	return
}
