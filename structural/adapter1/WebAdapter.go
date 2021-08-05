package adapter1

type WebAdapter struct {
	fancyRequester FancyRequester
}

func NewWebAdapter(fancyRequester FancyRequester) *WebAdapter {
	return &WebAdapter{fancyRequester}
}

func (w *WebAdapter) requestHandler() {
	w.fancyRequester.FancyRequester()
}
