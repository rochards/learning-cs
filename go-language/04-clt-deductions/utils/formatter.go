package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func FormatDecimalInBRL(value float64) string {

	parts := strings.Split(fmt.Sprintf("%.2f", value), ".")
	integerPart := parts[0]
	decimalPart := parts[1]

	var result strings.Builder
	integerPartLength := len(integerPart)
	for i, digit := range integerPart {
		if i > 0 && (integerPartLength-i)%3 == 0 {
			result.WriteString(".")
		}
		result.WriteRune(digit)
	}

	result.WriteString(",")
	result.WriteString(decimalPart)

	return result.String()
}

func FormatDecimalInBRLWithPrecision(value float64, precision int) string {

	parts := strings.Split(fmt.Sprintf("%."+strconv.Itoa(precision)+"f", value), ".")
	integerPart := parts[0]
	decimalPart := parts[1]

	var result strings.Builder
	integerPartLength := len(integerPart)
	for i, digit := range integerPart {
		if i > 0 && (integerPartLength-i)%3 == 0 {
			result.WriteString(".")
		}
		result.WriteRune(digit)
	}

	result.WriteString(",")
	result.WriteString(decimalPart)

	return result.String()
}
