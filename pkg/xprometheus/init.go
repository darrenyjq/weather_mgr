package xprometheus

import (
	"weather_mgr/helper"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	//请求
	CoinLabelCouter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "coin_bank",
			Namespace: "ime_one",
			Name:      "coin_label_count",
			Help:      "coin_bank coins transfer ",
		},
		[]string{"label", "app_name", "trade"},
	)
	CoinLabelTimesCouter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "coin_bank",
			Namespace: "ime_one",
			Name:      "coin_label_times",
			Help:      "coin_bank coins transfer times",
		},
		[]string{"label", "app_name", "trade"},
	)

	CoinAwardCouter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "coin_bank",
			Namespace: "ime_one",
			Name:      "coin_award",
			Help:      "coin_bank coin_award",
		},
		[]string{"label", "app_name", "error_code"},
	)
)

func init() {
	//注册
	prometheus.MustRegister(CoinLabelCouter, CoinLabelTimesCouter, CoinAwardCouter)

}
func AddCoinAwardCouter(label string, appName string, errCode int64) {
	CoinAwardCouter.With(
		prometheus.Labels{
			"label":      label,
			"app_name":   appName,
			"error_code": helper.Int642Str(errCode),
		}).Inc()
}

func AddCoinLabelCouter(label string, appName string, coinNum int64) {
	trade := "award"
	if coinNum < 0 {
		coinNum = -coinNum
		trade = "consume"
	}

	CoinLabelCouter.With(
		prometheus.Labels{
			"label":    label,
			"app_name": appName,
			"trade":    trade,
		}).Add(float64(coinNum))

	CoinLabelTimesCouter.With(
		prometheus.Labels{
			"label":    label,
			"app_name": appName,
			"trade":    trade,
		}).Inc()
}
