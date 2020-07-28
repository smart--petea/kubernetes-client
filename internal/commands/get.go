package commands

import (
    "github.com/spf13/cobra"
    "fmt"
)

var Get = &cobra.Command{
    Use:  "get",
    Run: func (cmd *cobra.Command, args []string)  {
        fmt.Println("get command")
    },
}
