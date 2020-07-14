package request

import (
    "github.com/smart--petea/kubernetes-client/internal/config"

    "net/http"
    "crypto/tls"
    "crypto/x509"
    "io/ioutil"
    "fmt"
)

type clientT struct {
    *http.Client
    Server string
}

var client *clientT

func InitClient( clusterName string, userName string) error {
    conf, err := config.GetConfig()
    if err != nil {
        return err
    }

    clusterConfig, err  := conf.Clusters.FindByName(clusterName)
    if err != nil {
        return err
    }

    userConfig, err := conf.Users.FindByName(userName)
    if err != nil {
        return err
    }

    cert, err := tls.LoadX509KeyPair(
        userConfig.Meta.ClientCertificate,
        userConfig.Meta.ClientKey,
    )
    if err != nil {
        return err
    }

    caCert, err := ioutil.ReadFile(clusterConfig.Meta.CertificateAuthority)
    if err != nil {
        return err
    }
    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)

    tlsConfig := &tls.Config{
        Certificates: []tls.Certificate{cert},
        RootCAs: caCertPool,
    }
    tlsConfig.BuildNameToCertificate()
    transport := &http.Transport{TLSClientConfig: tlsConfig}

    client = &clientT{
        Client: &http.Client{Transport: transport},
        Server: clusterConfig.Meta.Server,
    }

    return nil
}

func Get(endpoint string) ([]byte, error) {
    client, err := getHttpClient()
    if err != nil {
        return nil, err
    }

    resp, err := client.Get(client.Server + endpoint)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    return ioutil.ReadAll(resp.Body)
}

func getHttpClient() (*clientT, error) {
    if client == nil {
        return nil, fmt.Errorf("http client is not initialized")
    }

    return client, nil
}
