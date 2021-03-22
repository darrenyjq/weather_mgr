# 第一步: go build 可执行文件
#FROM harbor.cootekservice.com/library/golang:1.13.5 as builder
## 工作目录和项目一致, 把incentives替换为你自己的项目
#WORKDIR /weather_mgr
## 把项目目录下的所有文件复制到镜像内
##COPY . ./
### 执行 go build
##RUN ls && go env -w GO111MODULE="on" && go env -w GOPROXY="http://source.cootekos.com/goproxy/" && go env -w GONOSUMDB="gitlab.corp.cootek.com" && \
##    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go
#
## 第二步: 把编译好的可执行文件复制到镜像内, 并运行
#FROM harbor.cootekservice.com/library/ubuntu:18.04.v1
#WORKDIR /weather_mgr/
#COPY --from=builder /weather_mgr/main ./
#ENTRYPOINT exec /weather_mgr/main >> /ssd/weather_mgr-${INST_NO}-$(date +%Y%m%d%H).log 2>&1
#
FROM alpine

RUN mkdir /ssd && mkdir -p /src/weather_mgr/configs

COPY main /src/weather_mgr
ENTRYPOINT exec /src/weather_mgr/main >> /ssd/weather_mgr-${INST_NO}-$(date +%Y%m%d%H).log 2>&1
