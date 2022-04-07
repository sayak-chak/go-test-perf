package worker

type config struct {
	noOfReq int
}

func SetupConfig(noOfReq int) *config {
	return &config{
		noOfReq: noOfReq,
	}
}
