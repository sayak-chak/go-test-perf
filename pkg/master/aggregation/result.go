package aggregator

type Result struct {
	urls map[string]*urlResult
}

type urlResult struct { //TODO - refactor
	failCount, noOfReq                                                    int
	avgReqDuration, minReqDuration, maxReqDuration, cumulativeReqDuration float64
}

func (r *Result) UrlList() []string {
	list := make([]string, 0)
	for url := range r.urls {
		list = append(list, url)
	}
	return list
}

func (r *Result) AvgReqDur(url string) float64 {
	return r.urls[url].avgReqDuration
}

func (r *Result) MinReqDur(url string) float64 {
	return r.urls[url].minReqDuration
}

func (r *Result) MaxReqDur(url string) float64 {
	return r.urls[url].maxReqDuration
}

func (r *Result) FailCount(url string) int {
	return r.urls[url].failCount
}

func (r *Result) NoOfReq(url string) int {
	return r.urls[url].noOfReq
}
