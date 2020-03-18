package service

import (
	"bytes"
	"context"
	"crypto/tls"
	"html/template"
	"base/pkg/xzap"

	"github.com/go-mail/mail"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type emailService struct {
	dialer mail.Dialer
}

func newEmailService() (email *emailService) {
	d := mail.Dialer{
		Host:           "smtp.partner.outlook.cn",
		Port:           587,
		Username:       "noreply.ime@cootek.cn",
		Password:       `IIWMgic%B9ua`,
		StartTLSPolicy: mail.MandatoryStartTLS,
	}
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	email = &emailService{
		dialer: d,
	}
	return
}

func (this emailService) SendWithdrawReport(ctx context.Context) {
	/*title := "提现服务"
	templateName := "withdraw_report.html"
	data := ReportServ.WithdrawReport(ctx)
	this.sendMail(ctx, data, title, templateName)*/
}


func (this emailService) sendMail(ctx context.Context, data interface{}, title, templateName string) {
	temp := template.New("test")
	temp, err := template.ParseFiles("/go/src/withdraw/template/" + templateName)
	if err != nil {
		xzap.ErrorContext(ctx, "sendMail", zap.Error(err))
		return
	}

	b := bytes.NewBuffer([]byte{})
	err = temp.Execute(b, data)
	if err != nil {
		xzap.ErrorContext(ctx, "sendMail", zap.Error(err))
		return
	}

	m := mail.NewMessage()
	m.SetHeader("From", "noreply.ime@cootek.cn")
	if viper.GetString("current_env") == "dev" {
		m.SetHeader("To", "kangkang.han@cootek.cn")
		m.SetHeader("Subject", "(测试)"+title)
	} else {
		m.SetHeader("To", "kangkang.han@cootek.cn","elvin.zheng@cootek.cn","fengshun.fan@cootek.cn")
		//m.SetHeader("Cc", "joe.xie@cootek.cn", "jia.yin@cootek.cn")
		m.SetHeader("Subject", title)
	}
	m.SetBody("text/html", b.String())

	if err := this.dialer.DialAndSend(m); err != nil {
		xzap.ErrorContext(ctx, "sendMail", zap.Error(err))
	}
}
