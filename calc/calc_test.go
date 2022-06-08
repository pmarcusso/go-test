package calc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSub(t *testing.T) {
	num1 := 10
	num2 := 5
	expectedResult := 5
	result := Sub(num1, num2)

	if expectedResult != result {
		t.Errorf("Expected result: %d - Result: %d", expectedResult, result)
	}
}

func TestDivide(t *testing.T) {
	num := 10
	den := 2
	expectedResult := 5

	result, _ := Divide(num, den)
	assert.Equal(t, expectedResult, result)

}

func TestDivideError(t *testing.T) {
	num := 10
	den := 0
	expectedMsg := "o denominador não pode ser 0"

	_, err := Divide(num, den)

	if err != nil {
		assert.Equal(t, expectedMsg, err.Error())
	}
}
