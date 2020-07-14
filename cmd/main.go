package main

import (
    "log"
    "fmt"

    "github.com/smart--petea/kubernetes-client/internal/request"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    cluster := "minikube" //todo add from bash os.Args
    user := "minikube" //todo add from bash os.Args
    err = request.InitClient(cluster, user)
    if err != nil {
        log.Fatal(err)
    }

    data, err := request.Get("/api")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%s\n", string(data))
}
