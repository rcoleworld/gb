package gb

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
    "github.com/rcoleworld/gb/pkg/gb"
)

var uri string
var numberOfRequests int

var rootCmd = &cobra.Command {
    Use: "gb [OPTIONS]",
    Short: "gb - a benchmarking tool similar to ab, written in golang",
    Long: "gb - a benchmarking tool similar to ab, written in golang (long version)",
    Run: func(cmd *cobra.Command, args []string) {
        uriFlag, _ := cmd.Flags().GetString("uri")
        req, err :=  gb.NewGbHttpReq(uriFlag, gb.Get, nil)
        
        if err != nil {
            fmt.Printf("error creating request: %s\n", err)
        }

        req.SendRequest()
    },
}

func init() {
    rootCmd.Flags().StringVarP(&uri, "uri", "u", "", gb.UriUsage)
    rootCmd.MarkFlagRequired("uri")
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Error running gb: '%s'", err)
        os.Exit(1)
    }
}

