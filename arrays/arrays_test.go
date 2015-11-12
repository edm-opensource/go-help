package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

/*
TestReverseStrings_1
input := []string{"one", "two", "three"}
expected := []string{"three", "two", "one"}
*/
func TestReverseStrings_1(t *testing.T) {
	input := []string{"one", "two", "three"}
	expected := []string{"three", "two", "one"}
	actual := ReverseStrings(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Error(formatError(input, expected, actual))
	}
}

/*
TestReverseStrings_2
input := []string{"one"}
expected := []string{"one"}
*/
func TestReverseStrings_2(t *testing.T) {
	input := []string{"one"}
	expected := []string{"one"}
	actual := ReverseStrings(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Error(formatError(input, expected, actual))
	}
}

/*
TestReverseStrings_3
var input []string
var expected []string
*/
func TestReverseStrings_3(t *testing.T) {
	var input []string
	var expected []string
	actual := ReverseStrings(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Error(formatError(input, expected, actual))
	}
}

func formatError(input interface{}, expected interface{}, actual interface{}) string {
	return fmt.Sprintf("\nInput:    %v\nExpected: %v\nActual:   %v\n", input, expected, actual)
}
func formatErrorWithErr(input interface{}, expected interface{}, actual interface{}, err interface{}) string {
	return fmt.Sprintf("\nInput:    %v\nExpected: %v\nActual:   %v\nError:    %v\n", input, expected, actual, err)
}
