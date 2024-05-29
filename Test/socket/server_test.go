package socket

import (
	"fmt"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	var c rest.RestConf
	c.Host = "0.0.0.0"
	c.Port = 7000
	s := rest.MustNewServer(c)
	defer s.Stop()
	go h.run()
	s.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/ws",
		Handler: myws,
	})

	fmt.Printf("Starting websocket server at %s:%d...\n", c.Host, c.Port)
	s.Start()
}
