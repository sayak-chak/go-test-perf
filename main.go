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
	config := aggregator.SetupConfig(avgReqDur)

	aggregtr := aggregator.New(config)
	workerList := make([]master.Worker, 0)

	workerList = append(workerList, worker.New(worker.SetupConfig(noOfReqForEachWrkr)))
	workerList = append(workerList, worker.New(worker.SetupConfig(noOfReqForEachWrkr)))

	master := master.New(workerList, aggregtr)

	master.RunTests(constants.GET, "http://www.google.com", "")

}
