package main

import (
	"GetWebsocketsTest/lib"
	"net/url"
)

//主机地址
var wsUrl = "ws://192.168.1.240:9000/video"

//视频流地址
var rtspUrl = "rtsp://192.168.1.241/user=admin&password=&channel=1&stream=0.sdp"

func main() {
	//websocket连接
	var ws = wsUrl + "?url=" + url.QueryEscape(rtspUrl)

	wsConn := lib.GetConnect(ws)
	for {
		_, body, errRead := wsConn.ReadMessage()
		lib.CheckError(errRead)
		go lib.ProcessBytesToFile(body)
	}
}
