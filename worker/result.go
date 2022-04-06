package worker

import (
	"net/http"
)

type result struct {
	httpRes            *http.Response
	err                error
	timeToGetFirstByte float64
}

func (r *result) TimeToGetFirstByte() float64 {
	return r.timeToGetFirstByte
}

func (r *result) Err() error {
	return r.err
}

func (r *result) HttpResponse() *http.Response {
	return r.httpRes
}
