package arrays

import (
	"fmt"
	"testing"
)

/*
TestReverseString
input: [3]string{"one", "two", "three"}
expected: [3]string{"three", "two", "one"}
error: no
*/
func TestReverseString(t *testing.T) {
	input := []string{"one", "two", "three"}
	expected := []string{"three", "two", "one"}
	actual, err := ReverseString(input)
	if actual[0] != expected[0] && err == nil {
		t.Error(formatErrorWithErr(input, fmt.Sprintf("%v", expected)+", err == <nil>", actual, err))
	}
}

/*
TestReverseString_3
input := [1]string{"one"}
expected: [3]string{"three", "two", "one"}
error: yes
*/
func TestReverseString_3(t *testing.T) {
	input := []string{"one"}
	expected := []string{"one"}
	actual, err := ReverseString(input)
	if actual[0] != expected[0] && err == nil {
		t.Error(formatErrorWithErr(fmt.Sprintf("%v", input),
			fmt.Sprintf("%v", expected)+", err == <nil>",
			fmt.Sprintf("%v", actual),
			fmt.Sprintf("%v", err)))
	}
}

func formatError(input interface{}, expected interface{}, actual interface{}) string {
	return fmt.Sprintf("\nInput:    %v\nExpected: %v\nActual:   %v\n", input, expected, actual)
}
func formatErrorWithErr(input interface{}, expected interface{}, actual interface{}, err interface{}) string {
	return fmt.Sprintf("\nInput:    %v\nExpected: %v\nActual:   %v\nError:    %v\n", input, expected, actual, err)
}
