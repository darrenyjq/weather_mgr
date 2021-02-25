export LOCATION=cn
export CUSTOM_RUNTIME_ENV=DEV
export CONF_SERVER=not-conf-server
export ELETE_PROXY_HOST=elete-proxy-dev.cootekos.com:1921
export SERVICE=cootek.mig_one.weather_mgr
# 框架日志环境变量
export STDERR_REDIRECT=0
export APOLLO_ACCESSKEY_SECRET=ed56da201d8d4575bc1cd2b233018184
export BIND_DIRECT_IP="10.0.42.181"
go run main.go
