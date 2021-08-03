package main

import (
	"golang-designPattern/structural/adapter"
)

func main() {
	newAdapter := adapter.NewWebAdapter(adapter.FancyRequester{})
	client := adapter.NewWebClient(*newAdapter)
	client.DoWork()
}

/*
link : https://yaboong.github.io/design-pattern/2018/10/15/adapter-pattern/
*/
