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
	res := Result{
		MinReqDuration: math.MaxFloat64,
		TotalNoOfReq:   len(a.wrkrRslts),
	}
	a.updateResult(&res)

	return &res
}

func (a *aggregator) Add(res *worker.Result) {
	a.wrkrRslts = append(a.wrkrRslts, *res)
}

func (a *aggregator) updateResult(res *Result) { //TODO - refactor
	for idx := range a.wrkrRslts {
		if a.wrkrRslts[idx].Err != nil {
			res.FailCount++
			return
		}
		dur := a.wrkrRslts[idx].TimeToGetFirstByte
		if dur > a.cnfg.avgReqDur {
			res.FailCount++
		}
		res.MinReqDuration = math.Min(res.MinReqDuration, dur)
		res.MaxReqDuration = math.Max(res.MaxReqDuration, dur)
		res.AvgReqDuration += dur
	}
	res.AvgReqDuration /= float64((len(a.wrkrRslts)))
}
