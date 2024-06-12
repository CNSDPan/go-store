package socket

import (
	"fmt"
	"testing"
	"time"
)

const (
	socketUrl  = "ws://192.168.33.10:6991/ws"
	authToken  = "2gDGQwDxsrX0UG8yRbophdHxHqD"
	authToken2 = "2gDGQugkyFF4MI10hK7WfT3W3Pe"
)

var TClient *TestClient

func TestUserClient1(t *testing.T) {
	tClient, err := New(socketUrl)
	if err != nil {
		return
	}
	err = tClient.Auth(authToken)
	if err != nil {
		return
	}
	TClient = tClient
	tClient.Send()
	tClient.Read()

	go func() {
		for num := 0; num < 5; num++ {
			fmt.Printf("第【%d】发送消息 \r\n ", num)
			tClient.sendMsgChan <- "发送" + time.Now().Format("2006-01-02 15:04:05")
			time.Sleep(3 * time.Second)
		}
	}()

	for {
		select {
		case m := <-tClient.recvMsgChan:
			fmt.Printf("接收消息 %v  \r\n ", m)
		}
	}
}

func TestUserClient2(t *testing.T) {
	tClient, err := New(socketUrl)
	if err != nil {
		return
	}
	err = tClient.Auth(authToken2)
	if err != nil {
		return
	}
	TClient = tClient
	tClient.Send()
	tClient.Read()

	go func() {
		for num := 0; num < 5; num++ {
			fmt.Printf("第【%d】发送消息 \r\n ", num)
			tClient.sendMsgChan <- "发送" + time.Now().Format("2006-01-02 15:04:05")
			time.Sleep(3 * time.Second)
		}
	}()

	for {
		select {
		case m := <-tClient.recvMsgChan:
			fmt.Printf("接收消息 %v  \r\n ", m)
		}
	}
}
