FROM alpine

ENV CONF_ELETE_HTTP_SERVER_PORT ime__mig_base_api
ENV JAEGER_SERVICE_NAME ime__mig_base_api

RUN mkdir /ssd && mkdir -p /go/src/base/configs && mkdir -p /go/src/base/template

COPY configs /go/src/base/configs
COPY template   /go/src/base/template
COPY main /go/src/base

CMD ["/go/src/base/main"]
