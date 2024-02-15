package http

import "net/http"

type Request struct {
	*http.Request
}

func FromRequest(req *http.Request) *Request {
	return &Request{
		Request: req,
	}
}
