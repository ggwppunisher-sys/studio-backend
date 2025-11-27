package reqid

import (
	"net/http"
)

type roundTripperWrapper struct {
	tripper    http.RoundTripper
	headerName string
}

func NewTransportWrapper(transport http.RoundTripper, headerName string) http.RoundTripper {
	if transport == nil {
		transport = http.DefaultTransport
	}
	return &roundTripperWrapper{
		tripper:    transport,
		headerName: headerName,
	}
}

func (rt roundTripperWrapper) RoundTrip(request *http.Request) (*http.Response, error) {
	request.Header.Add(rt.headerName, FromContext(request.Context()))
	return rt.tripper.RoundTrip(request)
}
