package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestFollowList(t *testing.T) {
	url := "http://127.0.0.1:8989/douyin/relation/follow/list/?user_id=8&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6OCwiVXNlcm5hbWUiOiJ0ZXN0YWFxIiwiZXhwIjoxNjc3MDUxNTc1fQ.Ls49xguR5V0XabYTMNJCePa-JKz19RD4984HGll19OY"
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
