package socket

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"store/tools"
	"strconv"
	"sync/atomic"
	"testing"
	"time"
)

type Ch struct {
	Msg string
}

type Bucket struct {
	routines    []chan *Ch
	routinesNum uint64
}

func TestClientId(t *testing.T) {
	serverId := "299"
	fmt.Printf("server id %v \r\n", serverId)
	nodeId, err := strconv.ParseInt(serverId, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	node, e := snowflake.NewNode(nodeId)
	if e != nil {
		fmt.Println(e)
		return
	}
	for i := 0; i < 10; i++ {
		cliendId := node.Generate().String()
		c, e1 := tools.Encrypt([]byte(cliendId), []byte("Adba723b7fe06819"))
		if e1 != nil {
			panic(e1)
		}
		fmt.Printf("%v %v\r\n", cliendId, c)
	}
}

func TestSystemClient(t *testing.T) {
	var b = new(Bucket)
	b.routines = make([]chan *Ch, 10)
	for i := 1; i < 5; i++ {
		c := make(chan *Ch, 20)
		b.routines[i] = c
		go goCh(c)
	}
	str, err := tools.Encrypt([]byte("172.20.2.111:6901"), []byte("Adba723b7fe06819"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \r\n", str)
	for i := 1; i < 5; i++ {
		num := atomic.AddUint64(&b.routinesNum, 1) % uint64(32)
		fmt.Printf("num %d \r\n", num)
		b.routines[num] <- &Ch{Msg: "好呀你是：" + strconv.FormatUint(num, 10)}
	}
	time.Sleep(5 * time.Second)
	fmt.Println("结束了")
}

func goCh(ch chan *Ch) {
	var age *Ch
	for {
		age = <-ch
		fmt.Printf("接收msg:%v \r\n", age.Msg)
	}
}
