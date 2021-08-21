package bridge

import "fmt"

type mac struct {
}

func (m *mac) Print() {
	fmt.Println("Print request for mac")
}
