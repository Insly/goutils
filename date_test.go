package goutils

import (
	"testing"
	"time"
)

func TestCalculateBusinessDaysBetweenDates(t *testing.T) {
	type args struct {
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 day",
			args: args{
				from: time.Date(2022, time.August, 21, 24, 0, 0, 0, time.UTC),
				to:   time.Date(2022, time.August, 22, 24, 0, 0, 0, time.UTC),
			},
			want: 1,
		},
		{
			name: "day",
			args: args{
				from: time.Date(2022, time.August, 21, 24, 0, 0, 0, time.UTC),
				to:   time.Date(2022, time.August, 22, 23, 59, 59, 0, time.UTC),
			},
			want: 1,
		},
		{
			name: "2 day",
			args: args{
				from: time.Date(2022, time.August, 21, 24, 0, 0, 0, time.UTC),
				to:   time.Date(2022, time.August, 23, 24, 0, 0, 0, time.UTC),
			},
			want: 2,
		},
		{
			name: "3 day",
			args: args{
				from: time.Date(2022, time.August, 21, 24, 0, 0, 0, time.UTC),
				to:   time.Date(2022, time.August, 23, 24, 0, 1, 0, time.UTC),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult := CalculateBusinessDaysBetweenDates(tt.args.from, tt.args.to)
			if gotResult != tt.want {
				t.Errorf("CalculateBusinessDaysBetweenDates() gotResult = %v, want %v", gotResult, tt.want)
			}
		})
	}
}
