package templateMethod

import "fmt"

type GlassHouse struct {
	HouseTemplate
}

func (g *GlassHouse) buildWalls() {
	fmt.Println("Building Glass Walls")
}

func (g *GlassHouse) buildPillars() {
	fmt.Println("Building Pillars with glass coating")
}
