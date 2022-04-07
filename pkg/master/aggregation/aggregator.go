package aggregator

import (
	"go-test-perf/pkg/worker"
	"math"
)

type aggregator struct {
	cnfg      *config
	wrkrRslts []worker.Result
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
		if a.wrkrRslts[i].Err != nil {
			failCount++
			continue
		}
		dur := a.wrkrRslts[i].TimeToGetFirstByte
		if dur > a.cnfg.avgReqDur {
			failCount++
		}
		minReqDuration = math.Min(minReqDuration, dur)
		maxReqDuration = math.Max(maxReqDuration, dur)
		avgReqDuration += dur
	}
	avgReqDuration /= float64(len(a.wrkrRslts))

	return &Result{
		FailCount:      failCount,
		AvgReqDuration: avgReqDuration,
		MinReqDuration: minReqDuration,
		MaxReqDuration: maxReqDuration,
	}
}

func (a *aggregator) Add(res *worker.Result) {
	a.wrkrRslts = append(a.wrkrRslts, *res)
}
