// General purpose "utilities" that act as my own "standard library"
package basic

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type StrPair struct {
	Key   string
	Value string
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func MakeURLParams(base string, params ...StrPair) string {
	output := base

	for i, v := range params {
		if i == 0 {
			output += "?" + v.Key + "=" + v.Value
		} else {
			output += "&" + v.Key + "=" + v.Value
		}
	}

	return output
}

func SnakeCaseToTitleCase(s string) string {
	parts := strings.Split(s, "_")

	for i, part := range parts {
		parts[i] = strings.Title(part)
	}

	return strings.Join(parts, " ")
}

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

func HTMLDateToTime(date string) time.Time {
	t, _ := time.Parse("2006-01-02", date)
	return t
}

func TimeToSqliteString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func SqliteStringToTime(dateTimeStr string) time.Time {
	formats := []string{
		"2006-01-02 15:04:05.000000000",
		"2006-01-02 15:04:05",
	}

	var t time.Time
	var err error

	for _, format := range formats {
		t, err = time.Parse(format, dateTimeStr)
		if err == nil {
			return t
		}
	}

	return t
}

func TimeToTimeString(utcTime time.Time) string {
	return utcTime.Format("03:04 PM")
}

func TimeToString(utcTime time.Time) string {
	return utcTime.Format("01/02/06 03:04 PM")
}

func DateToString(utcTime time.Time) string {
	return utcTime.Format("01/02/06")
}

func StringToDate(ds string) time.Time {
	nt, _ := time.Parse("01/02/06", ds)
	return nt
}

func Reverse[T comparable](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
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