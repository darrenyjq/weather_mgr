package service

import (
	"context"
	"cootek.com/elete/sdk"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"base/cootek/account/TokenService"
	"base/cootek/pgd/account"
	"base/pkg/xzap"
	"strings"
	"time"
)

type tokenService struct {
	oldURL	map[string]bool
}

func newTokenService() (*tokenService) {
	return &tokenService{
		oldURL:map[string]bool {
			"ap.ime.cootek.com": true,
			"eu.ime.cootek.com": true,
			"usa.ime.cootek.com": true,
			"zh-cn.ime.cootek.com": true,
			"ap02-ap.ime.cootek.com": true,
			"pre-usa.ime.cootek.com": true,
			"eu.mobhey.com": true,
			"ap.mobhey.com": true,
			"usa.mobhey.com": true,
			"ap02-ap.mobhey.com": true,
		},
	}
}

func (this tokenService) shouldVerifyByV3Token(c *gin.Context) bool {
    return strings.HasPrefix(c.Request.RequestURI, "/b/") && !this.oldURL[c.Request.Host]
}

func (this tokenService) AuthToken(c *gin.Context,authToken string) (ok bool, ant *account.Account) {
	ant = new(account.Account)

	//测试环境
	if viper.GetString("current_env") == "dev" {
		ant,ok = this.testToken(c,authToken)
		this.SetUserInfo(c,ant)
		return
	}

	//requestId,_ := c.Get("RequestId")
	if authToken == "" {
		return
	}

	tokenIn := TokenService.TokenIn{Token:authToken}
	if this.shouldVerifyByV3Token(c) {
		// println(c.Request.Host, "is v3")
		tokenIn.Version = "v3"
	}

	ctx := c.Request.Context()
	ctx,_ = context.WithTimeout(ctx,1*time.Second)
	client := TokenService.NewTokenServiceClient(sdk.GrpcClientConnFromContext(ctx))
	fullinfo,err := client.QueryFullInfo(ctx, &tokenIn)
	if err != nil {
		//log.Error("AuthToken", zap.Any("requestId",requestId),zap.String("token",authToken), zap.Error(err))
		return
	}
	//token是否合法
	if fullinfo.GetResult() != TokenService.TokenStatus_TOKEN_VALID {
		return
	}
	ok = true
	ant = this.toAccount(c, fullinfo)
	this.SetUserInfo(c,ant)
	return
}

func (this tokenService) toAccount(c *gin.Context, fullinfo *TokenService.FullTokenInfo ) (*account.Account) {
	user_id := fullinfo.GetUserId()
	account_type := fullinfo.GetAccountType()
	token	:=	fullinfo.GetToken()
	exp_date	:=	fullinfo.GetExpDate()
	activation_code := fullinfo.GetActivationCode()
	account_name := fullinfo.GetAccountName()
	auth_info_3p := fullinfo.GetAuthInfo_3P()
	auth_3p_name := fullinfo.GetAuth_3PName()
	account_region := fullinfo.GetAccountRegion()
	forbidden :=	fullinfo.GetIsEzalterParticipant()
	media_source := fullinfo.GetMediaSource()

	ActivationInfo	:= fullinfo.GetActivationInfo()
	activate_type := ActivationInfo.GetActivateType()
	act_user_id	:=	ActivationInfo.GetUserId()
	app_name	:=	ActivationInfo.GetAppName()
	app_version	:=	ActivationInfo.GetAppVersion()
	os_name	:=	ActivationInfo.GetOsName()
	os_version	:=	ActivationInfo.GetOsVersion()
	device_info	:=	ActivationInfo.GetDeviceInfo()
	channel_code	:=	ActivationInfo.GetChannelCode()
	imei	:=	ActivationInfo.GetImei()
	uuid	:=	ActivationInfo.GetUuid()
	simid	:=	ActivationInfo.GetSimid()
	locale	:=	ActivationInfo.GetLocale()
	mnc	:=	ActivationInfo.GetMnc()
	manufacturer	:=	ActivationInfo.GetManufacturer()
	api_level	:=	ActivationInfo.GetApiLevel()
	host_app_name	:=	ActivationInfo.GetHostAppName()
	host_app_version	:=	ActivationInfo.GetHostAppVersion()
	resolution	:=	ActivationInfo.GetResolution()
	dpi	:=	ActivationInfo.GetDpi()
	pysical_size	:=	ActivationInfo.GetPysicalSize()
	recommend_channel	:=	ActivationInfo.GetRecommendChannel()
	last_activation_code	:=	ActivationInfo.GetLastActivationCode()
	identifier	:=	ActivationInfo.GetIdentifier()
	sys_app	:=	ActivationInfo.GetSysApp()
	timestamp	:=	ActivationInfo.GetTimestamp()
	idfa	:=	ActivationInfo.GetIdfa()
	idfv	:=	ActivationInfo.GetIdfv()
	apple_token	:=	ActivationInfo.GetAppleToken()
	release	:=	ActivationInfo.GetRelease()
	android_id	:=	ActivationInfo.GetAndroidId()
	mac_address	:=	ActivationInfo.GetMacAddress()
	device_verified	:=	ActivationInfo.GetDeviceVerified()
	package_name	:=	ActivationInfo.GetPackageName()
	meta_data	:=	ActivationInfo.GetMetaData()
	upgrade_timestamp	:=	ActivationInfo.GetUpgradeTimestamp()
	activate_timestamp	:=	ActivationInfo.GetActivateTimestamp()
	gaid	:=	ActivationInfo.GetGaid()


	accountType := account.AccountType(account_type)
	accountRegion := account_region.String()
	mediaSource := account.MediaSource(media_source)
	activateType := account.ActivationInfo_ActivateType(activate_type)

	ant := &account.Account{
		UserId:&user_id,
		AccountType:&accountType,
		Token:&token,
		ExpDate:&exp_date,
		ActivationCode:&activation_code,
		ActivationInfo:&account.ActivationInfo{
			ActivateType:&activateType,
			UserId:&act_user_id,
			AppName:&app_name,
			AppVersion:&app_version,
			OsName:&os_name,
			OsVersion:&os_version,
			DeviceInfo:&device_info,
			ChannelCode:&channel_code,
			Imei:&imei,
			Uuid:&uuid,
			Simid:&simid,
			Locale:&locale,
			Mnc:&mnc,
			Manufacturer:&manufacturer,
			ApiLevel:&api_level,
			HostAppName:&host_app_name,
			HostAppVersion:&host_app_version,
			Resolution:&resolution,
			Dpi:&dpi,
			PysicalSize:&pysical_size,
			RecommendChannel:&recommend_channel,
			LastActivationCode:&last_activation_code,
			Identifier:&identifier,
			SysApp:&sys_app,
			Timestamp:&timestamp,
			Idfa:&idfa,
			Idfv:&idfv,
			AppleToken:&apple_token,
			Release:&release,
			AndroidId:&android_id,
			MacAddress:&mac_address,
			DeviceVerified:&device_verified,
			PackageName:&package_name,
			MetaData:&meta_data,
			UpgradeTimestamp:&upgrade_timestamp,
			ActivateTimestamp:&activate_timestamp,
			Gaid:&gaid,
		},
		AccountName:&account_name,
		AuthInfo_3P:&auth_info_3p,
		Auth_3PName:&auth_3p_name,
		Forbidden:&forbidden,
		AccountRegion:&accountRegion,
		MediaSource:&mediaSource,
	}

	return ant
}

