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
	authToken3 = "2gDGQvEugR6Y5riFp2kVLdc7J0O"
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
		for {
			select {
			case <-time.After(10 * time.Second):
				tt := "用户1：哈喽 " + time.Now().Format("2006-01-02 15:04:05")
				fmt.Printf("发送消息 %s \r\n ", tt)
				tClient.sendMsgChan <- tt
			}
		}
	}()

	for {
		select {
		case m := <-tClient.recvMsgChan:
			fmt.Printf("%v  \r\n ", m)
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
		for num := 0; num < 2; num++ {
			tt := time.Now().Format("2006-01-02 15:04:05")
			fmt.Printf("第【%d】发送消息 %s \r\n ", num, tt)
			tClient.sendMsgChan <- "用户2：哈喽 " + tt
			time.Sleep(3 * time.Second)
		}
	}()

	for {
		select {
		case m := <-tClient.recvMsgChan:
			fmt.Printf("%v  \r\n ", m)
		}
	}
}

func TestUserClient3(t *testing.T) {
	tClient, err := New(socketUrl)
	if err != nil {
		return
	}
	err = tClient.Auth(authToken3)
	if err != nil {
		return
	}
	TClient = tClient
	tClient.Send()
	tClient.Read()

	go func() {
		time.After(2 * time.Second)
		for {
			select {
			case <-time.After(time.Microsecond):
				tt := "用户3：哈喽 " + time.Now().Format("2006-01-02 15:04:05.999999999")
				fmt.Printf("发送消息 %s \r\n ", tt)
				tClient.sendMsgChan <- tt
			}
		}
	}()

	for {
		select {
		case m := <-tClient.recvMsgChan:
			fmt.Printf("%v  \r\n ", m)
		}
	}
}

func TestUserClient4(t *testing.T) {
	tClient, err := New(socketUrl)
	if err != nil {
		return
	}
	err = tClient.Auth("2gDGQwhqJQczjkCikEvg3StOKSR")
	if err != nil {
		return
	}
	TClient = tClient
	tClient.Send()
	tClient.Read()

	for {
		select {
		case m := <-tClient.recvMsgChan:
			fmt.Printf("%v  \r\n ", m)
		}
	}
}
