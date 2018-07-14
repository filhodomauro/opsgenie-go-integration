# OpsGenie Go Integration

## Requirements

Generate an [OpsGenie Integration](https://app.opsgenie.com/integration#/) to use as environment variable
```
export OPSGENIE_API_KEY={opsgenie_api_key}
``` 

## Run

### List alerts

List alerts in an interval using the opsgenie createdAt attribute

```
go run main.go list-alerts {start_date} {end_date}(optional)
```

**pattern date**: yyyy-MM-dd. *Ex: 2018-07-10*

**interval rule**: >= start_date AND < end_date