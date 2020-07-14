package main

import (
    "log"
    "fmt"
    "os"

    "github.com/smart--petea/kubernetes-client/internal/request"
    "github.com/joho/godotenv"
)

func main() {
    if len(os.Args) == 1 {
        log.Fatal("xxx")
    }

    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    err = request.InitClient()
    if err != nil {
        log.Fatal(err)
    }

    data, err := request.Get("/api")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%s\n", string(data))
}
