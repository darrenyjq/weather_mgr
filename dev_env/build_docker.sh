#!/bin/sh
set -xe
tag=`date +%Y%m%d_%H%M`
docker build --no-cache  -f dev_env.dockerfile  -t harbor.bbbbservice.com/ime_us/go_dev_env -t harbor.bbbbservice.com/ime_us/go_dev_env:$tag .

