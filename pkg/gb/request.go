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
    uri string
    method string 
    body []byte
}

func NewGbHttpReq(uri string, method string, body []byte) (*GbHttpReq, error) {
    if method != Get {
       return &GbHttpReq {}, NotImplementedRequestMethodError("method '%s' is not implemented") 
    }
    return &GbHttpReq {
        uri,
        method,
        body,
    }, nil
} 

func (g *GbHttpReq) SendRequest() {
    switch g.method {
    case Get:
        handleGet(g.uri)
    }
}

func handleGet(uri string) {
    req, err := http.NewRequest(Get, uri, nil)
    if err != nil  {
        fmt.Printf("error making http request: %s\n", err)
        return
    }

    res, err:= http.DefaultClient.Do(req)

    if err != nil {
        fmt.Printf("error making http request: %s\n", err)
        return
    }

    resBody, err := ioutil.ReadAll(res.Body)

    if err != nil {
        fmt.Printf("error reading response body: %s\n", err)
        return
    }

    fmt.Printf("response body: %s\n", resBody)

}

