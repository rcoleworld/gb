package gb

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
    "github.com/rcoleworld/gb/pkg/gb"
)

var ( 
    numOfConcurrentRequests int 
    numOfRequests int
)

var rootCmd = &cobra.Command {
    Use: "gb [url] [OPTIONS]",
    Short: "gb - a benchmarking tool similar to ab, written in golang",
    Long: "gb - a benchmarking tool similar to ab, written in golang (long version)",
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        requestsFlag, _ := cmd.Flags().GetInt(gb.RequestsFlag)
        concurrencyFlag, _ := cmd.Flags().GetInt(gb.ConcurrencyFlag)
        url := args[0]
        req, err :=  gb.NewGbHttpReq(url, gb.Get, nil)
        
        if err != nil {
            fmt.Printf("error creating request: %s\n", err)
        }

        req.SendRequests(&gb.GbReqOptions{NumOfRequests: requestsFlag, NumOfConcurrentRequests: concurrencyFlag})
    },
}

func init() {
    rootCmd.Flags().IntVarP(&numOfRequests, gb.RequestsFlag, "n", 1, gb.RequestsUsage)
    rootCmd.Flags().IntVarP(&numOfConcurrentRequests, gb.ConcurrencyFlag, "c", 1, gb.ConcurrencyUsage)
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Error running gb: '%s'", err)
        os.Exit(1)
    }
}

