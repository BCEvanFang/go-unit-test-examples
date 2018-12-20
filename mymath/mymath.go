package mymath

import "errors"

func Add(x, y int) int {
	return x + y
}

func Multiply(x, y int) int {
	return x * y
}

func Devide(x, y float32) (float32, error) {

	// Check!
	if y == 0 {
		return 0, errors.New("Cannot devide by 0")
	}

	return x / y, nil
}
