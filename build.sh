#!/bin/sh
set -xe

# 在本地生成生产环境可执行文件
# 把可执行添加到镜像 dockerfile
# 生成镜像推上去
tag=`date +%Y%m%d_%H%M`

cd ~/go && git pull origin master

cd ~/go/src/base && git checkout master && git pull origin master

GOOS=linux GOARCH=amd64 go build -o main cmd/main.go

docker build --no-cache -t harbor.cootekservice.com/ime_us/base_api -t harbor.cootekservice.com/ime_us/base_api:$tag .
docker push harbor.cootekservice.com/ime_us/base_api:latest
docker push harbor.cootekservice.com/ime_us/base_api:$tag

rm -f main
