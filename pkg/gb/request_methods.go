package gb

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

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

func post(g *GbHttpReq, c *http.Client, w *sync.WaitGroup, num int) {
    defer w.Done()

    fmt.Printf("sending request %d\n", num)
    res, err := c.Post(g.url, "application/json", bytes.NewBuffer(g.body))
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

