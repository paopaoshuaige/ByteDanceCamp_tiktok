package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRegister(t *testing.T) {
	// 要用postman formdata测试才行
	url := "http://192.168.124.6:8989/douyin/user/register/?username=%3Cusername%3E&password=%3Cpassword%3E"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func TestQueryUserInfo(t *testing.T) {
	url := "http://192.168.124.6:8989/douyin/user/?user_id=5&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6NSwiVXNlcm5hbWUiOiJ0ZXN0IiwiZXhwIjoxNjc2NTQ2OTQxfQ.xlSS-VgbAAsCleVhV7stLYOISa8gqX1vKOyV_1pUB7A"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
