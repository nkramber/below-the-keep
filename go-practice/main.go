package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
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
