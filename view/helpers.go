package view

import (
	"reflect"
	"time"
)

func Empty(item interface{}) bool {
	val := reflect.ValueOf(item)
	return val.Len() < 1
}

func DateFormat(format string, location *time.Location) func(time.Time) string {
	return func(date time.Time) string {
		return date.In(location).Format(format)
	}
}
