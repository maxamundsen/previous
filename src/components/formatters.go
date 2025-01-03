package components

import (
	. "previous/basic"
	"previous/finance"

	. "maragu.dev/gomponents"

	"time"
)

func FormatDateTime(utcTime time.Time) Node {
	return Text(TimeToString(utcTime))
}

func FormatDate(utcTime time.Time) Node {
	return Text(DateToString(utcTime))
}

func FormatMoney(m int64) Node {
	return Text(finance.Int64ToMoney(m))
}

func ToText(i interface{}) Node {
	return Text(ToString(i))
}
