package lib

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"time"
)

func SaveBase64Image(str string) {
	bytes, errDecode := base64.StdEncoding.DecodeString(str)
	if errDecode != nil {
		panic(errDecode)
	}

	dirName, fullPath := MakeName()
	os.MkdirAll(dirName, os.ModePerm)
	os.Create(fullPath)
	errFile := ioutil.WriteFile(fullPath, bytes, 0666)
	CheckError(errFile)
}

// 1: dirName 2: fullPath
func MakeName() (string, string) {
	now := time.Now()
	dirName := "./images/" + now.Format("2006-01-02")
	fileName := now.Format("150405") + ".jpg"
	return dirName, dirName + "/" + fileName
}
