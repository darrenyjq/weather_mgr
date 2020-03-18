package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

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
	g.Use(func(c *gin.Context) {
		//只在测试环境加，生产环境在Nginx里
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Server,Date,Content-Length,Cache-Control,Keep-Alive,Connection,X-Requested-With,X-File-Name,Origin,Accept,Auth-Token,X-TZ-UTC,X-TS-Millisec")
		c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		c.Header("Access-Control-Max-Age", "1728000")
	})
	api.RegisterRoute(g)
	log.Println("step 3: test init gin route success")

	log.Fatal(http.ListenAndServe("0.0.0.0:9527",g))
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
