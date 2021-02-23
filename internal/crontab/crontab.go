package crontab

import (
	"context"
	"fmt"

	"github.com/robfig/cron/v3"
)

type crontab struct {
	cron *cron.Cron
}

func RunCrontab() {
	//生产环境只允许此主机上运行

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
