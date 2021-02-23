#!/bin/sh
set -xe

# 在本地生成生产环境可执行文件
# 把可执行添加到镜像 dockerfile
# 生成镜像推上去
user_name=`git config user.name`
tag=`date +%Y%m%d_%H%M_``git rev-parse --short HEAD`${user_name// /}

GOOS=linux GOARCH=amd64 go build -o main cmd/main.go
docker build --no-cache -t harbor.cootekservice.com/ime_us/mig_coin_bank -t harbor.cootekservice.com/ime_us/mig_coin_bank:$tag -f test.Dockerfile .

docker push harbor.cootekservice.com/ime_us/mig_coin_bank:latest
docker push harbor.cootekservice.com/ime_us/mig_coin_bank:$tag
docker rmi harbor.cootekservice.com/ime_us/mig_coin_bank:$tag

