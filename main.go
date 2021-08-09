package main

import (
	factorymethod "golang-designPattern/creational/factoryMethod"
)

func main() {
	factorymethod.GetShape("circle").Draw()
	factorymethod.GetShape("rectangle").Draw()
	factorymethod.GetShape("square").Draw()
}
