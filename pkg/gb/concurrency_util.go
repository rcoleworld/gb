package gb

import (
    "net/http"
    "sync"
    "time"
)

type requestFn func(*GbHttpReq, *http.Client, *sync.WaitGroup, int, chan time.Duration)

func handleConcurrentRequests(g *GbHttpReq, o *GbReqOptions, r requestFn) {
    wg := new(sync.WaitGroup)
    client := &http.Client {Timeout: time.Second * 4}
    wg.Add(o.NumOfConcurrentRequests)

    benchmarkingData := make(chan time.Duration)

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
        go r(g, client, wg, i, benchmarkingData)
        goRoutineCounter--
    }
    wg.Wait()
    close(benchmarkingData)
    benchmark(benchmarkingData)
}

