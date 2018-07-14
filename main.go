package main

import (
	"fmt"
	"os"
	"time"

	commands "github.com/filhodomauro/opsgenie-go-integration/commands"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		panic("Command required")
	}

	command, err := commands.Factory(args)
	if err != nil {
		panic(err)
	}
	command.Call(getOpsGenieCli())

	fmt.Printf("Args: %v", args)

	from := fmt.Sprintf("%vT00:00:00Z", args[2])

	t1, err := time.Parse(
		time.RFC3339, from,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data formatada: %v", t1)
}

func getOpsGenieCli() *ogcli.OpsGenieClient {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey("8cb098ae-354c-4edf-903c-3cac233a5891")
	return cli
}
