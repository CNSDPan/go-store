#!/bin/bash
cd /var/www/store/rpc/socket &&
# 设置打包环境
GOOS=linux GOARCH=amd64 go build -o bin/rpc_socket.bin -tags=rpc-socket socket.go