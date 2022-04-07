package aggregator

import (
	"go-test-perf/pkg/worker"
	"math"
)

type aggregator struct {
	cnfg      *config
	wrkrRslts []worker.Result
}

func Init(c *config) *aggregator {
	return &aggregator{
		cnfg: c,
	}
}

func (a *aggregator) Check() (r *Result) {
	res := Result{
		urls: make(map[string]*urlResult),
	}
	a.updateResult(&res)

	return &res
}

func (a *aggregator) Add(res *worker.Result) {
	a.wrkrRslts = append(a.wrkrRslts, *res)
}

func (a *aggregator) updateResult(res *Result) { //TODO - refactor
	for idx := range a.wrkrRslts {

		url := a.wrkrRslts[idx].Url
		if _, present := res.urls[url]; !present {
			res.urls[url] = &urlResult{
				minReqDuration: math.MaxFloat64,
			}
		}
		currRes := res.urls[url]
		currRes.noOfReq++

		if a.wrkrRslts[idx].Err != nil {
			currRes.failCount++
			return
		}
		dur := a.wrkrRslts[idx].TimeToGetFirstByte
		if dur > a.cnfg.avgReqDur {
			currRes.failCount++
		}
		currRes.minReqDuration = math.Min(currRes.minReqDuration, dur)
		currRes.maxReqDuration = math.Max(currRes.maxReqDuration, dur)
		currRes.cumulativeReqDuration += dur
	}

	for url := range res.urls {
		res.urls[url].avgReqDuration = res.urls[url].cumulativeReqDuration/float64(res.urls[url].noOfReq)
	}
}
