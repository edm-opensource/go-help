package slices

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
TestReverseInts_1
input := []int{1, 2, 3}
expected := []int{3, 2, 1}
*/
func TestReverseInts_1(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{3, 2, 1}
	actual := ReverseInts(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Error(formatError(input, expected, actual))
	}
}

/*
TestReverseInts_2
input := []int{1}
expected := []int{1}
*/
func TestReverseInts_2(t *testing.T) {
	input := []int{1}
	expected := []int{1}
	actual := ReverseInts(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Error(formatError(input, expected, actual))
	}
}

/*
TestReverseInts_3
var input []int
var expected []int
*/
func TestReverseInts_3(t *testing.T) {
	var input []int
	var expected []int
	actual := ReverseInts(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Error(formatError(input, expected, actual))
	}
}

/*
TestReverseFloat64s_1
input := []float64{1.123, 2.3, 3.12312}
expected := []float64{3.12312, 2.3, 1.123}
*/
func TestReverseFloat64s_1(t *testing.T) {
	input := []float64{1.123, 2.3, 3.12312}
	expected := []float64{3.12312, 2.3, 1.123}
	actual := ReverseFloat64s(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Error(formatError(input, expected, actual))
	}
}

/*
TestReverseFloat64s_2
input := []float64{1.451}
expected := []float64{1.451}
*/
func TestReverseFloat64s_2(t *testing.T) {
	input := []float64{1.451}
	expected := []float64{1.451}
	actual := ReverseFloat64s(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Error(formatError(input, expected, actual))
	}
}

/*
TestReverseFloat64s_3
var input []float64
var expected []float64
*/
func TestReverseFloat64s_3(t *testing.T) {
	var input []float64
	var expected []float64
	actual := ReverseFloat64s(input)
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

/*
	TestRemoveDuplicates
	input := []int{1.0, 1.0}
	expected := []int{1.0}
*/
func TestRemoveDuplicatesFloat(t *testing.T) {
	input := []float64{1.0, 1.0}
	expected := []float64{1.0}
	actual := RemoveDuplicateFloat64(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatError(fmt.Sprintf("%v", input), fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual)))
	}
}

/*
	TestRemoveDuplicates_2
	input := []int{9, 1, 1, 1, 100, 100, 9, 9}
	expected := []int{9, 1, 100}
*/
func TestRemoveDuplicatesFloat_2(t *testing.T) {
	input := []float64{9.0, 1.1, 1.0, 1.0, 100, 100, 9.0, 9.0}
	expected := []float64{9.0, 1.1, 1.0, 100}
	actual := RemoveDuplicateFloat64(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatError(fmt.Sprintf("%v", input), fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual)))
	}
}

/*
	TestRemoveDuplicates_3
	input := []int{}
	expected := []int{}
*/
func TestRemoveDuplicatesFloat_3(t *testing.T) {
	input := []float64{}
	expected := []float64{}
	actual := RemoveDuplicateFloat64(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatError(fmt.Sprintf("%v", input), fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual)))
	}
}

/*
	TestRemoveDuplicates_3
	input := []int{}
	expected := []int{}
*/
func TestRemoveDuplicatesFloat_4(t *testing.T) {
	input := []float64{-1, -100, -34.55, -100}
	expected := []float64{-1, -100, -34.55}
	actual := RemoveDuplicateFloat64(input)

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
