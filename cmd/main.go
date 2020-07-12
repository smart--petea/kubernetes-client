package main

import (
    "gopkg.in/yaml.v2"

    "io/ioutil"
    "log"
    "fmt"

    "github.com/smart--petea/kubernetes-client/internal/config"
)

func main() {
    configPath := "/home/check/.kube/config"
    configBytes, err := ioutil.ReadFile(configPath)
    if err != nil {
        log.Fatal(err)
    }

    var config config.ConfigT
    err = yaml.Unmarshal(configBytes, &config)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%+v\n", config)

    clusterName := "minikube" //todo add from bash os.Args
    clusterConfig, err  := config.Clusters.FindByName(clusterName)
    if err != nil {
        log.Fatal(err)
    }

    userName := "minikube" //todo add from bash os.Args
    userConfig, err := config.Users.FindByName(userName)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%+v", userConfig)

    fmt.Printf("%+v", clusterConfig.Meta.Server)
}
