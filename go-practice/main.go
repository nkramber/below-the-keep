package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("cannot divide by zero")

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide %d / %d: %w", a, b, ErrDivideByZero)
	}
	return a / b, nil
}

func main() {
	numsToDivide := [][2]int{
		{4, 2},
		{9, 3},
		{5, 0},
	}

	for _, nums := range numsToDivide {
		result, err := divide(nums[0], nums[1])
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result)
	}
}
