package main

import (
    "log"

    "github.com/smart--petea/kubernetes-client/internal/commands"
    "github.com/smart--petea/kubernetes-client/internal/request"
    "github.com/joho/godotenv"
    "github.com/jessevdk/go-flags"
)

type Options struct {
}

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    err = request.InitClient()
    if err != nil {
        log.Fatal(err)
    }

    var options Options
    var parser = flags.NewParser(&options, flags.Default)
    parser.AddCommand("cluster-info", "Display cluster info", "Display cluster info", &commands.ClusterInfo{})
    parser.AddCommand("get-nodes", "Display one or many resources", "Display one or many resources", &commands.GetNodes{})

    if _, err := parser.Parse(); err != nil {
        log.Fatal(err)
    }
}
