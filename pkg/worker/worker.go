package worker

import (
	"go-test-perf/pkg/constants"
	"net/http"
	"net/http/httptrace"
	"strings"
	"time"
)

type worker struct {
	cnfg *config
}

func Init(c *config) *worker {
	return &worker{cnfg: c}
}

func (w *worker) Execute() []*Result {
	resultlist := make([]*Result, 0)
	for i := 0; i < w.cnfg.noOfReq; i++ {
		for _, info := range w.cnfg.urlInfos {
			resultlist = append(resultlist, w.callUrl(info.hm, info.url, info.body))
		}
	}
	return resultlist
}

func (w *worker) callUrl(hm constants.HttpMethod, url, body string) *Result {
	req, err := http.NewRequest(string(hm), url, strings.NewReader(body))
	if err != nil {
		return &Result{
			Url: url,
			Err: err,
		}
	}
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
		Url:                url,
		HttpRes:            httpRes,
		Err:                err,
		TimeToGetFirstByte: float64(timeToGetFirstByte.Milliseconds()), //TODO - is this enough?
	}

}
