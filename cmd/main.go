package main

import (
    "gopkg.in/yaml.v2"

    "io/ioutil"
    "log"
    "fmt"
    "net/http"

    "crypto/tls"
    "crypto/x509"

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

    //fmt.Printf("%+v\n", config)

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

    //fmt.Printf("%+v", userConfig)

    clientKey := userConfig.Meta.ClientKey
    server := clusterConfig.Meta.Server
    certificateAuthority := clusterConfig.Meta.CertificateAuthority
    clientCertificate := userConfig.Meta.ClientCertificate
    fmt.Printf("server=%s\n", server)
    fmt.Printf("clientKey=%s\n", clientKey)
    fmt.Printf("certificateAuthority=%s\n", certificateAuthority)
    fmt.Printf("clientCertificate=%s\n", clientCertificate)

    cert, err := tls.LoadX509KeyPair(clientCertificate, clientKey)
    if err != nil {
        log.Fatal(err)
    }

    caCert, err := ioutil.ReadFile(certificateAuthority)
    if err != nil {
        log.Fatal(err)
    }
    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)

    tlsConfig := &tls.Config{
        Certificates: []tls.Certificate{cert},
        RootCAs: caCertPool,
    }
    tlsConfig.BuildNameToCertificate()
    transport := &http.Transport{TLSClientConfig: tlsConfig}
    client := &http.Client{Transport: transport}

    resp, err := client.Get(server + "/api")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    log.Println(string(data))
}
