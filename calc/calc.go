package calc

import "errors"

func Sub(num1, num2 int) int {
	return num1 - num2
}

func Divide(num, den int) (int, error) {
	if den == 0 {
		return 0, errors.New("o denominador não pode ser 0")
	}
	return num / den, nil
}
