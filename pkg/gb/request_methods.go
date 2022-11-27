package gb

import (
	"bytes"
	"fmt"
    "time"
	"io/ioutil"
	"net/http"
	"sync"
)

func get(g *GbHttpReq, c *http.Client, w *sync.WaitGroup, num int, benchmarkingData chan time.Duration) {
    defer w.Done()
    startTime := time.Now() 
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
    benchmarkingData <- time.Since(startTime)
}

func post(g *GbHttpReq, c *http.Client, w *sync.WaitGroup, num int, benchmarkingData chan time.Duration) {
    defer w.Done()
    startTime := time.Now() 
    fmt.Printf("sending request %d with body %s\n", num, string(g.body))
    res, err := c.Post(g.url, g.contentType, bytes.NewBuffer(g.body))
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
    benchmarkingData <- time.Since(startTime)
}

