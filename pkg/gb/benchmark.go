package gb

import (
	"time"
    "fmt"
)

func collectRequestTimes(benchmarkingData chan time.Duration) []time.Duration {
    dataPoints := make([]time.Duration, 0)
    for item := range benchmarkingData {
        dataPoints = append(dataPoints, item)
    }
    return dataPoints
}

func GetTotalRequestTime(requestTimes []time.Duration) time.Duration {
    totalRequestTime := time.Duration(0)
    for item := range requestTimes {
        totalRequestTime += time.Duration(item)
        fmt.Printf("item: %s\n", time.Duration(item))

    }
    return totalRequestTime
}

func GetAverageRequestTime(requestTimes []time.Duration) time.Duration {
    fmt.Printf("Debug: %s\n", requestTimes)
    return GetTotalRequestTime(requestTimes) / time.Duration(len(requestTimes)) 
}

