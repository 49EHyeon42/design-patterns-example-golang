package main

import (
	"golang-designPattern/behavioral/strategy"
)

type Person struct {
	language strategy.Language
}

func (p *Person) setLanguage(language strategy.Language) {
	p.language = language
}

func (p *Person) speakLanguage() {
	p.language.Speak()
}

func main() {
	person := Person{}
	person.setLanguage(&strategy.English{})
	person.speakLanguage()
}
