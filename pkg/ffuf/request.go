package ffuf

// Request holds the meaningful data that is passed for runner for making the query
type Request struct {
	Method   string
	Url      string
	Headers  map[string]string
	Data     []byte
	Input    []byte
	Position int
}

func NewRequest(conf *Config) Request {
	var req Request
	req.Method = conf.Method
	req.Url = conf.Url
	req.Headers = make(map[string]string)
	return req
}
