package finance

import (
	"testing"
)

func TestTax(t *testing.T) {
	var a int64 = 15345
	var b float64 = 8.625

	c := MultiplyByPercentageS64(a, b)

	const expected int64 = 1324
	if c == expected {
		m := Int64ToMoney(c)
		t.Logf("output: $%s", m)
	} else {
		t.Fatalf("assert failed. expected %d, got %d", expected, c)
	}
}
