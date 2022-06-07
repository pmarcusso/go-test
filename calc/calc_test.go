package calc

import "testing"

func TestSub(t *testing.T) {
	num1 := 10
	num2 := 5
	expectedResult := 5
	result := Sub(num1, num2)

	if expectedResult != result {
		t.Errorf("Expected result: %d - Result: %d", expectedResult, result)
	}
}
