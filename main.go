package main

import (
	"os"

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
}

func getOpsGenieCli() *ogcli.OpsGenieClient {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(os.Getenv("OPSGENIE_API_KEY"))
	return cli
}
