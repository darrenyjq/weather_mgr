package xprometheus

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	//请求
	HttpRequestsCouter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "mig_withdraw",
			Namespace: "pgd",
			Name:      "http_requests_total",
			Help:      "mig_withdraw http request counter",
		},
		[]string{"code", "path", "error_code", "version", "app_name"},
	)

	//用户事件
	UserEventCouter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "mig_withdraw",
			Namespace: "pgd",
			Name:      "user_event_total",
			Help:      "mig_withdraw user event counter",
		},
		[]string{"event", "thing", "num", "version", "app_name"},
	)


)

func init()  {
	//注册
	prometheus.MustRegister(HttpRequestsCouter, UserEventCouter)
}

func AddHttpRequestsCouter(httpCode, path, errorCode, version, appName string)  {
	HttpRequestsCouter.With(prometheus.Labels{"code": httpCode, "path": path, "error_code":errorCode, "version":version, "app_name":appName}).Inc()
}

func AddUserEventCouter(event , thing string, num int64, version , appName string)  {
	UserEventCouter.With(prometheus.Labels{"event":event, "thing":thing, "num":fmt.Sprintf("%d",num), "version":version, "app_name":appName}).Inc()
}



