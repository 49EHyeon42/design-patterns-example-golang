package bridge

import "fmt"

type windows struct {
}

func (w *windows) Print() {
	fmt.Println("Print request for windows")
}
