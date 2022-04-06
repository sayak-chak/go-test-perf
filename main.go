package main

import (
	"go-test-perf/master"
	"go-test-perf/master/aggregator"
	"go-test-perf/worker"
)

// "go-test-perf/constants"
// "go-test-perf/master"

func main() {
	// req, _ := http.NewRequest("GET", "http://www.google.com", nil)

	// var start, s time.Time
	// var x time.Duration

	// trace := &httptrace.ClientTrace{
	// 	DNSStart: func(dsi httptrace.DNSStartInfo) { s = time.Now() },
	// 	GotFirstResponseByte: func() {
	// 		x = time.Since(s)
	// 	},
	// }

	// req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	// start = time.Now()
	// if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Total time: %v\n", time.Since(start))

	// worker.Call(constants.GET, "http://www.google.com", "")
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
