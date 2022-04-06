package aggregator


type config struct {
	avgReqDur float64
	noOfReq   int
}

func SetupConfig(avgReqDur float64, noOfReq int) *config {
	return &config{
		avgReqDur: avgReqDur,
		noOfReq:   noOfReq,
	}
}
