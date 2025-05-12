package property

import (
	"strings"
)

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic uint16) string {
	var sb strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			sb.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return sb.String()
}

func ConvertToArabic(roman string) uint16 {
	var arabic uint16

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}

	return arabic
}

func Recursive(roman string) uint16 {
	var arabic uint16
	var newRoman string

	for _, numeral := range allRomanNumerals {
		if strings.HasPrefix(roman, numeral.Symbol) {
			if len(roman) == len(numeral.Symbol) {
				return numeral.Value
			}

			arabic = numeral.Value
			newRoman = strings.TrimPrefix(roman, numeral.Symbol)
			break
		}
	}

	return arabic + Recursive(newRoman)
}
