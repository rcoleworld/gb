package gb

/*
* Example output from ab:
*
* This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
* Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
* Licensed to The Apache Software Foundation, http://www.apache.org/
*
* Benchmarking localhost (be patient)
*
*
* Server Software:
* Server Hostname:        localhost
* Server Port:            1337
*
* Document Path:          /
* Document Length:        25 bytes
*
* Concurrency Level:      10
* Time taken for tests:   0.390 seconds
* Complete requests:      10000
* Failed requests:        0
* Total transferred:      1420000 bytes
* HTML transferred:       250000 bytes
* Requests per second:    25614.36 [#/sec] (mean)
* Time per request:       0.390 [ms] (mean)
* Time per request:       0.039 [ms] (mean, across all concurrent requests)
* Transfer rate:          3551.99 [Kbytes/sec] received
*
* Connection Times (ms)
*               min  mean[+/-sd] median   max
* Connect:        0    0   0.0      0       0
* Processing:     0    0   0.1      0       6
* Waiting:        0    0   0.1      0       3
* Total:          0    0   0.1      0       6
*
* Percentage of the requests served within a certain time (ms)
*   50%      0
*   66%      0
*   75%      0
*   80%      0
*   90%      0
*   95%      0
*   98%      1
*   99%      1
*  100%      6 (longest request)
*
 */

import (
	"fmt"
	"time"
)

type GbOutput struct {
    Url string
    ServerPort int
    // documentPath string
    // documentLength string // in bytes
    ConcurrencyLevel int
    TotalTime time.Duration 
    CompleteRequests int
    FailedRequests int
    AverageRequestTime time.Duration
    // totalTransfered int // in bytes
    // htmlTransfered int
    // requestsPerSecond float32
    // timePerRequest float32
    // transferRate float32
    // TODO: add rest
}

func GetOutput(o *GbOutput) {
    header := "This is GoBench, Version 0.10\nLicensed under the MIT license"

    fmt.Printf("%s\n", header)
    fmt.Printf("Server Url:               %s\n", o.Url)

    fmt.Printf("Server port:              %d\n", o.ServerPort)
    fmt.Printf("Concurrency Level:        %d\n", o.CompleteRequests)
    fmt.Printf("Time taken for tests:     %s\n", o.TotalTime)
    fmt.Printf("Average time per request: %s\n", o.AverageRequestTime)
    fmt.Printf("Complete requests:        %d\n", o.CompleteRequests)
    fmt.Printf("Failed requests:          %d\n", o.FailedRequests)
}
