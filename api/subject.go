package api

import (
	"FaceServer/lib"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// API 对外暴露API
var API *gin.Engine

func LoadAPI() {
	API.GET("/subject/:id", func(c *gin.Context) {
		id := c.Param("id")
		result, err := lib.Get(lib.Host+"/subject/"+id, nil)
		if err != nil {
			fmt.Println(err)
		}
		c.Header("Content-type", "application/json; charset=utf-8")
		c.String(200, "%s", string(result))
	})

	API.POST("/subject", func(c *gin.Context) {
		name := c.PostForm("name")
		photo := c.PostForm("photo")
		if name == "" || photo == "" {
			c.JSON(400, "缺少参数")
			return
		}
		// 创建用户
		subjectID, errCreate := lib.CreateSubject(name)
		if errCreate != nil {
			panic(errCreate)
		}
		// 解析Base64
		base64data := []byte(photo)[strings.IndexByte(photo, ',')+1:]         //  取得去掉头的base64
		bytesData, err := base64.StdEncoding.DecodeString(string(base64data)) //成图片文件并把文件写入到buffer
		if err != nil {
			panic(err)
		}
		// 为用户上传识别图片
		errUpload := lib.UploadImage(subjectID, bytesData)
		if errUpload != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, "ok")
		}
	})
}
