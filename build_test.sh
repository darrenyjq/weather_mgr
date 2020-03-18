#!/bin/sh
set -xe

# 在本地生成生产环境可执行文件
# 把可执行添加到镜像 dockerfile
# 生成镜像推上去
tag=`date +%Y%m%d_%H%M`

cd ~/go && git pull origin master

cd ~/go/src/base

GOOS=linux GOARCH=amd64 go build -o main cmd/main.go

docker build --no-cache -t harbor.cootekservice.com/ime_us/base_api_test -t harbor.cootekservice.com/ime_us/base_api_test:$tag .
docker push harbor.cootekservice.com/ime_us/base_api_test:latest
docker push harbor.cootekservice.com/ime_us/base_api_test:$tag

rm -f main
