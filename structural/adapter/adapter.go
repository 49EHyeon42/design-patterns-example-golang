package adapter

import "fmt"

// WebRequester Interface
type WebRequester interface {
	requestHandler()
}

// WebClient
type WebClient struct {
	webRequester WebRequester
}

func (w *WebClient) WebClient(webRequester WebRequester) {
	w.webRequester = webRequester
}

func (w *WebClient) DoWork() {
	w.webRequester.requestHandler()
}

// OldWebRequester
type OldWebRequester struct {
}

func (o *OldWebRequester) requestHandler() {
	fmt.Println("OldWebRequester is working")
}

// FancyRequester
type FancyRequester struct {
}

func (f *FancyRequester) fancyRequester() {
	fmt.Println("Yay! fancyRequestHandler is called!")
}

// WebAdapter
type WebAdapter struct {
	fancyRequester FancyRequester
}

func (w *WebAdapter) SetWebAdapter(fancyRequester FancyRequester) {
	w.fancyRequester = fancyRequester
}

func (w *WebAdapter) requestHandler() {
	w.fancyRequester.fancyRequester()
}
