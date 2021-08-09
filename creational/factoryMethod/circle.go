package factorymethod

import "fmt"

type Circle struct {
	// Shape
}

func (c *Circle) Draw() {
	fmt.Println("Circle - draw() Method.")
}
