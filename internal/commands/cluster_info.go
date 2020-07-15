package commands

import (
    "fmt"
    "github.com/smart--petea/kubernetes-client/internal/request"
)

type ClusterInfo struct {
}

func (clusterInfo *ClusterInfo) Execute(args []string) error {
    data, err := request.Get("/api")
    if err != nil {
        return err
    }

    fmt.Printf("%s\n", string(data))
    return nil
}

