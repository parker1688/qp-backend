package tool

import (
	"math"
	"unicode"
)

func TruncateFloat(num float64, precision int) float64 {
	scale := math.Pow10(precision)
	return math.Trunc(num*scale) / scale
}

func IsDigitsOnly(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return len(s) > 0
}
