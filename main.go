package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

//主机地址
var wsUrl = "ws://192.168.1.240:9000/video"

//视频流地址
var rtspUrl = "rtsp://192.168.1.241/user=admin&password=&channel=1&stream=0.sdp"

func main() {
	//websocket连接
	var ws = wsUrl + "?url=" + url.QueryEscape(rtspUrl)

	wsConn := GetConnect(ws)
	// go Send(wsConn, "王志成")
	for {
		_, body, errRead := wsConn.ReadMessage()
		if errRead != nil {
			panic(errRead)
		}
		result := &Result{}
		errJSON := json.Unmarshal(body, result)
		if errJSON != nil {
			panic(errJSON)
		}
		fmt.Println(result)
		// fmt.Println(code, string(body), errRead)

	}
}

func Send(wsConn *websocket.Conn, msg string) {
	time.Sleep(2e9)
	wsConn.WriteMessage(1, []byte(msg))
}

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

/*
	wsConn := GetConnect("ws://121.40.165.18:8800") //  测试IP
*/
type Result struct {
	Data Data
}

type Data struct {
	Status    string
	Track     int
	Timestamp int64
	Face      Face
	Person    Person
	Quality   float64
}

type Face struct {
	Image string
}

type Person struct {
	FeatureID  int `json:"feature_id"`
	Confidence float64
	Tag        string
	ID         string `json:"id"`
}
