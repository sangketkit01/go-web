package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := divide(100.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func divide(x, y float64) (float64, error) {
	var result float64

	if y == 0 {
		return result, errors.New("division by zero")
	}

	result = x / y
	return result, nil
}
