package main

import (
	"go-test-perf/master"
	"go-test-perf/master/aggregator"
	"go-test-perf/worker"
)


func main() {
	avgReqDur := 1000.0 // millisecond
	noOfReq := 10
	config := aggregator.SetupConfig(avgReqDur, noOfReq)

	aggregtr := aggregator.New(config)
	workerList := make([]master.Worker, 0)

	worker := worker.New()

	workerList = append(workerList, worker)

	master := master.New(workerList, aggregtr)

	master.RunTests()

}
