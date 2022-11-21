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
    Use: "gb",
    Short: "gb - a benchmarking tool similar to ab, written in golang",
    Long: "gb - a benchmarking tool similar to ab, written in golang (long version)",
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func init() {
    rootCmd.Flags().StringVarP(&uri, "uri", "u", "", gb.UriUsage)
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Error running gb: '%s'", err)
        os.Exit(1)
    }
}

