#!/bin/sh
set -xe
protoc -I . --go_out=plugins=grpc:../../../src ./*.proto
