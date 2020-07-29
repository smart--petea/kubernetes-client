package commands

import (
    "github.com/spf13/cobra"
    "fmt"
)

func NewGetCommand() *cobra.Command {
    var GetCommand = &cobra.Command{
        Use:  "get",
        Run: func (cmd *cobra.Command, args []string)  {
            fmt.Println("get command")
        },
    }

    GetCommand.AddCommand(NewGetNodesCommand())

    return GetCommand
}
