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
    contentType string
)

type CommandInput struct {
    requestsFlag int
    concurrencyFlag int
    requestMethodFlag string
    requestBodyFlag string
    contentType string
}

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

        commandInput := &CommandInput {
            requestsFlag,
            concurrencyFlag,
            requestMethodFlag,
            requestBody,
            contentType,
        }

        if !isValidInput(commandInput) {
            return
        }

        url := args[0]
        req, err :=  gb.NewGbHttpReq(url, requestMethodFlag, []byte(requestBodyFlag), contentType)
        
        if err != nil {
            fmt.Printf("error creating request: %s\n", err)
        }

        req.SendRequests(&gb.GbReqOptions{NumOfRequests: requestsFlag, NumOfConcurrentRequests: concurrencyFlag})
        output := &gb.GbOutput {
            Url: url,
            ConcurrencyLevel: concurrencyFlag,
            TotalTime: 1.2,
            CompleteRequests: 10,
            FailedRequests: 0,
        }
        gb.GetOutput(output)
    },
}

func isValidInput(ci *CommandInput) bool {
    if ci.requestMethodFlag != gb.Get && ci.requestMethodFlag != gb.Post {
        fmt.Printf("Invalid request method: %s. allowed: GET, POST", ci.requestMethodFlag)
        return false
    }

    if ci.requestMethodFlag == gb.Get && ci.requestBodyFlag != "" {
        fmt.Printf("cannot have a request body for method: %s\n", ci.requestMethodFlag)
        return false
    }

    return true
}

func init() {
    rootCmd.Flags().IntVarP(&numOfRequests, gb.RequestsFlag, "n", 1, gb.RequestsUsage)
    rootCmd.Flags().IntVarP(&numOfConcurrentRequests, gb.ConcurrencyFlag, "c", 1, gb.ConcurrencyUsage)
    rootCmd.Flags().StringVarP(&requestMethod, gb.RequestMethodFlag, "m", gb.Get , gb.RequestMethodUsage)
    rootCmd.Flags().StringVarP(&requestBody, gb.RequestBodyFlag, "b", "", gb.RequestBodyUsage)
    rootCmd.Flags().StringVarP(&contentType, gb.ContentTypeFlag, "T", "application/json", gb.ContentTypeUsage)
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Error running gb: '%s'", err)
        os.Exit(1)
    }
}

