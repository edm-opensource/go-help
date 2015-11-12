package arrays

import (
	"fmt"
	"reflect"
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

/*
	TestRemoveDuplicates
	input := []int{1, 1}
	expected := []int{1}
*/
func TestRemoveDuplicates(t *testing.T) {
	input := []int{1, 1}
	expected := []int{1}
	actual := RemoveDuplicateInts(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatError(fmt.Sprintf("%v", input), fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual)))
	}
}

/*
	TestRemoveDuplicates_2
	input := []int{9, 1, 1, 1, 100, 100, 9, 9}
	expected := []int{9, 1, 100}
*/
func TestRemoveDuplicates_2(t *testing.T) {
	input := []int{9, 1, 1, 1, 100, 100, 9, 9}
	expected := []int{9, 1, 100}
	actual := RemoveDuplicateInts(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatError(fmt.Sprintf("%v", input), fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual)))
	}
}

/*
	TestRemoveDuplicates_3
	input := []int{}
	expected := []int{}
*/
func TestRemoveDuplicates_3(t *testing.T) {
	input := []int{}
	expected := []int{}
	actual := RemoveDuplicateInts(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatError(fmt.Sprintf("%v", input), fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual)))
	}
}

func formatError(input interface{}, expected interface{}, actual interface{}) string {
	return fmt.Sprintf("\nInput:    %v\nExpected: %v\nActual:   %v\n", input, expected, actual)
}
func formatErrorWithErr(input interface{}, expected interface{}, actual interface{}, err interface{}) string {
	return fmt.Sprintf("\nInput:    %v\nExpected: %v\nActual:   %v\nError:    %v\n", input, expected, actual, err)
}
