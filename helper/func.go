package helper

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"weather_mgr/bbbb/pgd/account"
	"weather_mgr/bbbb/pgd/ysession"
	"weather_mgr/pkg/xzap"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

// GetAccount 获取用户对象
func GetAccount(ctx context.Context) (userInfo *account.Account, exist bool) {
	userInfo = new(account.Account)
	ant := ctx.Value("account")
	if ant == nil {
		return
	}
	userInfo = ant.(*account.Account)
	exist = true
	return
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

func GetAccountUid(ctx context.Context) (uid uint64) {

	val := ctx.Value("uid")
	if val != nil {
		uid, ok := val.(uint64)
		if !ok {
			panic("uid 类型不正确")
		}
		return uid
	}

	panic("uid 不存在")
}

func Str2int64(str string) (int64, error) {
	i64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return i64, nil
}

// GetOSName 获取系统类型-小写-android/ios
func GetOSName(ctx context.Context) (OSName string) {
	OSName = "android"
	val := ctx.Value("os")
	if val != nil {
		strVal, ok := val.(string)
		if ok && strVal != "" {
			OSName = strings.ToLower(strVal)
		}
	}
	return
}

func GetRequestId(ctx context.Context) (requestId string) {
	val := ctx.Value("X-Request-Id")
	if val != nil {
		requestId, _ = val.(string)
	}
	return
}

func GetAuthToken(c *gin.Context) (authToken string) {
	authToken, _ = c.Cookie("auth_token")
	if authToken == "" {
		authToken = c.GetHeader("Auth-Token")
		if authToken == "" {
			authToken = c.Query("auth_token")
		}
	}
	return
}

func GetTsMsNow(ctx context.Context) uint64 {
	ms := time.Now().UnixNano() / 1e6
	if os.Getenv("CUSTOM_RUNTIME_ENV") != "pro" {
		millisecVal := ctx.Value("X-TS-Millisec")
		if millisecVal != nil {
			ms_str, _ := millisecVal.(string)
			_ms, err := strconv.ParseInt(ms_str, 10, 64)
			if err == nil {
				ms = _ms
			}
		}
	}
	return uint64(ms)
}

func GetTimeNow(ctx context.Context) time.Time {
	xzap.Debug("", zap.Any("CUSTOM_RUNTIME_ENV", os.Getenv("CUSTOM_RUNTIME_ENV")))
	if os.Getenv("CUSTOM_RUNTIME_ENV") != "pro" {
		millisecVal := ctx.Value("X-TS-Millisec")
		if millisecVal != nil {
			xzap.Debug("", zap.Any("millisecVal", millisecVal))
			millisec, _ := millisecVal.(string)
			userTime, err := MsToTime(millisec)
			if err == nil {
				return userTime
			}
		}
	}
	return time.Now()
}

func MsToTime(ms string) (time.Time, error) {
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	tm := time.Unix(0, msInt*int64(time.Millisecond))
	return tm, nil
}

func Abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

func MillisecToTime(ms int64) time.Time {
	tm := time.Unix(0, ms*int64(time.Millisecond))
	return tm
}
func JsonStrConnect(data map[string]interface{}) (s string) {
	keys := make([]string, 0, len(data))
	for k, _ := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	//fmt.Println(keys)

	for _, k := range keys {
		v := data[k]
		switch v.(type) {
		case map[string]interface{}:
			//对象值
			s += fmt.Sprintf("%s={%s}&", k, JsonStrConnect(v.(map[string]interface{})))
		case []interface{}:
			//数组值
			arrStr := ""
			for _, sonV := range v.([]interface{}) {
				switch sonV.(type) {
				case map[string]interface{}:
					//数组里的值是对象值
					arrStr += fmt.Sprintf("{%s},", JsonStrConnect(sonV.(map[string]interface{})))
				default:
					//数组里的值是单元素值.[1,2]
					arrStr += fmt.Sprintf("%v,", sonV)
				}
			}
			arrStr = strings.TrimRight(arrStr, ",")
			s += fmt.Sprintf("%s=[%s]&", k, arrStr)
		case nil:
			//排除掉空值 null
		case string:
			sv := v.(string)
			//排除掉空字符串
			if sv != "" {
				s += fmt.Sprintf("%s=%s&", k, sv)
			}
		default:
			//其他的直接相连接
			s += fmt.Sprintf("%s=%v&", k, v)
		}
	}
	s = strings.TrimRight(s, "&")
	return
}

func GetSeqId(sessionParam *ysession.SessionParam) string {
	id := sessionParam.GetCurSpanId()
	id += 1
	sessionParam.CurSpanId = id
	perID := sessionParam.GetPreSpanId()
	seqID := fmt.Sprintf("%s.%d", perID, id)
	return seqID
}

func Str2int(str string) (int, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func Str2Uint64(str string) (uint64, error) {
	i64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64(i64), nil
}

func Int642Str(num int64) string {
	return strconv.FormatInt(num, 10)
}

func Int2Str(num int) string {
	return strconv.Itoa(num)
}

func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MinInt64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MaxInt64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
