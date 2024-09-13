package fixer

import (
	"strconv"
)

func BintoDecimal(a string) string {
	decimalValue, err := strconv.ParseInt(a, 2, 64)
	if err != nil {
		return "Invalid binary string"
	}
	return strconv.Itoa(int(decimalValue))
}
