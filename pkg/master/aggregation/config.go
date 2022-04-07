package aggregator

type config struct {
	avgReqDur float64
}

func SetupConfig(avgReqDur float64) *config {
	return &config{
		avgReqDur: avgReqDur,
	}
}
