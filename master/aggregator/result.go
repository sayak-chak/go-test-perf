package aggregator

type Result struct { //TODO - make this unexported
	failCount                                      int
	avgReqDuration, minReqDuration, maxReqDuration float64
}

func (r *Result) FailCount() int {
	return r.failCount
}

func (r *Result) AvgReqDur() float64 {
	return r.avgReqDuration
}

func (r *Result) MinReqDuration() float64 {
	return r.minReqDuration
}

func (r *Result) MaxReqDuration() float64 {
	return r.maxReqDuration
}
