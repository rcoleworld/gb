package gb

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
    Use: "gb",
    Short: "gb - a benchmarking tool similar to ab, written in golang",
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := rootCmd.Execute; err != nil {
        fmt.Fprintf(os.Stderr, "Error running gb: '%s'", err)
        os.Exit(1)
    }
}

