package master

import (
	"fmt"
	"go-test-perf/pkg/master/aggregation"
	"go-test-perf/pkg/worker"
	"sync"
)

type Worker interface {
	Execute() []*worker.Result
}

type Aggregator interface {
	Check() (r *aggregator.Result)
	Add(res *worker.Result)
}

type master struct {
	wkrs []Worker
	aggr Aggregator
}

func Init(wrks []Worker, agr Aggregator) *master {
	return &master{
		wkrs: wrks,
		aggr: agr,
	}
}

func (m *master) ExecuteWorkers() {
	var workerWg sync.WaitGroup
	var aggregatorWg sync.WaitGroup
	results := make(chan worker.Result)

	for i := range m.wkrs {
		workerWg.Add(1)
		go m.executeWorker(i, &workerWg, results)
	}

	aggregatorWg.Add(1)
	go m.aggregate(&aggregatorWg, results)

	workerWg.Wait()
	close(results)
	aggregatorWg.Wait()

	m.displayMetrics(m.aggr.Check())
}

func (*master) displayMetrics(aggrResults *aggregator.Result) {
	urlList := aggrResults.UrlList()
	for _, url := range urlList {
		fmt.Println("--------------------------------------------------")
		fmt.Println("For url", url)
		fmt.Println("Failure count = ", aggrResults.FailCount(url))
		fmt.Println("Average req duration = ", aggrResults.AvgReqDur(url))
		fmt.Println("Min req duration = ", aggrResults.MinReqDur(url))
		fmt.Println("Max req duration = ", aggrResults.MaxReqDur(url))
		fmt.Println("Total request count = ", aggrResults.NoOfReq(url))
		fmt.Println("--------------------------------------------------")
	}
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

func (m *master) executeWorker(idx int, wg *sync.WaitGroup, results chan<- worker.Result) { //TODO - refactor
	defer wg.Done()

	wkr := m.wkrs[idx]
	wrkrResults := wkr.Execute()
	for i := range wrkrResults {
		results <- *wrkrResults[i]
	}
}
