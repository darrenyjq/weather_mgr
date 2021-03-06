package main

import (
	"context"
	"go.uber.org/zap"
	"log"
	"weather_mgr/bbbb/pgd/weather_mgr"
	"weather_mgr/pkg/xzap"

	"gitlab.corp.aaaa/cloud_infra/elete-go/pkg/elete/sdk"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"syscall"
	_ "weather_mgr/internal"
	"weather_mgr/internal/service"
)

func main() {
	go grpcServer()
	notify()
}

func grpcServer() {
	// 拦截器
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		xzap.Info("params:", zap.Any("", req))
		// 继续处理请求
		resp, err = handler(ctx, req)
		xzap.Info("resp:", zap.Any("", resp))

		// 转换grpc类型错误码
		// var errCode int64
		if err != nil {
			err = nil
			resp = &weather_mgr.TodayResp{
				AlertDesc: "该服务已下线",
			}
			// errCode = helper.GetErrCode(err)
			// err = status.Errorf(codes.Code(errCode), err.Error())
			// errCode = helper.GetErrCode(err)

		}
		// 上报数据
		return
	}
	// 添加拦截器
	sdk.SetupApplicationServerInterceptor(interceptor)
	// 注册要调用的下游服务，多个就调用多次
	sdk.AddNormalProtoFileToMetadatda("bbbb.pgd.ysession.proto")
	sdk.AddNormalProtoFileToMetadatda("account.proto")
	sdk.AddServiceProtoFileToMetadata("bbbb.pgd.weather_mgr.proto")
	server := sdk.NewGrpcServer(sdk.EMPTY_RROTOS)
	weather_mgr.RegisterWeatherMgrServer(server, service.WeatherServ)
	sdk.PublishGrpcServer(server)
}

func notify() {
	// wait for exit
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
