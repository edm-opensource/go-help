package integers

import (
	"fmt"
	"testing"
)

func formatError(input interface{}, expected interface{}, actual interface{}) string {
	//return "Input: " + input.(string) + "\nExpected: " + expected.(string) + "\nActual: " + actual.(string)
	return fmt.Sprintf("\nInput:    %v\nExpected: %v\nActual:   %v", input, expected, actual)
}

type testdataModPow struct {
	x      int
	y      int
	n      int
	result int
}

func TestModPow(t *testing.T) {
	tests := []testdataModPow{
		{875898931, 93086444, 10494, 9799},
		{1650865812, 597236599, 33252, 22188},
		{1091473074, 2038076398, 4808, 3760},
		{1564115234, 514985747, 35489, 31261},
	}

	for _, data := range tests {
		result := ModPow(data.x, data.y, data.n)
		if result != data.result {
			t.Error(formatError(data, data, result))
		}
	}

	result := ModPow(875898931, 93086444, 10494)
	if result != 9799 {
		t.Error("\nfel")
	}
}
