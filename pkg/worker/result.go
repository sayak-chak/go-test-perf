package worker

import (
	"net/http"
)

type Result struct {
	Url                string
	HttpRes            *http.Response
	Err                error
	TimeToGetFirstByte float64
}
