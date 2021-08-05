package adapter1

import "fmt"

type FancyRequester struct {
}

func (f *FancyRequester) FancyRequester() {
	fmt.Println("Yay! fancyRequestHandler is called!")
}
