#!/bin/sh
set -xe

protoc  -I . --go_out=plugins=grpc:../../../src ./account.proto ./cootek.pgd.ysession.proto ./cootek.pgd.coin_mgr.proto ./cootek.pgd.config_service.proto ./cootek.pgd.weather_mgr.proto
