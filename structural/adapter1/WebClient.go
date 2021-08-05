package adapter1

type WebClient struct {
	webRequester WebRequester
}

func NewWebClient(webRequester WebAdapter) *WebClient {
	return &WebClient{&webRequester}
}

func (w *WebClient) DoWork() {
	w.webRequester.requestHandler()
}
