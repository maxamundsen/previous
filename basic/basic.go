// General purpose "utilities" that act as my own "standard library"
package basic

import (
	"fmt"
	"reflect"
	"time"
)

func ToString(i interface{}) string {
	v := reflect.ValueOf(i)

	output := ""

	switch v.Kind() {
	case reflect.String:
		output = v.String()
	case reflect.Int:
		output = fmt.Sprintf("%d", v.Int())
	case reflect.Float64:
		output = fmt.Sprintf("%f", v.Float())
	case reflect.Bool:
		output = fmt.Sprintf("%t", v.Bool())
	default:
		output = fmt.Sprintf("%v", i)
	}

	return output
}

func TimeToString(utcTime time.Time) string {
	// Convert to EST (Eastern Standard Time)
	loc, _ := time.LoadLocation("America/New_York")
	estTime := utcTime.In(loc)

	// Format the time as mm/dd/yy hh:mm AM/PM
	return estTime.Format("01/02/06 03:04 PM")
}

func DateToString(utcTime time.Time) string {
	// Convert to EST (Eastern Standard Time)
	loc, _ := time.LoadLocation("America/New_York")
	estTime := utcTime.In(loc)

	// Format the time as mm/dd/yy hh:mm AM/PM
	return estTime.Format("01/02/06")
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func RemoveDuplicates[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
