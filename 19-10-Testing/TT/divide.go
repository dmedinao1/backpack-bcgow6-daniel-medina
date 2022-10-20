package main

import "fmt"

var ErrDivisionBy0 = fmt.Errorf("couldn't divde by zero")

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionBy0
	}

	return a / b, nil
}
