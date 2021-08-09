package factorymethod

import "fmt"

type Square struct {
	Shape
}

func (c *Square) Draw() {
	fmt.Println("Square - draw() Method.")
}
