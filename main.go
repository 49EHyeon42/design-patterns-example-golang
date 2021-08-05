package main

import (
	"fmt"
	"golang-designPattern/behavioral/templateMethod"
)

func main() {
	fmt.Println("1.")
	var wood templateMethod.WoodenHouse
	wood.BuildHouse()
	// fmt.Println("2.")
	// glass := templateMethod.GlassHouse{}
	// glass.BuildHouse()
}

// 자바에서 golang 으로 구현
// 참고 : https://stackoverflow.com/questions/38123911/golang-method-override
// 링크 : https://niceman.tistory.com/142
