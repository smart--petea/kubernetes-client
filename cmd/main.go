package main

import (
    "log"

    "github.com/smart--petea/kubernetes-client/internal/commands"
    "github.com/smart--petea/kubernetes-client/internal/request"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    err = request.InitClient()
    if err != nil {
        log.Fatal(err)
    }

    commands.
        NewRootCommand().
        Execute()
}
