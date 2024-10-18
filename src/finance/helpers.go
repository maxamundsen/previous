package finance

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Int64ToMoney(value int64) string {
	decimalValue := float64(value) / 100.0
	moneyString := fmt.Sprintf("%.2f", decimalValue)
	moneyString = strings.ReplaceAll(moneyString, ",", ".")
	return moneyString
}

func MoneyToInt64(input string) int64 {
	if !strings.Contains(input, ".") {
		input += ".00"
	} else if strings.Count(input, ".") == 1 {
		digits := strings.Split(input, ".")
		if len(digits[1]) == 1 {
			input += "0"
		}
	}
	processedString := strings.ReplaceAll(input, ".", "")
	processedString = strings.ReplaceAll(processedString, ",", "")
	processedString = strings.ReplaceAll(processedString, " ", "")
	result, _ := strconv.Atoi(processedString)
	return int64(result)
}

func RoundUpToCeiling(input int64) int64 {
	remainder := input % 100
	if remainder == 0 {
		return input
	}
	return input + 100 - remainder
}

func RoundDownToFloor(input int64) int64 {
	remainder := input % 100
	if remainder == 0 {
		return input
	}
	return input - remainder
}

func MultiplyByPercentageS64(amount int64, percentage float64) int64 {
	amount_float64 := float64(amount) / 100
	out_float64 := (amount_float64 * (percentage / 100))

	out_int64 := int64(math.Round(out_float64 * 100))

	return out_int64
}

func MultiplyByPercentageF64(amount int64, percentage float64) float64 {
	amount_float64 := float64(amount) / 100

	return amount_float64 * percentage
}

func ProcessDiscount(discount string) float64 {
	discount_f64, _ := strconv.ParseFloat(discount, 64)

	if discount_f64 < 0 {
		return 0
	} else if discount_f64 > 100 {
		return 100
	}

	return discount_f64
}
