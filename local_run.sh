if [[ $1 == "docker" ]]; then
  #  取本地最新镜像 TAG 存入环境变量
  tag=$(docker images | grep "weather_mgr" | awk '{print $2}' | head -n 1)
  export TAG=$tag
  #  以 docker 启动当前镜像服务
  docker-compose down
  docker-compose up -d --build
else
  export LOCATION=cn
  export CUSTOM_RUNTIME_ENV=DEV
  export CONF_SERVER=not-conf-server
  export ELETE_PROXY_HOST=elete-proxy-dev.bbbbos.com:1921
  export SERVICE=bbbb.mig_one.weather_mgr
  # 框架日志环境变量
  export STDERR_REDIRECT=0
  export APOLLO_ACCESSKEY_SECRET=ed56da201d8d4575bc1cd2b233018184
  export BIND_DIRECT_IP="10.0.42.181"
  go run main.go
fi
