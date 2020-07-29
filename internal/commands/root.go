package commands

import (
    "github.com/spf13/cobra"
    "fmt"
)

func NewRootCommand() *cobra.Command {
    var rootCmd = &cobra.Command{
        Use:  "use",
        Short: "short", 
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("ok")
        },
    }

    rootCmd.AddCommand(NewClusterInfo())
    rootCmd.AddCommand(NewGetCommand())

    return rootCmd
}
