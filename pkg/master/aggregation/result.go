package aggregator

type Result struct { 
	FailCount                                      int
	AvgReqDuration, MinReqDuration, MaxReqDuration float64
}
