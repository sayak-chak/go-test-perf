package worker

import (
	"go-test-perf/constants"
	"net/http"
	"net/http/httptrace"
	"strings"
	"time"
)

type worker struct {
	// TODO - is this the right place?
}

func New() *worker {
	return &worker{}
}

func (w *worker) Execute(hm constants.HttpMethod, url, body string) *Result {

	req, _ := http.NewRequest(string(hm), url, strings.NewReader(body))

	var start time.Time
	var timeToGetFirstByte time.Duration

	trace := &httptrace.ClientTrace{
		DNSStart: func(dsi httptrace.DNSStartInfo) { start = time.Now() },
		GotFirstResponseByte: func() {
			timeToGetFirstByte = time.Since(start)
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	httpRes, err := http.DefaultTransport.RoundTrip(req)

	return &Result{
		HttpRes:            httpRes,
		Err:                err,
		TimeToGetFirstByte: float64(timeToGetFirstByte.Milliseconds()), //TODO - is this enough?
	}

}
