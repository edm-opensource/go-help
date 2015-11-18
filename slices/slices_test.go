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
	TestRemoveAllInstancesOfStrings_1
	input := []string{"one", "one", "two", "oone"}
	input2 := []string{"one"}
	expected := []string{"two", "oone"}
*/
func TestRemoveAllInstancesOfStrings_1(t *testing.T) {
	input := []string{"one", "one", "two", "oone"}
	input2 := []string{"one"}
	expected := []string{"two", "oone"}
	actual := RemoveAllInstancesOfStrings(input, input2...)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

/*
	TestRemoveDuplicatesFloat
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
	TestRemoveAllInstancesOfStrings_2
	input := []string{"one", "\none", "two", "oone"}
	input2 := []string{"one"}
	expected := []string{"\none", "two", "oone"}
*/
func TestRemoveAllInstancesOfStrings_2(t *testing.T) {
	input := []string{"one", "\none", "two", "oone"}
	input2 := []string{"one"}
	expected := []string{"\none", "two", "oone"}
	actual := RemoveAllInstancesOfStrings(input, input2...)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

/*
	TestRemoveDuplicatesFloat_2
	input := []float64{9.0, 1.1, 1.0, 1.0, 100, 100, 9.0, 9.0}
	expected := []float64{9.0, 1.1, 1.0, 100}
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
	TestRemoveAllInstancesOfStrings_3
	input := []string{"one", "one", "two", "oone"}
	input2 := []string{""}
	expected := []string{"one", "one", "two", "oone"}
*/
func TestRemoveAllInstancesOfStrings_3(t *testing.T) {
	input := []string{"one", "one", "two", "oone"}
	input2 := []string{""}
	expected := []string{"one", "one", "two", "oone"}
	actual := RemoveAllInstancesOfStrings(input, input2...)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

/*
	TestRemoveDuplicatesFloat_3
	input := []float64{}
	expected := []float64{}
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
	TestRemoveAllInstancesOfInts_1
	input := []int{1, 2, 3, 2, 2, 4, 2, 6, 12}
	input2 := []int{2, 1}
	expected := []int{3, 4, 6, 12}
*/
func TestRemoveAllInstancesOfInts_1(t *testing.T) {
	input := []int{1, 2, 3, 2, 2, 4, 2, 6, 12}
	input2 := []int{2, 1}
	expected := []int{3, 4, 6, 12}
	actual := RemoveAllInstancesOfInts(input, input2...)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

/*
	TestRemoveDuplicatesFloat_4
	input := []float64{-1, -100, -34.55, -100}
	expected := []float64{-1, -100, -34.55}
*/
func TestRemoveDuplicatesFloat_4(t *testing.T) {
	input := []float64{-1, -100, -34.55, -100}
	expected := []float64{-1, -100, -34.55}
	actual := RemoveDuplicateFloat64(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatError(fmt.Sprintf("%v", input), fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual)))
	}
}

/*
	TestRemoveAllInstancesOfInts_2
	input := []int{1, 2, 3, 2, 2, 4, 2, 6, 12}
	input2 := []int{16}
	expected := []int{1, 2, 3, 2, 2, 4, 2, 6, 12}
*/
func TestRemoveAllInstancesOfInts_2(t *testing.T) {
	input := []int{1, 2, 3, 2, 2, 4, 2, 6, 12}
	input2 := []int{16}
	expected := []int{1, 2, 3, 2, 2, 4, 2, 6, 12}
	actual := RemoveAllInstancesOfInts(input, input2...)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

/*
	TestRemoveDuplicatesString
	input := []string{"Hello", "HELLO", "cat", "cat", "dog"}
	expected := []string{"Hello", "HELLO", "cat", "dog"}
*/
func TestRemoveDuplicatesString(t *testing.T) {
	input := []string{"Hello", "HELLO", "cat", "cat", "dog"}
	expected := []string{"Hello", "HELLO", "cat", "dog"}
	actual := RemoveDuplicateString(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatError(fmt.Sprintf("%v", input), fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual)))
	}
}

/*
	TestRemoveAllInstancesOfInts_3
	input := []int{}
	input2 := []int{}
	expected: Length of 0
*/
func TestRemoveAllInstancesOfInts_3(t *testing.T) {
	input := []int{}
	input2 := []int{}
	actual := RemoveAllInstancesOfInts(input, input2...)
	expectedStr := "Length of 0"
	if len(actual) != 0 {
		t.Error(formatErrorTwoInputs(input, input2, expectedStr, actual))
	}
}

/*
	TestRemoveAllInstancesOfFloat64s
	input := []int{1, 2, 3, 2, 2, 4, 2, 6, 12}
	input2 := []int{2, 1}
	expected := []int{3, 4, 6, 12}
*/
func TestRemoveAllInstancesOfFloat64s_1(t *testing.T) {
	input := []float64{2, 2, 2}
	input2 := []float64{2, 2}
	expected := []float64{}
	actual := RemoveAllInstancesOfFloat64s(input, input2...)

	if len(actual) != 0 {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

/*
	TestRemoveAllInstancesOfFloat64s_2
	input := []float64{1.123, 2.53, 3.0000000000001, 3, 2, 4, 2, 6, 12}
	input2 := []float64{3}
	expected := []float64{1.123, 2.53, 3.0000000000001, 2, 4, 2, 6, 12}
*/
func TestRemoveAllInstancesOfFloat64s_2(t *testing.T) {
	input := []float64{1.123, 2.53, 3.0000000000001, 3, 2, 4, 2, 6, 12}
	input2 := []float64{3}
	expected := []float64{1.123, 2.53, 3.0000000000001, 2, 4, 2, 6, 12}
	actual := RemoveAllInstancesOfFloat64s(input, input2...)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

/*
	TestRemoveDuplicatesString_2
	input := []string{"", "", "Hey65", "Hey65"}
	expected := []string{"", "Hey65"}
*/
func TestRemoveDuplicatesString_2(t *testing.T) {
	input := []string{"", "", "Hey65", "Hey65"}
	expected := []string{"", "Hey65"}
	actual := RemoveDuplicateString(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatError(fmt.Sprintf("%v", input), fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual)))
	}
}

/*
	TestRemoveAllInstancesOfFloat64s_3
	input := []float64{1.123, 2.53, 3.0000000000001, 3, 2, 4, 2, 6, 12}
	input2 := []float64{3}
	expected := []float64{1.123, 2.53, 3.0000000000000, 2, 4, 2, 6, 12}
*/
func TestRemoveAllInstancesOfFloat64s_3(t *testing.T) {
	input := []float64{1.123, 2.53, 3.0000000000000, 3, 2, 4, 2, 6, 12}
	input2 := []float64{3}
	expected := []float64{1.123, 2.53, 2, 4, 2, 6, 12}
	actual := RemoveAllInstancesOfFloat64s(input, input2...)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

/*
	TestContainsString_1
	input := []string{"one", "too", "tree"}
	input2 := "too"
	expected := true
*/
func TestContainsString_1(t *testing.T) {
	input := []string{"one", "too", "tree"}
	input2 := "too"
	expected := true
	actual := ContainsString(input, input2)

	if actual != expected {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

/*
	TestContainsString_2
	input := []string{"one", "too", "tree"}
	input2 := "two"
	expected := false
*/
func TestContainsString_2(t *testing.T) {
	input := []string{"one", "too", "tree"}
	input2 := "two"
	expected := false
	actual := ContainsString(input, input2)

	if actual != expected {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

/*
	TestContainsInt_1
	input := []int{1, 2, 6}
	input2 := 1
	expected := true
*/
func TestContainsInt_1(t *testing.T) {
	input := []int{1, 2, 6}
	input2 := 1
	expected := true
	actual := ContainsInt(input, input2)

	if actual != expected {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

/*
	TestContainsInt_2
	input := []int{1, 2, 6}
	input2 := 3
	expected := false
*/
func TestContainsInt_2(t *testing.T) {
	input := []int{1, 2, 6}
	input2 := 3
	expected := false
	actual := ContainsInt(input, input2)

	if actual != expected {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

/*
	TestContainsFloat64_1
	input := []float64{1.00000, 2, 6}
	input2 := 1.0
	expected := true
*/
func TestContainsFloat64_1(t *testing.T) {
	input := []float64{1.00000, 2, 6}
	input2 := 1.0
	expected := true
	actual := ContainsFloat64(input, input2)

	if actual != expected {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

/*
	TestContainsFloat64_2
	input := []int{1.00000000001, 2, 6}
	input2 := 1.0
	expected := false
*/
func TestContainsFloat64_2(t *testing.T) {
	input := []float64{1.00000000001, 2, 6}
	input2 := 1.0
	expected := false
	actual := ContainsFloat64(input, input2)

	if actual != expected {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

/*
	TestRemoveDuplicatesString_3
	input := []string{"Hey"}
	expected := []string{"Hey"}
*/
func TestRemoveDuplicatesString_3(t *testing.T) {
	input := []string{"Hey"}
	expected := []string{"Hey"}
	actual := RemoveDuplicateString(input)

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
func formatErrorTwoInputs(input interface{}, input2 interface{}, expected interface{}, actual interface{}) string {
	return fmt.Sprintf("\nInput:    %v\nInput2:   %v\nExpected: %v\nActual:   %v\n", input, input2, expected, actual)
}
