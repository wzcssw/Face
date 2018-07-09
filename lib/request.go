package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/kataras/iris/core/errors"
)

const (
	Loginname = "303573139@qq.com"
	Password  = "123456"
	Host      = "http://192.168.1.240"
	LoginURL  = "/auth/login"
)

var SessionID = ""

type LoginResult struct {
	Code int
	Desc string
}

func Request(method, url string, params, header map[string]string) ([]byte, error) {
	var r http.Request
	r.ParseForm()
	if params != nil {
		for k, v := range params {
			r.Form.Add(k, v)
		}
	}

	bodystr := strings.TrimSpace(r.Form.Encode())
	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(bodystr))
	// Add Session
	req.Header.Add("Cookie", "session="+SessionID)
	if header != nil {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println("测试：", resp.Header.Get("Content-Type"))
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body, err
}

func Post(url string, params map[string]string) ([]byte, error) {
	header := make(map[string]string)
	header["Content-Type"] = "application/x-www-form-urlencoded"
	return Request("POST", url, params, header)
}

func Get(url string, params map[string]string) ([]byte, error) {
	return Request("GET", url, params, nil)
}

func GetSession() (string, error) {
	var resultError error
	session := ""
	params := make(map[string]string)
	params["username"] = Loginname
	params["password"] = Password
	var r http.Request
	r.ParseForm()
	for k, v := range params {
		r.Form.Add(k, v)
	}
	bodystr := strings.TrimSpace(r.Form.Encode())
	client := &http.Client{}
	req, err := http.NewRequest("POST", Host+LoginURL, strings.NewReader(bodystr))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Koala Admin")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	for _, Cookie := range resp.Cookies() {
		if Cookie.Name == "session" {
			session = Cookie.Value
			break
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	LoginResult := &LoginResult{}
	errJSON := json.Unmarshal(body, LoginResult)
	if errJSON != nil {
		fmt.Println("请求登录接口 JSON 解析错误:", errJSON)
		resultError = errors.New("请求登录接口 JSON 解析错误")
	} else {
		if LoginResult.Code != 0 {
			fmt.Println("请求登录接口 登录失败:", LoginResult.Desc)
			resultError = errors.New("请求登录接口 登录失败")
		}
	}
	SessionID = session
	return session, resultError
}

func GetSessionInterval() {
	for {
		time.Sleep(time.Minute * 20)
		GetSession()
	}
}
