package main

import (
	"go-test-perf/pkg/constants"
	"go-test-perf/pkg/master"
	"go-test-perf/pkg/master/aggregation"
	"go-test-perf/pkg/worker"
)

func main() {
	avgReqDur := 1000.0 // millisecond
	noOfReqForEachWrkr := 10
	aggrConfig := aggregator.SetupConfig(avgReqDur)

	aggregtr := aggregator.Init(aggrConfig)
	workerList := make([]master.Worker, 0)

	for i := 0; i < 2; i++ {
		wrkrCnfg := worker.SetupConfig(noOfReqForEachWrkr)
		wrkrCnfg.Update("http://www.google.com", "", constants.GET)
		wrkrCnfg.Update("http://www.chess.com", "", constants.GET)
		wrkr := worker.Init(wrkrCnfg)
		workerList = append(workerList, wrkr)
	}

	master := master.Init(workerList, aggregtr)

	master.ExecuteWorkers()

}
