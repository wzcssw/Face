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

/*
{"description": "", "remark": "", "subject_type": 0, "name": "\u738b\u5fd7\u6210", "inviter_id": null, "start_time": 0, "title": "", "interviewee": "", "avatar": "/static/upload/avatar/2018-06-27/v2_7ccc8a5dd4f55619b783ea24e6d70469a451260b.jpg", "origin_photo_id": 3, "birthday": null, "id": 4, "entry_date":null, "department": "", "interviewee_pinyin": "", "job_number": "", "end_time": 0}
*/
