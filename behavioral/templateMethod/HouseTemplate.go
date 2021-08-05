package templateMethod

import "fmt"

type iBuildHouse interface {
	buildFoundation()
	buildPillars()
	buildWalls()
	buildWindows()
}

type HouseTemplate struct {
	iBuildHouse
}

func (h *HouseTemplate) BuildHouse() {
	h.buildFoundation()
	h.buildPillars()
	h.buildWalls()
	h.buildWindows()
	fmt.Println("House is built.")
}

func (h *HouseTemplate) buildFoundation() {
	fmt.Println("Building foundation with cement, icon rods and sand")
}

func (h *HouseTemplate) buildWindows() {
	fmt.Println("Building Glass Windows")
}
