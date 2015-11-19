package gohelp

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	m0 = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	m1 = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	m2 = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	m3 = []string{"", "M", "MM", "MMM", "I̅V̅",
		"V̅", "V̅I̅", "V̅I̅I̅", "V̅I̅I̅I̅", "I̅X̅"}
	m4 = []string{"", "X̅", "X̅X̅", "X̅X̅X̅", "X̅L̅",
		"L̅", "L̅X̅", "L̅X̅X̅", "L̅X̅X̅X̅", "X̅C̅"}
	m5 = []string{"", "C̅", "C̅C̅", "C̅C̅C̅", "C̅D̅",
		"D̅", "D̅C̅", "D̅C̅C̅", "D̅C̅C̅C̅", "C̅M̅"}
	m6 = []string{"", "M̅", "M̅M̅", "M̅M̅M̅"}
)

var mapRomanToNumbers = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// EncodeRomanNumerals function
func EncodeRomanNumerals(n int) (string, bool) {
	if n < 1 || n > 4e6 {
		return "", false
	}

	return m6[n/1e6] + m5[n%1e6/1e5] + m4[n%1e5/1e4] + m3[n%1e4/1e3] +
		m2[n%1e3/1e2] + m1[n%100/10] + m0[n%10], true
}

// DecodeRomanNumerals function
func DecodeRomanNumerals(s string) (int, bool) {
	result := 0
	for index := 0; index < len(s)-1; index++ {
		s1 := string(s[index])
		s2 := string(s[index+1])
		if mapRomanToNumbers[s1] < mapRomanToNumbers[s2] {
			result -= mapRomanToNumbers[s1]
		} else {
			result += mapRomanToNumbers[s1]
		}
	}
	s3 := string(s[len(s)-1])
	result += mapRomanToNumbers[s3]
	return result, true
}

/*
TitleFirstWord bra skit va
*/
func TitleFirstWord(s string) string {
	i := 0
	w := 0
	for ; i < len(s); i += w {
		r, width := utf8.DecodeRuneInString(s[i:])
		w = width
		if unicode.IsLetter(r) {
			break
		}
	}
	res := s[0:i] + strings.ToUpper(string([]rune(s)[i])) + s[i+w:]
	return res
}
