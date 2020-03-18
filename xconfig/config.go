package xconfig

import (
	"base/pkg/xzap"
	"bytes"
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"log"
	"base/xconfig/model"
	"strconv"
	"strings"
	"sync"
	"time"
)

type config struct {
	data  sync.Map
}

func (this *config) getConfig(ctx context.Context, ckey configKey) (value []byte,err error) {
	cacheKey := ckey.GetKeyGroupOSNameString(ctx)
	cacheVal,ok := this.data.Load(cacheKey)
	if ok {
		value = cacheVal.([]byte)
		return
	}

	//获取原始配置
	value,err = this.readConfig(ctx, ckey.GetKey(ctx), ckey.GetGroup(ctx), ckey.GetAppName(ctx))
	if err != nil {
		return
	}
	//缓存
	this.data.Store(cacheKey,value)
	return
}

//定时循环更新配置
func (this *config)  updateConfig(ctx context.Context)  {
	tick := time.Tick(10*time.Second)
	for {
		select {
		case <-tick:
			//fmt.Println("tick range")
			this.data.Range(func(key, value interface{}) bool {
				cacheKey	:= key.(string)
				oldJsonByteValue := value.([]byte)

				//找到其原型
				tmpKs := strings.Split(cacheKey,"-")
				originKey := tmpKs[0]
				ckey,ok := configMap[originKey]
				if !ok {
					return true
				}
				originGroup,err := strconv.ParseInt(tmpKs[1],10,64)
				if err != nil {
					xzap.Error("updateConfigError",zap.Error(err))
					return true
				}
				originAppName := tmpKs[2]

				//获取最新的配置内容,group用名字里的，因为abtest
				newJsonByteValue,err := this.readConfig(ctx, originKey, originGroup, originAppName)
				if err != nil {
					return true
				}
				//相同就不更新
				if bytes.Equal(newJsonByteValue,oldJsonByteValue) {
					return true
				}
				//检查新的json值是否合法,保证修改错误不影响现在的正确缓存
				objStruct := ckey.GetObjectStruct()
				err = json.Unmarshal(newJsonByteValue, &objStruct)
				if err != nil {
					xzap.Error("updateConfigError",zap.Error(err),zap.String("key",cacheKey),zap.ByteString("byte",newJsonByteValue))
					return true
				}

				//缓存
				this.data.Store(cacheKey, newJsonByteValue)
				//fmt.Println(key, string(newJsonByteValue))
				return true
			})
		}
	}
}


func (this *config) readConfig(ctx context.Context, key string, group int64, appName string) (val []byte,err error) {
	//fmt.Println(key,group,osName)
	//刚开始从数据库读取
	val,err = model.ConfigModel.GetConfig(key, group, appName)
	return
}

var cfg config

func init()  {
	log.Println("init xconfig")
	cfg = config{}
	go cfg.updateConfig(context.Background())
}


//初始化全局变量
//配置格式：配置名，json，abtest版本(1，2，3....)
//用到的时候去取配置，并缓存下来
//定时更新








