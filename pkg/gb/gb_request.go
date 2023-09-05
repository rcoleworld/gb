package gb

import (
	"fmt"
	"time"
)

const (
    Get  = "GET" 
    Post = "POST" 
)

var SupportedMethods = map[string]bool{
    Get: true,
    Post: true,
}

type NotImplementedRequestMethodError string

func (e NotImplementedRequestMethodError) Error() string {
    return fmt.Sprintf("%s", string(e))
}

type GbHttpReq struct {
    url string
    method string 
    body []byte
    contentType string
}

type GbReqOptions struct {
    NumOfRequests int
    NumOfConcurrentRequests int
}

func NewGbHttpReq(url string, method string, body []byte, contentType string) (*GbHttpReq, error) {
    if !SupportedMethods[method] {
        return &GbHttpReq {}, NotImplementedRequestMethodError("method '%s' is not implemented") 
    }

    return &GbHttpReq {
        url,
        method,
        body,
        contentType,
    }, nil
} 

func (g *GbHttpReq) SendRequests(options *GbReqOptions) []time.Duration {
    if options.NumOfRequests < 1 {
        options.NumOfRequests = 1 
    }

    if options.NumOfConcurrentRequests < 1 {
        options.NumOfConcurrentRequests = 1
    }

    if options.NumOfConcurrentRequests > options.NumOfRequests {
        fmt.Println(ConcurrencyExceedsRequestsWarning)
        options.NumOfConcurrentRequests = options.NumOfRequests
    }

    switch g.method {
    case Get:
        return handleConcurrentRequests(g, options, get)
    case Post:
        return handleConcurrentRequests(g, options, post)
    default:
        return nil
    }
}

