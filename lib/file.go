package lib

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"time"
)

func SaveBase64Image(dataStr, name string) {
	bytes, errDecode := base64.StdEncoding.DecodeString(dataStr)
	if errDecode != nil {
		panic(errDecode)
	}

	dirName, fullPath := MakeName(name)
	os.MkdirAll(dirName, os.ModePerm)
	os.Create(fullPath)
	errFile := ioutil.WriteFile(fullPath, bytes, 0666)
	CheckError(errFile)
}

// 1: dirName 2: fullPath
func MakeName(name string) (string, string) {
	now := time.Now()
	dirName := "./images/" + now.Format("2006-01-02")
	fileName := now.Format("15:04:05") + "_" + name + ".jpg"
	return dirName, dirName + "/" + fileName
}
