#!/bin/bash
cd /var/www/store/websocket/ &&
# 设置打包环境
GOOS=linux GOARCH=amd64 go build -o bin/websocket.bin -tags=websocket websocket.go