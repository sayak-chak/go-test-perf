package aggregator

import (
	"math"
	"net/http"
)

type WorkerResult interface {
	TimeToGetFirstByte() float64
	Err() error
	HttpResponse() *http.Response
}

type aggregator struct {
	cnfg      *config
	wrkrRslts []WorkerResult
}

func New(c *config) *aggregator {
	return &aggregator{
		cnfg: c,
	}
}

func (a *aggregator) Check() (r *Result) {
	failCount := 0
	avgReqDuration := 0.0
	minReqDuration := math.MaxFloat64
	maxReqDuration := 0.0
	for i := range a.wrkrRslts {
		dur := a.wrkrRslts[i].TimeToGetFirstByte()
		if dur > a.cnfg.avgReqDur {
			failCount++
		}
		minReqDuration = math.Min(minReqDuration, dur)
		maxReqDuration = math.Max(maxReqDuration, dur)
		avgReqDuration += dur
	}
	avgReqDuration /= float64(len(a.wrkrRslts))

	return &Result{
		failCount:      failCount,
		avgReqDuration: avgReqDuration,
		minReqDuration: minReqDuration,
		maxReqDuration: maxReqDuration,
	}
}

func (a *aggregator) Add(res WorkerResult) {
	a.wrkrRslts = append(a.wrkrRslts, res)
}
