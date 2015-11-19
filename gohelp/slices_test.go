package gohelp

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseStrings_1(t *testing.T) {
	input := []string{"one", "two", "three"}
	expected := []string{"three", "two", "one"}
	actual := ReverseStrings(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Error(formatError(input, expected, actual))
	}
}

func TestReverseStrings_2(t *testing.T) {
	input := []string{"one"}
	expected := []string{"one"}
	actual := ReverseStrings(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Error(formatError(input, expected, actual))
	}
}

func TestReverseStrings_3(t *testing.T) {
	var input []string
	var expected []string
	actual := ReverseStrings(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Error(formatError(input, expected, actual))
	}
}

func TestReverseInts_1(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{3, 2, 1}
	actual := ReverseInts(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Error(formatError(input, expected, actual))
	}
}

func TestReverseInts_2(t *testing.T) {
	input := []int{1}
	expected := []int{1}
	actual := ReverseInts(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Error(formatError(input, expected, actual))
	}
}

func TestReverseInts_3(t *testing.T) {
	var input []int
	var expected []int
	actual := ReverseInts(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Error(formatError(input, expected, actual))
	}
}

func TestReverseFloat64s_1(t *testing.T) {
	input := []float64{1.123, 2.3, 3.12312}
	expected := []float64{3.12312, 2.3, 1.123}
	actual := ReverseFloat64s(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Error(formatError(input, expected, actual))
	}
}

func TestReverseFloat64s_2(t *testing.T) {
	input := []float64{1.451}
	expected := []float64{1.451}
	actual := ReverseFloat64s(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Error(formatError(input, expected, actual))
	}
}

func TestReverseFloat64s_3(t *testing.T) {
	var input []float64
	var expected []float64
	actual := ReverseFloat64s(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Error(formatError(input, expected, actual))
	}
}

func TestRemoveDuplicates(t *testing.T) {
	input := []int{1, 1}
	expected := []int{1}
	actual := RemoveDuplicateInts(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatError(fmt.Sprintf("%v", input), fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual)))
	}
}

func TestRemoveDuplicates_2(t *testing.T) {
	input := []int{9, 1, 1, 1, 100, 100, 9, 9}
	expected := []int{9, 1, 100}
	actual := RemoveDuplicateInts(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatError(fmt.Sprintf("%v", input), fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual)))
	}
}

func TestRemoveDuplicates_3(t *testing.T) {
	input := []int{}
	expected := []int{}
	actual := RemoveDuplicateInts(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatError(fmt.Sprintf("%v", input), fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual)))
	}
}

func TestRemoveAllInstancesOfStrings_1(t *testing.T) {
	input := []string{"one", "one", "two", "oone"}
	input2 := []string{"one"}
	expected := []string{"two", "oone"}
	actual := RemoveAllInstancesOfStrings(input, input2...)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

func TestRemoveDuplicatesFloat(t *testing.T) {
	input := []float64{1.0, 1.0}
	expected := []float64{1.0}
	actual := RemoveDuplicateFloat64(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatError(fmt.Sprintf("%v", input), fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual)))

	}
}

func TestRemoveAllInstancesOfStrings_2(t *testing.T) {
	input := []string{"one", "\none", "two", "oone"}
	input2 := []string{"one"}
	expected := []string{"\none", "two", "oone"}
	actual := RemoveAllInstancesOfStrings(input, input2...)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

func TestRemoveDuplicatesFloat_2(t *testing.T) {
	input := []float64{9.0, 1.1, 1.0, 1.0, 100, 100, 9.0, 9.0}
	expected := []float64{9.0, 1.1, 1.0, 100}
	actual := RemoveDuplicateFloat64(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatError(fmt.Sprintf("%v", input), fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual)))
	}
}

func TestRemoveAllInstancesOfStrings_3(t *testing.T) {
	input := []string{"one", "one", "two", "oone"}
	input2 := []string{""}
	expected := []string{"one", "one", "two", "oone"}
	actual := RemoveAllInstancesOfStrings(input, input2...)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

func TestRemoveDuplicatesFloat_3(t *testing.T) {
	input := []float64{}
	expected := []float64{}
	actual := RemoveDuplicateFloat64(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatError(fmt.Sprintf("%v", input), fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual)))
	}
}

