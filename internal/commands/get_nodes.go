package commands

import (
    "github.com/smart--petea/kubernetes-client/internal/request"
    "github.com/smart--petea/kubernetes-client/internal/helper"

    "github.com/spf13/cobra"
    "encoding/json"
    "strings"
)

type NodesTable struct {
    Kind string `json:"kind"`
    ApiVersion string `json:"apiVersion"`
    Metadata map[string]string `json:"metadata"`
    ColumnDefinitions []struct{
        Name string `json:"name"`
        Type string `json:"type"`
        Format string `json:"format"`
        Description string `json:"description"`
        Priority int `json:"priority"`
    } `json:"columnDefinitions"`
    Rows []struct{
        Cells []string `json:"cells"`
    } `json:"rows"`
}

func NewGetNodesCommand() *cobra.Command {
    return &cobra.Command{
        Use:  "nodes",
        RunE: func (cmd *cobra.Command, args []string) error {
            data, err := request.
                Get("/api/v1/nodes").
                AsTable("meta.k8s.io", "v1").
                AsTable("meta.k8s.io", "v1beta1").
                Do()
            if err != nil {
                return err
            }

            var nodesTables NodesTable 
            err = json.Unmarshal(data, &nodesTables)
            if err != nil {
                return err
            }

            var columnIndexes []int
            var columnNames []string
            var columnName string
            for columnIndex, columnDefinition := range nodesTables.ColumnDefinitions {
                if columnDefinition.Priority != 0 {
                    continue
                }

                columnIndexes = append(columnIndexes, columnIndex)

                columnName = strings.ToUpper(columnDefinition.Name)
                columnNames = append(columnNames, columnName)
            }

            printer := helper.NewPrinter(len(columnIndexes) + 1)
            printer.AddRow(columnNames...)

            var forPrint []string
            for _, row := range nodesTables.Rows {
                forPrint = forPrint[:0]
                for _, columnIndex := range columnIndexes {
                    forPrint = append(forPrint, row.Cells[columnIndex])
                }

                printer.AddRow(forPrint...)
            }

            printer.Print()
            return nil
        },
    }
}
