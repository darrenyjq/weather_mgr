package main

import (
	"base/cootek/pgd/base/account"
	"base/internal/grpc"
	"cootek.com/elete/sdk"
	"log"
	//全局初始化配置
	_ "base/internal"
	"base/internal/crontab"
	"os"
	"os/signal"
	"syscall"
)


func main()  {
	go httpServer()
	crontab.RunCrontab()
	notify()
}

func httpServer()  {
	server := sdk.NewGrpcServer([]string{"cootek.pgd.base.account.proto"})
	account.RegisterAccountServer(server,&grpc.AccountService{})
	sdk.PublishGrpcServer(server)
}

func notify()  {
	//wait for exit
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
	for {
		sig := <-ch
		switch sig {
		case syscall.SIGTERM:
			log.Println("step 5: http server exit success by SIGTERM")
			os.Exit(0)
		case syscall.SIGINT:
			log.Println("step 5: http server exit success by SIGINT")
			os.Exit(0)
		case syscall.SIGUSR1:
			log.Println("step 5: http server exit success by SIGUSR1")
			os.Exit(0)
		}
	}
}
