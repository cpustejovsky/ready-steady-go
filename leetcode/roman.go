package leetcode

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

var allRomanNumerals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{49, "XLIX"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

var romanNumerMap = map[string]int{
	"M":    1000,
	"CM":   900,
	"D":    500,
	"CD":   400,
	"C":    100,
	"XC":   90,
	"L":    50,
	"XLIX": 49,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"V":    5,
	"IV":   4,
	"I":    1,
}

func ConvertToRomanNumeral(arabic int) string {

	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

type RomanNumeralMap map[string]int

func (r RomanNumeralMap) ValuesOf(symbols ...byte) int {
	s := string(symbols)
	val, ok := r[s]
	if !ok {
		return 0
	}
	return val
}

func (r RomanNumeralMap) Exists(symbols ...byte) bool {
	if r.ValuesOf(symbols...) == 0 {
		return false
	}
	return true
}

var rn = RomanNumeralMap{
	"M":    1000,
	"CM":   900,
	"D":    500,
	"CD":   400,
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"V":    5,
	"IV":   4,
	"I":    1,
}

func RomanToArabic(roman string) int {
	var result int
	for i := 0; i < len(roman); i++ {
		symbol := roman[i]
		notAtEnd := i+1 < len(roman)
		//If it's not the end of the slice and the value of the current and next symbol exists in the map, add that value
		//then increment the index to skip the next value
		if notAtEnd && rn.Exists(symbol, roman[i+1]) {
			result += rn.ValuesOf(symbol, roman[i+1])
			i++
		} else {
			result += rn.ValuesOf(symbol)
		}

	}
	return result
}
