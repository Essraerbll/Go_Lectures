package main

import (
	"errors"
	"fmt"
)

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0.0, errors.New("division by zero")
	}
	result := dividend / divisor
	return result, nil
}

func main2() {
	result, err := divide(10, 5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}
