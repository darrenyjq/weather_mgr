
# 第二步: 把编译好的可执行文件复制到镜像内, 并运行
FROM harbor.cootekservice.com/library/ubuntu:18.04.v1
WORKDIR /coin_bank/
COPY main ./
ENTRYPOINT exec /coin_bank/main