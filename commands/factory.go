package commands

import (
	"fmt"
	"time"

	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
)

// Command represents the api command
type Command interface {
	Call(cli *ogcli.OpsGenieClient)
}

// Factory creates an OpsGenie command from args
func Factory(args []string) (Command, error) {
	commandName := args[1]
	switch commandName {
	case "list-alerts":
		if len(args) < 3 {
			return nil, fmt.Errorf("Argument [from] required. Ex: ... list-alerts yyyy-MM-dd")
		}
		command := ListAlertsCommand{
			from: args[2],
		}
		if len(args) > 3 {
			command.to = args[3]
		}
		return command, nil
	default:
		return nil, fmt.Errorf("Invalid command %v", commandName)
	}
}

func formatDate(from string) int64 {
	parsed, err := time.Parse(
		time.RFC3339, fmt.Sprintf("%vT00:00:00Z", from),
	)
	if err != nil {
		panic(err)
	}
	return parsed.UnixNano() / 1000000
}
