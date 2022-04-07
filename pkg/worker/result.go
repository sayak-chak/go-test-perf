package worker

import (
	"net/http"
)

type Result struct {
	HttpRes            *http.Response
	Err                error
	TimeToGetFirstByte float64
}