func (this tokenService) testToken(c *gin.Context,authToken string) (ant *account.Account,ok bool) {
	ant =  new(account.Account)
	data := map[string]string{
		"120-1-SVXN9tsJuLbHLVfC7PrU":`{"user_id":1,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money","app_version":"120"}}`,
		"120-2-yfqCgfayKJdnERxIgr4W":`{"user_id":2,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money","app_version":"120"}}`,
		"cn01:d37d44cd-a3cd-4a06-b3a9-3069bc123883":`{"user_id":3,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money","app_version":"120"}}`,
		"120-4-wRJP8sWpZmBs6srBsBFH":`{"user_id":4,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money","app_version":"120"}}`,

		"125-5-CNSjoCfJiys2vdqjOg31":`{"user_id":5,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money","app_version":"125"}}`,
		"125-6-3OmdFekxtEG1mDqYULmh":`{"user_id":6,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money","app_version":"125"}}`,
		"125-7-e9bYv8WgngSG3D77zt9V":`{"user_id":7,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money","app_version":"125"}}`,
		"125-8-b0qfHuHMGX2np6T9Nb3A":`{"user_id":8,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money","app_version":"125"}}`,

		"125-9-O7Ds6wXUs7aVzdt22js6":`{"user_id":9,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money","app_version":"125"}}`,
		"125-10-69evLywDJatiCxU0OHdz":`{"user_id":10,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money","app_version":"125"}}`,
		"125-11-ycJjsLAS7O6D0XUbbAuN":`{"user_id":11,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money","app_version":"125"}}`,
		"125-12-9dksIK2DMGx63vAZfvEV":`{"user_id":12,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money","app_version":"125"}}`,

		"131-13-vLywDJatiksIK2DM":`{"user_id":13,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money","app_version":"131"}}`,
		"131-14-ssdesfdtiksIK2DM":`{"user_id":14,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money","app_version":"131"}}`,

		"133-15-jfslkjflsljvewwv":`{"user_id":15,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money","app_version":"133"}}`,
		"133-16-ysiMkjhHKNHKjjkua":`{"user_id":16,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money","app_version":"133"}}`,
		"133-17-MudkckejljslfkHjk":`{"user_id":17,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money","app_version":"133"}}`,

		"161-18-MudkckejljslfkHjk":`{"user_id":18,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money","app_version":"161"}}`,
		"171-19-095adc37a323d452c":`{"user_id":19,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money","app_version":"171"}}`,

		"172-20-095adc37a323d452c":`{"user_id":20,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money.ios","app_version":"172","os_name": "iOS"}}`,
		"172-21-095adc37a323d452c":`{"user_id":21,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money.ios","app_version":"172","os_name": "iOS"}}`,
		"172-22-095adc37a323d452c":`{"user_id":22,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money.ios","app_version":"172","os_name": "iOS"}}`,
		"172-23-095adc37a323d452c":`{"user_id":23,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money.ios","app_version":"172","os_name": "iOS"}}`,

		"187-24-095adc37a323d452c":`{"user_id":24,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money","app_version":"187"}}`,

		"1024-25-095adc37a323d452c":`{"user_id":25,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money","app_version":"1024"}}`,
		"1024-26-095adc37a323d452c":`{"user_id":26,"account_type":1,"activation_info":{"app_name":"com.androidhealth.steps.money","app_version":"1024","os_name": "iOS"}}`,
	}
	s,ok := data[authToken]
	if !ok {
		return
	}
	err := json.Unmarshal([]byte(s),ant)
	if err != nil {
		xzap.ErrorContext(c,"testToken",zap.Error(err))
		return
	}
	return
}

func (this tokenService) SetUserInfo(c *gin.Context, ant *account.Account) {
	c.Set("account",ant)
	c.Set("app_name",ant.ActivationInfo.GetAppName())
	c.Set("version",ant.ActivationInfo.GetAppVersion())
	c.Set("uid",int64(ant.GetUserId()))
}