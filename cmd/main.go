package main

import (
	"cootek.com/elete/sdk"
	"github.com/gin-gonic/gin"
	"log"
	//全局初始化配置
	_ "base/internal"
	"base/internal/api"
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
	g := gin.Default()
	api.RegisterRoute(g)
	log.Println("step 3: init gin route success")

	server := sdk.NewHttpServer([]string{"token_service.proto"})
	sdk.Handle("/",g)
	log.Fatal(server.ListenAndServe())
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
