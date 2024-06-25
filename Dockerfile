FROM golang:1.19 AS builder
ENV TZ=Asia/Shanghai
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /var/www/store

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN apt update&&apt upgrade -y
RUN apt install vim -y