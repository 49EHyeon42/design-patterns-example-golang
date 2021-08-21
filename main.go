package main

import (
	"fmt"
	"golang-designPattern/creational/singleton"
)

func main() {
	s := singleton.New()
	s["this"] = "that"
	s2 := singleton.New()
	fmt.Println("This is ", s2["this"])
}
