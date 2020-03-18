package api

import (
	"github.com/stretchr/testify/assert"
	"base/helper"
	"base/pkg/xaes"

	"github.com/gavv/httpexpect/v2"
	"net/http"
	"testing"
)

var testurl ="http://127.0.0.1/walkingformoney/withdraw"

func TestHttpGetWelcomePass(t *testing.T) {
	e:= httpexpect.New(t, testurl)  //创建一个httpexpect实例
	e.GET("/open/welcome").   //ge请求
		Expect().
		Status(http.StatusOK).
		Text().
		Contains("welcome")
}

func TestHttpPostInfoPass(t *testing.T) {
	postdata:= map[string]interface{}{   //创建一个json变量
		"data":"",
		"sign_type":"aes",
	}
	path := "/open/info"
	fields := []string{"amount","goods","rules","ticket_num","uid"}

	response := testRequest(t,path,postdata).Path("$.data")
	object := response.Object()
	for _,v := range fields {
		object = object.ContainsKey(v)
	}
}

func TestHttpPostOrdersPass(t *testing.T)  {
	postdata:= map[string]interface{}{   //创建一个json变量
		"data":"",
		"sign_type":"aes",
	}
	path := "/open/orders"
	fields := []string{"amount","bank","bank_account","comment","created_at","goods_type","id","id_card","name","order_no","phone","status","tickets","uid","updated_at"}

	response := testRequest(t,path,postdata).Path("$.data")
	object := response.Array().First().Object()
	for _,v := range fields {
		object = object.ContainsKey(v)
	}
}

func TestHttpPostTicketsPass(t *testing.T)  {
	postdata:= map[string]interface{}{   //创建一个json变量
		"data":"",
		"sign_type":"aes",
	}
	path := "/open/tickets"
	fields := []string{"event","expired","status","ticket_id","uid"}

	response := testRequest(t,path,postdata).Path("$.data")
	object := response.Array().First().Object()
	for _,v := range fields {
		object = object.ContainsKey(v)
	}
}

//一个商品只能提现一次
func TestHttpPostDealFailed_3009(t *testing.T)  {
	aes := xaes.NewAes("dea5a27becaa2e054b959b3b7c8c95034949e4ab")
	data,err := aes.Encrypt([]byte(`{"goods_id":1,"bank":"alipay","phone":"15611111111","name":"曾繁然","id_card":"330401199208018142","bank_account":"1234567@qq.com","device":{"is_root":true,"is_emulator":false,"is_xposed_exist":false}}`))
	assert.NoError(t,err)

	postdata:= map[string]interface{}{   //创建一个json变量
		"data":data,
		"sign_type":"aes",
	}
	path := "/open/deal"
	testRequest(t,path,postdata).
		Object().
		ContainsKey("error_code").
		ValueEqual("error_code",helper.ERROR_GOODS_WITHDRAW_LIMIT.GetCode())
}

//提现金额不足
func TestHttpPostDealFailed_3005(t *testing.T)  {
	aes := xaes.NewAes("dea5a27becaa2e054b959b3b7c8c95034949e4ab")
	data,err := aes.Encrypt([]byte(`{"goods_id":2,"bank":"alipay","phone":"15611111111","name":"曾繁然","id_card":"330401199208018142","bank_account":"1234567@qq.com","device":{"is_root":true,"is_emulator":false,"is_xposed_exist":false}}`))
	assert.NoError(t,err)

	postdata:= map[string]interface{}{   //创建一个json变量
		"data":data,
		"sign_type":"aes",
	}
	path := "/open/deal"
	testRequest(t,path,postdata).
		Object().
		ContainsKey("error_code").
		ValueEqual("error_code",helper.ERROR_AMOUNT_NOT_ENOUGH.GetCode())
}

//用户信息
func TestHttpPostAccountInfoPass(t *testing.T) {
	postdata:= map[string]interface{}{   //创建一个json变量
		"account":map[string]interface{}{
			"user_id": 5874802553,
			"account_type": 0,
			"activation_info": map[string]interface{}{
				"app_name": "com.androidhealth.steps.money",
				"app_version": "120",
			},
		},
	}
	path := "/inner/account/info"
	fields := []string{"amount","ticket_num"}

	response := testRequest(t,path,postdata).Path("$.data")
	object := response.Object()
	for _,v := range fields {
		object = object.ContainsKey(v)
	}
}

//增加金额成功
func TestHttpPostIncrAmountPass(t *testing.T) {
	postdata:= map[string]interface{}{   //创建一个json变量
		"account":map[string]interface{}{
			"user_id": 5874802553,
			"account_type": 0,
			"activation_info": map[string]interface{}{
				"app_name": "com.androidhealth.steps.money",
				"app_version": "120",
			},
		},
		"incr_num":20,
		"event":"act_red_envelope_rain",
	}
	path := "/inner/account/incr-amount"
	fields := []string{"amount"}

	response := testRequest(t,path,postdata).Path("$.data")
	object := response.Object()
	for _,v := range fields {
		object = object.ContainsKey(v)
	}
}

//金额操作，超过增加限制
func TestHttpPostIncrAmountFailed_3002(t *testing.T) {
	postdata:= map[string]interface{}{   //创建一个json变量
		"account":map[string]interface{}{
			"user_id": 5874802553,
			"account_type": 0,
			"activation_info": map[string]interface{}{
				"app_name": "com.androidhealth.steps.money",
				"app_version": "120",
			},
		},
		"incr_num":20,
		"event":"test",
	}
	path := "/inner/account/incr-amount"
	testRequest(t,path,postdata).
		Object().
		ContainsKey("error_code").
		ValueEqual("error_code",helper.ERROR_AMOUNT_EVENT_NOT_EXIST.GetCode())
}

//金额操作，超过增加限制
func TestHttpPostIncrAmountFailed_3003(t *testing.T) {
	postdata:= map[string]interface{}{   //创建一个json变量
		"account":map[string]interface{}{
			"user_id": 5874802553,
			"account_type": 0,
			"activation_info": map[string]interface{}{
				"app_name": "com.androidhealth.steps.money",
				"app_version": "120",
			},
		},
		"incr_num":300,
		"event":"act_red_envelope_rain",
	}
	path := "/inner/account/incr-amount"

	testRequest(t,path,postdata).
		Object().
		ContainsKey("error_code").
		ValueEqual("error_code",helper.ERROR_AMOUNT_EVENT_LIMIT.GetCode())
}

//增加提现券成功
func TestHttpPostIncrTicketPass(t *testing.T) {
	postdata:= map[string]interface{}{   //创建一个json变量
		"account":map[string]interface{}{
			"user_id": 5874802553,
			"account_type": 0,
			"activation_info": map[string]interface{}{
				"app_name": "com.androidhealth.steps.money",
				"app_version": "120",
			},
		},
		"incr_num":1,
		"event":"checkin_day_reward",
	}
	path := "/inner/incr-ticket"
	fields := []string{"ticket_num"}

	response := testRequest(t,path,postdata).Path("$.data")
	object := response.Object()
	for _,v := range fields {
		object = object.ContainsKey(v)
	}
}


func testRequest(t *testing.T,path string, postdata interface{}) (response *httpexpect.Value) {
	response = httpexpect.New(t, testurl).
		POST(path).   //ge请求
		WithHeader("Auth-Token","cn01:d37d44cd-a3cd-4a06-b3a9-3069bc123883").
		WithHeader("ContentType", "application/json;charset=utf-8"). //定义头信息
		WithJSON(postdata).  //传入json body
		Expect().
		Status(http.StatusOK).   //判断请求是否200
		JSON()
	return
}








