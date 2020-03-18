package crontab

import (
	"base/helper"
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"os"
)

type crontab struct {
	cron  *cron.Cron
}

func RunCrontab()  {
	//生产环境只允许此主机上运行
	if viper.GetString("current_env") == "pro" {
		if os.Getenv("HOSTNAME") != "ime03.corp.cootek.com" {
			fmt.Println("crontab forbid! this host is: " + os.Getenv("HOSTNAME"))
			return
		}
	}

	//因为是部署到多个容器里的，只允许一个容器执行crontab任务
	if helper.GetInstNo() != "1" {
		return
	}

	ctx := context.WithValue(context.Background(),"X-Request-Id","crontab")
	crontab := new(crontab)
	crontab.cron = cron.New()

	crontab.WithdrawRequestCash(ctx)

	crontab.cron.Start()
	fmt.Println("run crontab success")
}

func (this *crontab) WithdrawRequestCash(ctx context.Context) {
	fmt.Println("run crontab WithdrawRequestCash")
	/*_, err := this.cron.AddFunc("00 11 * * ?", func() {
		service.WithdrawServ.CrontabCash(ctx)
	})
	if err != nil {
		panic(err)
	}*/
}
