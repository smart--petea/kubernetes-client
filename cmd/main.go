package main

import (
    "log"

    "github.com/smart--petea/kubernetes-client/internal/commands"
    "github.com/smart--petea/kubernetes-client/internal/request"
    "github.com/joho/godotenv"
    "github.com/spf13/cobra"
    "fmt"
)

var rootCmd = &cobra.Command{
    Use:  "use",
    Short: "short", 
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("ok")
    },
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

    commands.Get.AddCommand(commands.GetNodes)
    rootCmd.AddCommand(commands.ClusterInfo)
    rootCmd.AddCommand(commands.Get)
    rootCmd.Execute()

    /*
    var options Options
    var parser = flags.NewParser(&options, flags.Default)
    parser.AddCommand("cluster-info", "Display cluster info", "Display cluster info", &commands.ClusterInfo{})
    parser.AddCommand("get-nodes", "Display one or many resources", "Display one or many resources", &commands.GetNodes{})

    if _, err := parser.Parse(); err != nil {
        log.Fatal(err)
    }
    */
}
