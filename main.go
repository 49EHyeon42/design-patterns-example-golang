package main

import (
	"golang-designPattern/structural/adapter"
)

func main() {
	server := adapter.WebAdapter{}
	server.SetWebAdapter(adapter.FancyRequester{})
	client := adapter.WebClient{}
	client.WebClient(&server)
	client.DoWork()
}

/*
link : https://yaboong.github.io/design-pattern/2018/10/15/adapter-pattern/
링크에 계시된 자바를 고로 변환
소스코드 지저분함, 이해 안됨
2021년 8월 2일
*/
