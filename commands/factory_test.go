package commands

import (
	"reflect"
	"testing"
)

func Test_formatDate(t *testing.T) {
	type args struct {
		from string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "convert date 10/07/1991",
			args: args{
				from: "1991-07-10",
			},
			want: 679104000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatDate(tt.args.from); got != tt.want {
				t.Errorf("formatDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactory(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		want    Command
		wantErr bool
	}{
		{
			name: "Testing valid list-alerts start end date",
			args: args{
				args: []string{"execution-info", "list-alerts", "2018-01-30", "2018-01-31"},
			},
			want: ListAlertsCommand{
				from: "2018-01-30",
				to:   "2018-01-31",
			},
			wantErr: false,
		},
		{
			name: "Testing valid list-alerts only start date",
			args: args{
				args: []string{"execution-info", "list-alerts", "2018-01-30"},
			},
			want: ListAlertsCommand{
				from: "2018-01-30",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factory(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Factory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Factory() = %v, want %v", got, tt.want)
			}
		})
	}
}
