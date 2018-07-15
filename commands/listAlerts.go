package commands

import (
	"bufio"
	"fmt"
	"os"

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
	alertOcurrences int
	errorOcurrences int
}

// Call OpsGenie list alert api
func (command ListAlertsCommand) Call(cli *ogcli.OpsGenieClient) {
	alertChannel := make(chan alerts.Alert)
	go listAlerts(command, cli, alertChannel)

	report := make(map[string]AlertResult)
	for alert := range alertChannel {
		alertResult, found := report[alert.Alias]
		if !found {
			alertResult = AlertResult{
				alertOcurrences: 0,
				errorOcurrences: 0,
			}
		}
		alertResult.alertOcurrences++
		alertResult.errorOcurrences += alert.Count
		report[alert.Alias] = alertResult
	}
	_, err := print(report)
	if err != nil {
		panic(err)
	}
}

func listAlerts(command ListAlertsCommand, cli *ogcli.OpsGenieClient, alertChannel chan alerts.Alert) {
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
				alertChannel <- alert
			}
			offset = offset + limit
		}
	}
	close(alertChannel)
}

func print(report map[string]AlertResult) (bool, error) {
	file, err := os.Create("report.csv")
	if err != nil {
		return false, err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for k, alertResult := range report {
		_, err := writer.WriteString(fmt.Sprintf("%v,%v,%v\n", k, alertResult.alertOcurrences, alertResult.errorOcurrences))
		if err != nil {
			return false, err
		}
	}
	writer.Flush()
	return true, nil
}

func generateQuery(command ListAlertsCommand) string {
	query := fmt.Sprintf("createdAt>%v", formatDate(command.from))
	if command.to != "" {
		query = fmt.Sprintf("%v AND createdAt<%v", query, formatDate(command.to))
	}
	return query
}
