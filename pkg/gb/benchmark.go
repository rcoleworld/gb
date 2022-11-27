package gb

import (
	"fmt"
	"time"
)


func benchmark(benchmarkingData chan time.Duration) {
    totalTime := time.Duration(0)
    for item := range benchmarkingData {
        totalTime += item
        fmt.Printf("request took %s\n", item)
    }

    fmt.Printf("total time %s\n", totalTime)
}
