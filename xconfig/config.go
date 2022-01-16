package xconfig

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
	"weather_mgr/bbbb/pgd/config_service"
	"weather_mgr/pkg/xzap"

	"gitlab.corp.aaaa/cloud_infra/elete-go/pkg/elete/sdk"
	"go.uber.org/zap"
)

type configServiceResponse struct {
	Data struct {
		Config string `json:"config"`
	} `json:"data"`
	ErrorCode int64  `json:"error_code"`
	Msg       string `json:"msg"`
}

type config struct {
	data    sync.Map
	setTime sync.Map
	host    string
}

func (this *config) getConfig(ctx context.Context, configKey string, group int64, appName string, cache bool) (value []byte, err error) {
	//不指定group，直接请求配置服务来计算
	if group == 0 || cache {
		return this.readConfig(ctx, configKey, group, appName, cache)
	}

	//指定了group,并且不用被配置服务缓存
	cacheKey := fmt.Sprintf("%s-%d-%s", configKey, group, appName)
	//在缓存里存在
	cacheVal, ok := this.data.Load(cacheKey)
	if ok {
		//先赋值
		value = cacheVal.([]byte)
		//超出缓存时间
		st, ok := this.setTime.Load(cacheKey)
		if ok && (time.Now().Unix()-st.(int64)) > 60 {
			//获取原始配置，第n次，有错误忽略，别影响业务服务
			tempNewValue, err2 := this.readConfig(ctx, configKey, group, appName, false)
			if err2 != nil {
				return
			}
			value = tempNewValue
			//缓存
			this.data.Store(cacheKey, value)
			this.setTime.Store(cacheKey, time.Now().Unix())
		}
	} else {
		//获取原始配置，基本上第一次获取，有错误直接返回
		value, err = this.readConfig(ctx, configKey, group, appName, false)
		if err != nil {
			return
		}
		//缓存
		this.data.Store(cacheKey, value)
		this.setTime.Store(cacheKey, time.Now().Unix())
	}

	return
}

func (this *config) readConfig(ctx context.Context, key string, group int64, appName string, cache bool) (val []byte, err error) {
	if appName == "" {
		err = errors.New("get config but app name is empty")
		xzap.ErrorContext(ctx, "readConfig", zap.Error(err), zap.String("key", key))
		return
	}
	in := &config_service.GetConfigParam{
		Key:     key,
		Group:   group,
		AppName: appName,
		Cache:   cache,
	}

	//xzap.InfoContext(ctx,"readConfig",zap.Error(err),zap.Any("param",in))

	val, err = this.readFromGrpc(ctx, in)
	if err != nil {
		val, err = this.readFromHttp(ctx, in)
	}

	return
}

func (this *config) readFromGrpc(ctx context.Context, in *config_service.GetConfigParam) (val []byte, err error) {
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	client := config_service.NewConfigServiceClient(sdk.GrpcClientConnFromContext(ctx))
	result, err := client.Get(ctx, in)
	if err != nil {
		xzap.ErrorContext(ctx, "readConfig", zap.Error(err), zap.Any("param", in))
		return
	}
	val = []byte(result.GetConfig())
	return
}

func (this *config) readFromHttp(ctx context.Context, in *config_service.GetConfigParam) (val []byte, err error) {
	jb, err := json.Marshal(in)
	if err != nil {
		xzap.ErrorContext(ctx, "readFromHttp", zap.Error(err), zap.Any("param", in))
		return
	}
	path := this.host + "/coin_bankformoney/config/inner/get"
	resp, err := http.Post(path, "application/json", bytes.NewReader(jb))
	if err != nil {
		xzap.ErrorContext(ctx, "readFromHttp", zap.Error(err), zap.Any("param", in))
		return
	}
	defer resp.Body.Close()

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		xzap.ErrorContext(ctx, "readFromHttp", zap.Error(err), zap.Any("param", in))
		return
	}

	response := new(configServiceResponse)
	err = json.Unmarshal(body, response)
	if err != nil {
		xzap.ErrorContext(ctx, "readFromHttp", zap.Error(err), zap.Any("param", in), zap.Any("resp", body))
		return
	}

	if response.ErrorCode != 0 {
		err = fmt.Errorf("err_code:%d,msg:%s", response.ErrorCode, response.Msg)
		xzap.ErrorContext(ctx, "readFromHttp", zap.Error(err), zap.Any("param", in))
		return
	}
	val = []byte(response.Data.Config)
	return
}

var cfg config

func init() {
	host := "http://mig-config-service.corp.aaaa"
	if strings.ToLower(os.Getenv("CUSTOM_RUNTIME_ENV")) != "pro" {
		host = "http://pgd-beta.bbbbservice.com"
	}
	log.Println("init xconfig, http host: " + host)
	cfg = config{
		host: host,
	}
}

//初始化全局变量
//配置格式：配置名，json，abtest版本(1，2，3....)
//用到的时候去取配置，并缓存下来
//定时更新
