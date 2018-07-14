package commands

import (
	"fmt"

	alerts "github.com/opsgenie/opsgenie-go-sdk/alertsv2"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
)

// ListAlertsCommand is a List Alert Command
type ListAlertsCommand struct {
	from string
	to   string
}

// Call OpsGenie list alert api
func (command ListAlertsCommand) Call(cli *ogcli.OpsGenieClient) {
	alertCli, _ := cli.AlertV2()

	response, err := alertCli.List(alerts.ListAlertRequest{
		Limit:                2,
		Offset:               0,
		SearchIdentifierType: alerts.Name,
		Sort:                 alerts.CreatedAt,
		Order:                "desc",
	})

	if err != nil {
		panic(err)
	} else {
		for _, alert := range response.Alerts {
			fmt.Println(fmt.Printf("%v,%v,%v", alert.CreatedAt, alert.Alias, alert.Count))
		}
	}
}
