package master

import (
	"fmt"
	"go-test-perf/pkg/constants"
	"go-test-perf/pkg/master/aggregation"
	"go-test-perf/pkg/worker"
	"sync"
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
	var workerWg sync.WaitGroup
	var aggregatorWg sync.WaitGroup
	results := make(chan worker.Result)

	for i := range m.wkrs {
		workerWg.Add(1)
		go m.executeWorker(i, hm, url, body, &workerWg, results)
	}

	aggregatorWg.Add(1)
	go m.aggregate(&aggregatorWg, results)

	workerWg.Wait()
	close(results)
	aggregatorWg.Wait()

	m.displayMetrics(m.aggr.Check())

}

func (*master) displayMetrics(testResults *aggregator.Result) {
	fmt.Println("--------------------------------------------------")
	fmt.Println("Failure count = ", testResults.FailCount)
	fmt.Println("Average req duration = ", testResults.AvgReqDuration)
	fmt.Println("Min req duration = ", testResults.MinReqDuration)
	fmt.Println("Max req duration = ", testResults.MaxReqDuration)
	fmt.Println("Total request count = ", testResults.TotalNoOfReq)
	fmt.Println("--------------------------------------------------")
}

func (m *master) aggregate(wg *sync.WaitGroup, results <-chan worker.Result) {
	defer wg.Done()

	for {
		res, isOpen := <-results
		if !isOpen {
			break
		}

		m.aggr.Add(&res)
	}
}

func (m *master) executeWorker(idx int, hm constants.HttpMethod, url string, body string, wg *sync.WaitGroup, results chan<- worker.Result) { //TODO - refactor
	defer wg.Done()

	wkr := m.wkrs[idx]
	wrkrResults := wkr.Execute(hm, url, body)
	for i := range wrkrResults {
		results <- *wrkrResults[i]
	}
}
