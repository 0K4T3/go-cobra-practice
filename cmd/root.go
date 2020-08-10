package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use: "hoge",
    Short: "A brief description of your application",
    Long: "A longer description that spans multiple lines and likely contain examples and usage of using your application. For example: \nCobra is a CLI library for Go that empowers applications. \nThis application is a tool  to generate the needed files to quickly create a Cobra application.",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("rootCmd Run")
        fmt.Println(args)
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func SetArgs(args []string) {
    rootCmd.SetArgs(args)
}
