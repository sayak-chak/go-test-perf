package aggregator

type Result struct {
	FailCount, TotalNoOfReq                        int
	AvgReqDuration, MinReqDuration, MaxReqDuration float64
}
