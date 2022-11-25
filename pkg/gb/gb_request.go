package gb

import (
	"fmt"
)

const (
    Get  = "GET" 
    Post = "POST" 
)

var SupportedMethods = map[string]bool{
    Get: true,
    Post:true,
}

type NotImplementedRequestMethodError string

func (e NotImplementedRequestMethodError) Error() string {
    return fmt.Sprintf("%s", string(e))
}

type GbHttpReq struct {
    url string
    method string 
    body []byte
}

type GbReqOptions struct {
    NumOfRequests int
    NumOfConcurrentRequests int
}

func NewGbHttpReq(url string, method string, body []byte) (*GbHttpReq, error) {
    if !SupportedMethods[method] {
        return &GbHttpReq {}, NotImplementedRequestMethodError("method '%s' is not implemented") 
    }
    return &GbHttpReq {
        url,
        method,
        body,
    }, nil
} 

func (g *GbHttpReq) SendRequests(options *GbReqOptions) {
    if options.NumOfRequests ==  0 {
        options.NumOfRequests = 1 
    }

    if options.NumOfConcurrentRequests == 0 {
        options.NumOfConcurrentRequests = 1
    }

    switch g.method {
    case Get:
        handleConcurrentRequests(g, options, get)
    case Post:
        handleConcurrentRequests(g, options, post)
    }
}

