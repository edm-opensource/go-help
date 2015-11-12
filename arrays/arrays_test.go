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