func TestRemoveAllInstancesOfInts_1(t *testing.T) {
	input := []int{1, 2, 3, 2, 2, 4, 2, 6, 12}
	input2 := []int{2, 1}
	expected := []int{3, 4, 6, 12}
	actual := RemoveAllInstancesOfInts(input, input2...)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

func TestRemoveDuplicatesFloat_4(t *testing.T) {
	input := []float64{-1, -100, -34.55, -100}
	expected := []float64{-1, -100, -34.55}
	actual := RemoveDuplicateFloat64(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatError(fmt.Sprintf("%v", input), fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual)))
	}
}

func TestRemoveAllInstancesOfInts_2(t *testing.T) {
	input := []int{1, 2, 3, 2, 2, 4, 2, 6, 12}
	input2 := []int{16}
	expected := []int{1, 2, 3, 2, 2, 4, 2, 6, 12}
	actual := RemoveAllInstancesOfInts(input, input2...)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

func TestRemoveDuplicatesString(t *testing.T) {
	input := []string{"Hello", "HELLO", "cat", "cat", "dog"}
	expected := []string{"Hello", "HELLO", "cat", "dog"}
	actual := RemoveDuplicateString(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatError(fmt.Sprintf("%v", input), fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual)))
	}
}

func TestRemoveAllInstancesOfInts_3(t *testing.T) {
	input := []int{}
	input2 := []int{}
	actual := RemoveAllInstancesOfInts(input, input2...)
	expectedStr := "Length of 0"
	if len(actual) != 0 {
		t.Error(formatErrorTwoInputs(input, input2, expectedStr, actual))
	}
}

func TestRemoveAllInstancesOfFloat64s_1(t *testing.T) {
	input := []float64{2, 2, 2}
	input2 := []float64{2, 2}
	expected := []float64{}
	actual := RemoveAllInstancesOfFloat64s(input, input2...)

	if len(actual) != 0 {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

func TestRemoveAllInstancesOfFloat64s_2(t *testing.T) {
	input := []float64{1.123, 2.53, 3.0000000000001, 3, 2, 4, 2, 6, 12}
	input2 := []float64{3}
	expected := []float64{1.123, 2.53, 3.0000000000001, 2, 4, 2, 6, 12}
	actual := RemoveAllInstancesOfFloat64s(input, input2...)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

func TestRemoveDuplicatesString_2(t *testing.T) {
	input := []string{"", "", "Hey65", "Hey65"}
	expected := []string{"", "Hey65"}
	actual := RemoveDuplicateString(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatError(fmt.Sprintf("%v", input), fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual)))
	}
}

func TestRemoveAllInstancesOfFloat64s_3(t *testing.T) {
	input := []float64{1.123, 2.53, 3.0000000000000, 3, 2, 4, 2, 6, 12}
	input2 := []float64{3}
	expected := []float64{1.123, 2.53, 2, 4, 2, 6, 12}
	actual := RemoveAllInstancesOfFloat64s(input, input2...)

	if !reflect.DeepEqual(expected, actual) {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

func TestContainsString_1(t *testing.T) {
	input := []string{"one", "too", "tree"}
	input2 := "too"
	expected := true
	actual := ContainsString(input, input2)

	if actual != expected {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

func TestContainsString_2(t *testing.T) {
	input := []string{"one", "too", "tree"}
	input2 := "two"
	expected := false
	actual := ContainsString(input, input2)

	if actual != expected {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

func TestContainsInt_1(t *testing.T) {
	input := []int{1, 2, 6}
	input2 := 1
	expected := true
	actual := ContainsInt(input, input2)

	if actual != expected {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

func TestContainsInt_2(t *testing.T) {
	input := []int{1, 2, 6}
	input2 := 3
	expected := false
	actual := ContainsInt(input, input2)

	if actual != expected {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

func TestContainsFloat64_1(t *testing.T) {
	input := []float64{1.00000, 2, 6}
	input2 := 1.0
	expected := true
	actual := ContainsFloat64(input, input2)

	if actual != expected {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

func TestContainsFloat64_2(t *testing.T) {
	input := []float64{1.00000000001, 2, 6}
	input2 := 1.0
	expected := false
	actual := ContainsFloat64(input, input2)

	if actual != expected {
		t.Error(formatErrorTwoInputs(input, input2, expected, actual))
	}
}

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
