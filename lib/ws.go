package lib

import (
	"fmt"
	"net"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

func GetConnect(rawurl string) *websocket.Conn {
	u, err := url.Parse(rawurl)
	if err != nil {
		fmt.Println(err)
	}
	rawConn, err := net.Dial("tcp", u.Host)
	if err != nil {
		fmt.Println(err)
	}
	wsHeaders := http.Header{
		"Origin":                   {"http://local.host:80"},
		"Sec-WebSocket-Extensions": {"permessage-deflate; client_max_window_bits, x-webkit-deflate-frame"},
	}
	wsConn, resp, errNewClient := websocket.NewClient(rawConn, u, wsHeaders, 1024, 1024)
	if errNewClient != nil {
		fmt.Println(resp)
		panic(errNewClient)
	}
	return wsConn
}
