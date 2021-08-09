package factorymethod

import "fmt"

type Rectangle struct {
	Shape
}

func (c *Rectangle) Draw() {
	fmt.Println("Rectangle - draw() Method.")
}
