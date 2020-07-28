package commands

import (
    "github.com/smart--petea/kubernetes-client/internal/request"
    "github.com/spf13/cobra"
//    "github.com/smart--petea/kubernetes-client/internal/helper"
 //   "encoding/json"
    "fmt"
)

type NodeList struct {
    Kind string `json:"kind"`
    ApiVersion string `json:"apiVersion"`
    Metadata map[string]string `json:"metadata"`
    Items []struct{
        Metadata struct{
            Name string `json:"name"`
        } `json:"metadata"`
        Labels map[string]string `json:"labels"`
        Annotations map[string]string `json:"annotations"`
        ManagedFields map[string]string `json:"managedFields"`
        Spec map[string]interface{} `json:"spec"`
        Status map[string]interface{} `json:"status"`
    } `json:"items"`
}

var GetNodes = &cobra.Command{
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

        fmt.Printf("%s\n", data)
        return nil
    },
}

//func (getNodes *GetNodes) Execute(args []string) error {

    /*
    var nodeList NodeList
    err = json.Unmarshal(data, &nodeList)
    if err != nil {
        return err
    }

    printer := helper.NewPrinter(len(nodeList.Items) + 1)
    printer.AddRow("NAME", "STATUS", "ROLES", "AGE", "VERSION")

    for _, item := range nodeList.Items {
        printer.AddRow(item.Metadata.Name)
    }
    printer.Print()
    */

 //   return nil
//}
