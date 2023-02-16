package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_feed(t *testing.T) {
	url := "http://192.168.124.6:8989/douyin/feed/?latest_time=%3Clatest_time%3E&token=%3Ctoken%3E"
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
