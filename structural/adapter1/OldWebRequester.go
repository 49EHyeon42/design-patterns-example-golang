package adapter1

import "fmt"

type OldWebRequester struct {
}

func (o *OldWebRequester) requestHandler() {
	fmt.Println("OldWebRequester is working")
}
