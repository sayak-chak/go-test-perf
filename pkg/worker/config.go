package worker

import "go-test-perf/pkg/constants"

type config struct {
	noOfReq    int
	urlInfos []*urlInfo
}

type urlInfo struct {
	url  string
	body string
	hm   constants.HttpMethod
}

func SetupConfig(noOfReq int) *config {
	return &config{
		noOfReq:    noOfReq,
		urlInfos: make([]*urlInfo, 0),
	}
}

func (c *config) Update(url, body string, hm constants.HttpMethod) {
	c.urlInfos = append(c.urlInfos, &urlInfo{
		url:  url,
		body: body,
		hm:   hm,
	})
}
