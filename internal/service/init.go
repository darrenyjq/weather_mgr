package service

import (
	"crypto/tls"
	"net/http"
	"weather_mgr/pkg/xrsa"
)

var (
	httpClient  *http.Client
	rsaHandle   *xrsa.XRsa
	WeatherServ *WeatherService
)

func init() {
	WeatherServ = new(WeatherService)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient = &http.Client{Transport: tr}
}
