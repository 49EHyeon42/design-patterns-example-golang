package templateMethod

import "fmt"

type WoodenHouse struct {
	HouseTemplate
}

func (w *WoodenHouse) buildWalls() {
	fmt.Println("Building Wooden Walls")
}

func (w *WoodenHouse) buildPillars() {
	fmt.Println("Building Pillars with Wood coating")
}
