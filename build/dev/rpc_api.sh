#!/bin/bash
cd /var/www/store/rpc/api &&
# 设置打包环境
GOOS=linux GOARCH=amd64 go build -o bin/rpc_api.bin -tags=rpc-api api.go