package main

import (
    "gopkg.in/yaml.v2"

    "io/ioutil"
    "log"
    "fmt"
)

type ClusterConfig struct {
    Cluster struct {
        CertificateAuthority string
        Server string
    }
    Name string
}

type Config struct {
    Clusters []ClusterConfig
}

func (config Config) getClusterByName(name string) (clusterConfig *ClusterConfig, err error) {
    for _, cluster := range config.Clusters {
        if cluster.Name == name {
            return &cluster, err
        }
    }

    return nil, fmt.Errorf("Cluster with name %s not found", name)
}

func main() {
    configPath := "/home/check/.kube/config"
    configBytes, err := ioutil.ReadFile(configPath)
    if err != nil {
        log.Fatal(err)
    }

    var config Config
    err = yaml.Unmarshal(configBytes, &config)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%+v\n", config)
    clusterConfig, err  := config.getClusterByName("minikube")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%+v", clusterConfig.Cluster.Server)
}
