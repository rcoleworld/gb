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
    if options.NumOfRequests ==  0 {
        options.NumOfRequests = 1 
    }

    if options.NumOfConcurrentRequests == 0 {
        options.NumOfConcurrentRequests = 1
    }

    switch g.method {
    case Get:
        handleConcurrentRequests(g, options, get)
    }
}

type request func(*GbHttpReq, *http.Client, *sync.WaitGroup, int)

func handleConcurrentRequests(g *GbHttpReq, o *GbReqOptions, r request) {
    wg := new(sync.WaitGroup)
    client := &http.Client {Timeout: time.Second * 4}
    wg.Add(o.NumOfConcurrentRequests)

    goRoutineCounter := o.NumOfConcurrentRequests 
    for i := 0; i < o.NumOfRequests; i++ {
        if goRoutineCounter == 0 {
            wg.Wait()
            numLeft := o.NumOfRequests- i
            if numLeft < o.NumOfConcurrentRequests {
                goRoutineCounter = numLeft
                wg.Add(numLeft)
            } else {
                goRoutineCounter = o.NumOfConcurrentRequests 
                wg.Add(o.NumOfConcurrentRequests)
            }
        } 
        go r(g, client, wg, i)
        goRoutineCounter--
    }
    wg.Wait()
}

func get(g *GbHttpReq, c *http.Client, w *sync.WaitGroup, num int) {
    defer w.Done()

    fmt.Printf("sending request %d\n", num)
    res, err := c.Get(g.url)
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

