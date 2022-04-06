package master

import (
	"fmt"
	"go-test-perf/constants"
	"go-test-perf/master/aggregator"
)

type Worker interface {
	Execute(hm constants.HttpMethod, url, body string) (res aggregator.WorkerResult)
}

type Aggregator interface {
	Check() (r *aggregator.Result)
	Add(res aggregator.WorkerResult)
}

// TODO - see how to have this without having import cycle
// type AggregatorResult interface {
// 	FailCount() int
// 	AvgReqDur() float64
// 	MinReqDuration() float64
// 	MaxReqDuration() float64
// }

type master struct {
	wkrs []Worker
	aggr Aggregator
}

func New(wrks []Worker, agr Aggregator) *master {
	return &master{
		wkrs: wrks,
		aggr: agr,
	}
}

func (m *master) RunTests() {
	for i := range m.wkrs {
		wkr := m.wkrs[i]
		aggr := m.aggr
		aggr.Add(wkr.Execute(constants.GET, "http://www.google.com", "")) //TODO - remove dummy val
	}

	testResults := m.aggr.Check()

	fmt.Println("--------------------------------------------------")
	fmt.Println("Failure count = ", testResults.FailCount())
	fmt.Println("Average req duration = ", testResults.AvgReqDur())
	fmt.Println("Min req duration = ", testResults.MinReqDuration())
	fmt.Println("Max req duration = ", testResults.MaxReqDuration())
	fmt.Println("--------------------------------------------------")

}
