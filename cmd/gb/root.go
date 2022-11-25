package gb

import (
	"fmt"
	"os"
	"strings"

	"github.com/rcoleworld/gb/pkg/gb"
	"github.com/spf13/cobra"
)

var ( 
    numOfConcurrentRequests int 
    numOfRequests int
    requestMethod string
    requestBody string
)


var rootCmd = &cobra.Command {
    Use: "gb [url] [OPTIONS]",
    Short: "gb - a benchmarking tool similar to ab, written in golang",
    Long: "gb - a benchmarking tool similar to ab, written in golang (long version)",
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        requestsFlag, _ := cmd.Flags().GetInt(gb.RequestsFlag)
        concurrencyFlag, _ := cmd.Flags().GetInt(gb.ConcurrencyFlag)
        requestMethodFlag, _ := cmd.Flags().GetString(gb.RequestMethodFlag)
        requestBodyFlag, _ := cmd.Flags().GetString(gb.RequestBodyFlag)

        requestMethodFlag = strings.ToUpper(requestMethodFlag)

        if requestMethodFlag != gb.Get && requestMethodFlag != gb.Post {
            fmt.Printf("Invalid request method: %s. allowed: GET, POST", requestMethodFlag)
            return
        }

        if requestMethodFlag == gb.Get && requestBodyFlag != "" {
            fmt.Printf("cannot have a request body for method: %s\n", requestMethodFlag)
            return
        }

       // we will prevent concurrent requests to exceed total requests for now
        if concurrencyFlag > requestsFlag {
            fmt.Println(gb.ConcurrencyExceedsRequestsWarning)
            concurrencyFlag = requestsFlag
        }
        url := args[0]
        req, err :=  gb.NewGbHttpReq(url, requestMethodFlag, []byte(requestBodyFlag))
        
        if err != nil {
            fmt.Printf("error creating request: %s\n", err)
        }

        req.SendRequests(&gb.GbReqOptions{NumOfRequests: requestsFlag, NumOfConcurrentRequests: concurrencyFlag})
    },
}

func init() {
    rootCmd.Flags().IntVarP(&numOfRequests, gb.RequestsFlag, "n", 1, gb.RequestsUsage)
    rootCmd.Flags().IntVarP(&numOfConcurrentRequests, gb.ConcurrencyFlag, "c", 1, gb.ConcurrencyUsage)
    rootCmd.Flags().StringVarP(&requestMethod, gb.RequestMethodFlag, "m", gb.Get , gb.RequestMethodUsage)
    rootCmd.Flags().StringVarP(&requestBody, gb.RequestBodyFlag, "b", "", gb.RequestBodyUsage)
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Error running gb: '%s'", err)
        os.Exit(1)
    }
}

