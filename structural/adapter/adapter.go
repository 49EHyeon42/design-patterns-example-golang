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

func NewWebClient(webRequester WebAdapter) *WebClient {
	return &WebClient{&webRequester}
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

func (f *FancyRequester) FancyRequester() {
	fmt.Println("Yay! fancyRequestHandler is called!")
}

// WebAdapter
type WebAdapter struct {
	fancyRequester FancyRequester
}

func (w *WebAdapter) SetWebAdapter(fancyRequester FancyRequester) {
	w.fancyRequester = fancyRequester
}

func NewWebAdapter(fancyRequester FancyRequester) *WebAdapter {
	return &WebAdapter{fancyRequester}
}

func (w *WebAdapter) requestHandler() {
	w.fancyRequester.FancyRequester()
}
