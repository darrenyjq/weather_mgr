package service

import (
	"base/pkg/xrsa"
	"crypto/tls"
	"math/rand"
	"net/http"
	"time"
)

var (
	TokenServ *tokenService
	EmailServ		*emailService

	httpClient *http.Client
	rsaHandle	*xrsa.XRsa
)

func init()  {
	TokenServ = newTokenService()
	EmailServ	=	newEmailService()

	rand.Seed(time.Now().UnixNano())

	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	httpClient = &http.Client{Transport: tr}
}