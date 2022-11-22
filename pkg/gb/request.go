package gb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

const (
    Get  = "GET" 
    Post = "POST" 
)

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
    if method != Get {
        return &GbHttpReq {}, NotImplementedRequestMethodError("method '%s' is not implemented") 
    }
    return &GbHttpReq {
        url,
        method,
        body,
    }, nil
} 

func (g *GbHttpReq) SendRequests(options *GbReqOptions) {
    numOfRequests := 1 
    if options.NumOfRequests !=  0 {
        numOfRequests = options.NumOfRequests
    }

    numOfConcurrentRequests := 1
    if options.NumOfConcurrentRequests != 0 {
        numOfConcurrentRequests = options.NumOfConcurrentRequests
    }

    switch g.method {
    case Get:
        // TODO: fix this broken concurrency
        wg := new(sync.WaitGroup)
        client := &http.Client {Timeout: time.Second * 4}
        for i := 0; i < numOfRequests; i += numOfConcurrentRequests {
            wg.Add(numOfConcurrentRequests)
            for j := 0; j < numOfConcurrentRequests; j++ {
                go get(g.url, client, wg, i + j)
            }
            wg.Wait()
        }
    }
}

func get(url string, c *http.Client, w *sync.WaitGroup, num int) {
    defer w.Done()

    fmt.Printf("sending request %d\n", num)
    res, err := c.Get(url)
    if err != nil {
        fmt.Printf("error: %s\n", err)
        return 
    }

    defer res.Body.Close()
    resBody, err := ioutil.ReadAll(res.Body)

    if err != nil {
        fmt.Printf("error: %s\n", err)
        return
    }

    fmt.Printf("response body, %d: %s\n", num, resBody)
}

