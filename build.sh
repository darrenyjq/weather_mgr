#!/bin/sh
set -xe

# 在本地生成生产环境可执行文件
# 把可执行添加到镜像 dockerfile
# 生成镜像推上去
tag=`date +%Y%m%d_%H%M_``git rev-parse --short HEAD`

docker build --no-cache -t harbor.cootekservice.com/ime_us/mig_coin_bank -t harbor.cootekservice.com/ime_us/mig_coin_bank:$tag .

docker push harbor.cootekservice.com/ime_us/mig_coin_bank:latest
docker push harbor.cootekservice.com/ime_us/mig_coin_bank:$tag
docker rmi harbor.cootekservice.com/ime_us/mig_coin_bank:$tag

