package request

import (
    "github.com/smart--petea/kubernetes-client/internal/config"

    "os"
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

type Request struct {
    Method string
    Endpoint string
    Headers map[string]string
}

func InitClient() error {
    conf, err := config.GetConfig()
    if err != nil {
        return err
    }

    clusterName := os.Getenv("CLUSTER_NAME")
    clusterConfig, err  := conf.Clusters.FindByName(clusterName)
    if err != nil {
        return err
    }

    userName := os.Getenv("USER_NAME")
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

func NewRequest(method string, endpoint string) *Request {
    return &Request{
        Method: method,
        Endpoint: endpoint,
        Headers: make(map[string]string),
    }
}

func Get(endpoint string) *Request {
    return NewRequest("GET", endpoint)
}

func (request *Request) AddHeader(name string, value string) *Request {
    oldValue := request.Headers["Accept"]
    if len(oldValue) == 0 {
        request.Headers["Accept"] = value
    } else {
        request.Headers["Accept"] = oldValue + "," + value
    }

    return request
}

func (request *Request) AsJson(group string, version string) *Request {
    request.AddHeader("Accept", fmt.Sprintf("application/json;g=%s;v=%s", group, version))
    return request
}

func (request *Request) AsTable(group string, version string) *Request {
    request.AddHeader("Accept", fmt.Sprintf("application/json;as=Table;g=%s;v=%s", group, version))
    return request
}

func (request *Request) Do() ([]byte, error) {
    client, err := getHttpClient()
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest(request.Method, client.Server + request.Endpoint, nil)
    if err != nil {
        return nil, err
    }

    for headerName, headerValue := range request.Headers {
        fmt.Println(headerValue)
        req.Header.Add(headerName, headerValue)
    }

    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    return ioutil.ReadAll(resp.Body)
}

func GetServerAddress() (string, error) {
    client, err := getHttpClient()
    if err != nil {
        return "", err
    }

    return client.Server, nil
}

func getHttpClient() (*clientT, error) {
    if client == nil {
        return nil, fmt.Errorf("http client is not initialized")
    }

    return client, nil
}
