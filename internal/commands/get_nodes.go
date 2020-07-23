package commands

import (
    "github.com/smart--petea/kubernetes-client/internal/request"
    "github.com/smart--petea/kubernetes-client/internal/helper"
    "encoding/json"
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

type GetNodes struct {
}

func (getNodes *GetNodes) Execute(args []string) error {
    data, err := request.Get("/api/v1/nodes")
    if err != nil {
        return err
    }

    var nodeList NodeList
    err = json.Unmarshal(data, &nodeList)
    if err != nil {
        return err
    }

    printer := helper.NewPrinter(len(nodeList.Items) + 1)
    printer.AddRow("NAME")

    for _, item := range nodeList.Items {
        printer.AddRow(item.Metadata.Name)
    }
    printer.Print()

    return nil
}
