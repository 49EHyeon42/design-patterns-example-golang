package main

import "fmt"

type ServiceProcess interface {
	RunSomething() string
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) RunSomething() string {
	return "Service Run"
}

type Proxy struct{}

func NewProxy() *Proxy {
	return &Proxy{}
}

func (p *Proxy) RunSomething() string {
	fmt.Println("호출에 대한 흐름 제어, 반환 결과 전달")
	newService := NewService()
	return newService.RunSomething()
}

func main() {
	proxy := NewProxy()
	fmt.Println(proxy.RunSomething())
}

/*
reference link : https://limkydev.tistory.com/79
- 흐름제어 목적으로 사용

*/
