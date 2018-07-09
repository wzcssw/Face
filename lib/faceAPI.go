package lib

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/kataras/iris/core/errors"
)

// UploadImage 为用户上传识别图片
func UploadImage(subjectID string, bytesData []byte) error {
	extraParams := map[string]string{
		"subject_id": subjectID,
	}
	request, err := NewFileUploadRequest("http://192.168.1.240/subject/photo", extraParams, "photo", bytesData)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	Result := &Result{}
	json.Unmarshal(body, Result)
	if Result.Code != 0 {
		return errors.New(Result.Desc)
	}
	return nil
}

// CreateSubject 新增用户
func CreateSubject(name string) (string, error) {
	param := make(map[string]string)
	param["gender"] = "1"
	param["subject_type"] = "0"
	param["name"] = name
	bytes, err := Post("http://192.168.1.240/subject", param)
	if err != nil {
		return "", err
	}
	Result := &Result{}
	errMarsh := json.Unmarshal(bytes, Result)

	return strconv.Itoa(Result.Data.ID), errMarsh
}
