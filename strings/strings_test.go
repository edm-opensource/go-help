package strings

import (
	"fmt"
	"testing"
)

func formatError(input interface{}, expected interface{}, actual interface{}) string {
	//return "Input: " + input.(string) + "\nExpected: " + expected.(string) + "\nActual: " + actual.(string)
	return fmt.Sprintf("\nInput:    %v\nExpected: %v\nActual:   %v", input, expected, actual)
}

type testpairRoman struct {
	value int
	roman string
}

func TestEncodeRomanNumeralsCorrect(t *testing.T) {
	tests := []testpairRoman{
		{1990, "MCMXC"},
		{2008, "MMVIII"},
		{1992, "MCMXCII"},
		{1666, "MDCLXVI"},
		{1700, "MDCC"},
		{1995, "MCMXCV"},
	}
	for _, pair := range tests {
		roman, ok := EncodeRomanNumerals(pair.value)
		if !ok || roman != pair.roman {
			t.Error(formatError(pair.value, pair.roman, roman))
		}
	}
}

func TestDecodeToRomanNumeralsCorrect(t *testing.T) {
	tests := []testpairRoman{
		{1990, "MCMXC"},
		{2008, "MMVIII"},
		{1992, "MCMXCII"},
		{1666, "MDCLXVI"},
		{1700, "MDCC"},
		{1995, "MCMXCV"},
	}
	for _, pair := range tests {
		number, ok := DecodeRomanNumerals(pair.roman)
		if !ok || number != pair.value {
			t.Error(formatError(pair.roman, pair.value, number))
		}
	}
}

// TestToRomanNumeralsInCorrect function
func TestToRomanNumeralsInCorrect(t *testing.T) {
	_, ok := EncodeRomanNumerals(0)
	if ok {
		t.Error(formatError(0, "false and empty string", ok))
	}
}

func TestTitleFirstWord_1(t *testing.T) {
	if TitleFirstWord("hello world") != "Hello world" {
		t.Error("hello world was not parsed sucessfully")
	}
}
func TestTitleFirstWord_2(t *testing.T) {
	if TitleFirstWord("Hello world") != "Hello world" {
		t.Error("Hello world was not parsed sucessfully")
	}
}
