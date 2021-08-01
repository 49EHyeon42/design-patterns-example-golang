package strategy

import "fmt"

type Language interface {
	Speak()
}

type Korean struct{}

func (K *Korean) Speak() {
	fmt.Println("안녕")
}

type English struct{}

func (E *English) Speak() {
	fmt.Println("Hi")
}
