package calculator

import (
	"errors"

	"github.com/g-utils/overflow"
)

func Add(a, b int) (int, error) {
	sum, ok := overflow.Add(a, b)
	if !ok {
		return 0, errors.New("overflow")
	}
	return sum, nil
}

func Subtract(a, b int) (int, error) {
	diff, ok := overflow.Sub(a, b)
	if !ok {
		return 0, errors.New("overflow")
	}
	return diff, nil
}

func Multiply(a, b int) (int, error) {
	prod, ok := overflow.Mul(a, b)
	if !ok {
		return 0, errors.New("overflow")
	}
	return prod, nil
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
