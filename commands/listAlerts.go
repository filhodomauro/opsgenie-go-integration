package commands

import (
	"fmt"

	alerts "github.com/opsgenie/opsgenie-go-sdk/alertsv2"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
)

const limit int = 100

// ListAlertsCommand is a List Alert Command
type ListAlertsCommand struct {
	from string
	to   string
}

type AlertResult struct {
	alias string
}

// Call OpsGenie list alert api
func (command ListAlertsCommand) Call(cli *ogcli.OpsGenieClient) {
	alertCli, _ := cli.AlertV2()

	query := generateQuery(command)

	hasMoreResult := true
	offset := 0

	for hasMoreResult {
		response, err := alertCli.List(alerts.ListAlertRequest{
			Limit:                limit,
			Offset:               offset,
			SearchIdentifierType: alerts.Name,
			Sort:                 alerts.CreatedAt,
			Order:                "desc",
			Query:                query,
		})

		if err != nil {
			panic(err)
		} else {
			hasMoreResult = len(response.RateLimitState) == limit
			for _, alert := range response.Alerts {
				fmt.Println(fmt.Sprintf("%v,%v,%v", alert.CreatedAt, alert.Alias, alert.Count))
			}
			offset = offset + limit
		}
	}
}

func generateQuery(command ListAlertsCommand) string {
	query := fmt.Sprintf("createdAt>%v", formatDate(command.from))
	if command.to != "" {
		query = fmt.Sprintf("%v AND createdAt<%v", query, formatDate(command.to))
	}
	return query
}
