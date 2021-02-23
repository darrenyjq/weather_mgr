package service

import (
	"crypto/tls"
	"math/rand"
	"net/http"
	"time"
	"weather_mgr/pkg/xrsa"
)

var (
	httpClient *http.Client
	rsaHandle  *xrsa.XRsa
)
var (
	TokenServ *tokenService
	// CoinServ  *CoinService
	WeatherServ *WeatherService
	// LogServ     *LogService
)

func init() {
	TokenServ = newTokenService()
	// CoinServ = new(CoinService)
	WeatherServ = new(WeatherService)
	// LogServ = new(LogService)
	rand.Seed(time.Now().UnixNano())

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient = &http.Client{Transport: tr}
}
