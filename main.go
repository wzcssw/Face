package main

import (
	"FaceServer/api"
	"FaceServer/lib"
	"net/url"

	"github.com/gin-gonic/gin"
)

//主机地址
var wsUrl = "ws://192.168.1.240:9000/video"

//视频流地址
var rtspUrl = "rtsp://192.168.1.241/user=admin&password=&channel=1&stream=0.sdp"

func init() {
	go lib.GetSessionInterval()
	_, err := lib.GetSession()
	if err != nil {
		panic("face++ 本地登录失败:" + err.Error())
	}
}
func main() {
	// ProcessWebSocket()
	api.API = gin.Default()
	api.LoadAPI()
	api.API.Run()
}

func ProcessWebSocket() {
	//websocket连接
	var ws = wsUrl + "?url=" + url.QueryEscape(rtspUrl)

	wsConn := lib.GetConnect(ws)
	for {
		_, body, errRead := wsConn.ReadMessage()
		lib.CheckError(errRead)
		// go lib.ProcessBytesToFile(body)
		go lib.Send(body)
	}
}
