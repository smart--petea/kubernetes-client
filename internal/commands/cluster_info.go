package commands

import (
    "fmt"
    "github.com/smart--petea/kubernetes-client/internal/request"
    "encoding/json"

    ct "github.com/daviddengcn/go-colortext"
)

type ServiceList struct {
    Kind string `json:"kind"`
    ApiVersion string `json:"apiVersion"`
    Metadata struct {
        SelfLink string `json:"selfLink"`
        ResourceVersion string `json:"resourceVersion"`
    } `json:"metadata"`
    Items []struct{
        Metadata struct {
            Name string `json:"name"`
            Namespace string `json:"namespace"`
            SelfLink string `json:"selfLink"`
            Uid string `json:"uid"`
            ResourceVersion string `json:"resourceVersion"`
            CreationTimestamp string `json:"creationTimestamp"`
            Labels map[string]string `json:"labels"`
            Annotations map[string]interface{} `json:"annotations"`
            ManagedFields []struct{
                Manager string `json:"manager"`
                Operation string `json:"operation"`
                ApiVersion string `json:"apiVersion"`
                Time string `json:"time"`
                FieldsType string `json:"fieldsType"`
                FieldsV1 struct {
                    FMetadata struct {
                        FAnnotations map[string] map[string] interface{} `json:"f:annotations"`
                        FLabels map[string] map[string] interface{} `json:"f:labels"`
                    } `json:"f:metadata"`
                    FSpec map[string] interface{} `json:"f:spec"`
                } `json:"fieldsV1"`
            } `json:"managedFields"`
        } `json:"metadata"`
        Spec struct {
            Ports []struct{
                Name string `json:"name"`
                Protocol string `json:"protocol"`
                Port int `json:"port"`
                TargetPort int `json:"targetPort"`
            } `json:"ports"`
            Selector struct {
                K8sApp string `json:"k8s-app"`
            } `json:"selector"`
            ClusterIP string `json:"type"`
            Type string `json:"type"`
            SessionAffinity string `json:"sessionAffinity"`
        } `json:"spec"`
        Status struct {
            LoadBalancer struct {
            } `json:"loadBalancer"`
        } `json:"status"`
    } `json:"items"`
}


type ClusterInfo struct {
}

func (clusterInfo *ClusterInfo) Execute(args []string) error {
    serverAddress, err := request.GetServerAddress()
    if err != nil {
        return err
    }
    printOutput("Kubernetes master", " is running at ", serverAddress)

    data, err := request.
        Get("/api/v1/namespaces/kube-system/services?l").
        AsJson("meta.k8s.io", "v1").
        Do()
    if err != nil {
        return err
    }

    var serviceList ServiceList
    err = json.Unmarshal(data, &serviceList)
    if err != nil {
        return err
    }

    for _, service := range serviceList.Items {
        printOutput(service.Metadata.Labels["kubernetes.io/name"], " is running at ", serverAddress + service.Metadata.SelfLink)
    }

    printOutput("", "", "")
    printOutput("", "To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.", "")

    return nil
}

func printOutput(green string, white string, yellow string) {
    ct.ChangeColor(ct.Green, false, ct.None, false) 
    fmt.Print(green)

    ct.ResetColor()
    fmt.Print(white)

    ct.ChangeColor(ct.Yellow, false, ct.None, false) 
    fmt.Println(yellow)
}
