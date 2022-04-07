package master

import (
	"fmt"
	"go-test-perf/pkg/constants"
	"go-test-perf/pkg/master/aggregation"
	"go-test-perf/pkg/worker"
)

type Worker interface {
	Execute(hm constants.HttpMethod, url, body string) []*worker.Result
}

type Aggregator interface {
	Check() (r *aggregator.Result)
	Add(res *worker.Result)
}

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

func (m *master) RunTests(hm constants.HttpMethod, url string, body string) {
	for i := range m.wkrs {
		wkr := m.wkrs[i]
		aggr := m.aggr
		results := wkr.Execute(hm, url, body)
		for i := range results {
			aggr.Add(results[i])
		}
	}

	testResults := m.aggr.Check()

	fmt.Println("--------------------------------------------------")
	fmt.Println("Failure count = ", testResults.FailCount)
	fmt.Println("Average req duration = ", testResults.AvgReqDuration)
	fmt.Println("Min req duration = ", testResults.MinReqDuration)
	fmt.Println("Max req duration = ", testResults.MaxReqDuration)
	fmt.Println("Total req fired = ", len(m.wkrs))
	fmt.Println("--------------------------------------------------")

}
