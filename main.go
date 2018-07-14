package main

import (
	"fmt"
	"strconv"

	alerts "github.com/opsgenie/opsgenie-go-sdk/alertsv2"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey("8cb098ae-354c-4edf-903c-3cac233a5891")

	alertCli, _ := cli.AlertV2()

	response, err := alertCli.List(alerts.ListAlertRequest{
		Limit:                100,
		Offset:               0,
		SearchIdentifierType: alerts.Name,
	})

	if err != nil {
		panic(err)
	} else {
		for _, alert := range response.Alerts {
			fmt.Println(alert.Alias + "," + strconv.Itoa(alert.Count))
		}
	}
}
