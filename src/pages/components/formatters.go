package components

import (
	"webdawgengine/finance"

	"fmt"
	. "maragu.dev/gomponents"
	"reflect"
	"time"
)

func FormatDateTime(utcTime time.Time) Node {
	// Convert to EST (Eastern Standard Time)
	loc, _ := time.LoadLocation("America/New_York")
	estTime := utcTime.In(loc)

	// Format the time as mm/dd/yy hh:mm AM/PM
	formattedTime := estTime.Format("01/02/06 03:04 PM")

	return Text(formattedTime)
}

func FormatDate(utcTime time.Time) Node {
	// Convert to EST (Eastern Standard Time)
	loc, _ := time.LoadLocation("America/New_York")
	estTime := utcTime.In(loc)

	// Format the time as mm/dd/yy hh:mm AM/PM
	formattedTime := estTime.Format("01/02/06")

	return Text(formattedTime)
}

func FormatMoney(m int64) Node {
	return Text(finance.Int64ToMoney(m))
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

func ToText(i interface{}) Node {
	return Text(ToString(i))
}
