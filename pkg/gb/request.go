package gb

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

const (
    Get  = "GET" 
    Post = "POST" 
)

type NotImplementedRequestMethodError string

func (e NotImplementedRequestMethodError) Error() string {
    return fmt.Sprintf("%s", string(e))
}

// TODO: Implement headers
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

    switch g.method {
    case Get:
        for i := 0; i < numOfRequests; i++ {
            err := handleGet(g.url)
            if err != nil {
                fmt.Printf("error making get request: %s\n", err) 
            }
        }
    }
}

func handleGet(uri string) error {
    req, err := http.NewRequest(Get, uri, nil)
    if err != nil  {
        return err
    }

    res, err:= http.DefaultClient.Do(req)

    if err != nil {
        return err
    }

    resBody, err := ioutil.ReadAll(res.Body)

    if err != nil {
        return err
    }

    fmt.Printf("response body: %s\n", resBody)
    return nil
}

