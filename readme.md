## 业务基础框架
1. 把代码clone下来，重命名为项目名称
```
git clone git@gitlab.corp.cootek.com:pgd_portfolio/mig/server/base_go.git newProject
```
2. 替换代码中的包引用路径，base为项目名即可
```
go run artisan.go -name newProject
```

# 依赖
项目依赖elete_go项目，该项目因为把整个gopath下的目录都上传到仓库中了，所以把该项目clone下来，再改为本地的gopath路径，把当前项目再放到src目录下。
这样就能找到依赖的elete_go中的sdk包。

```$xslt
git clone git@gitlab.corp.cootek.com:cloud_infra/elete_go.git go

cd go/src

git clone git@gitlab.corp.cootek.com:gb_server/mig2.git mig

cd mig

//编译
./build
```

# 编译
直接执行命令即可，会拉取master分支最新代码，并在本地build，产生的可执行文件放到docker中。
再把docker推到生产环境仓库，在titan中选择最新的镜像发布上线。

```$xslt
./build
```

# 本地开发配置
```
配置环境变量
export CUSTOM_RUNTIME_ENV=dev
export CUSTOM_RUNTIME_ENV=test   # 都可以，test不用验证token
export LOCATION=cn
export GO111MODULE=on 

创建连接
sudo ln -s /Users/yaodongen/work/go  /go

增加日志文件
mkdir /ssd
sudo chmod 777 /ssd/


```

## 本地搭建可执行环境

```
1.本地安装redis
brew install redis

2.启动redis
redis-server

3.设置环境变量
export CUSTOM_RUNTIME_ENV=dev
export LOCATION=cn

4.创建日志目录
mkdir /ssd && chmod 777 /ssd

5.创建项目目录
mkdir -p ~/go/src/mig
ln -s ~/go /go

6.把main,config,template目录放到~/go/src/mig目录下
后面更新，直接从这步开始，替换main

7.给main加上可执行权限
chmod +x main

8.启动
./main

9.如果要清数据
redis-cli

FLUSHALL
```