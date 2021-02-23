FROM golang

RUN ln -sf /bin/bash /bin/sh && echo "set -o history">>~/.bashrc && apt-get update && apt-get install apt-file -y && apt-file update && apt-get install vim -y
ENV CUSTOM_RUNTIME_ENV=dev
ENV LOCATION=cn
ENV GO111MODULE=on
RUN mkdir /ssd
ENTRYPOINT ["bash", "-c","while true;do echo hello docker;sleep 1;done"]
