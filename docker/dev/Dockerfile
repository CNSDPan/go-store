FROM golang:1.19
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /var/www/k

#COPY go.mod .
#COPY go.sum .
#RUN go mod download

RUN apt update&&apt upgrade -y
RUN apt install vim -y