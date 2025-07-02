package main

import (
	"errors"
	"fmt"
)

func main() {
	x, y := 56, 78
	op := func(x, y int) int {
		return x + y
	}
	add := genMyCalculator(op)
	result, err := add(x, y)
	fmt.Printf("The result: %d (error: %v)\n", result, err)
}

type operate1 func(x, y int) int

type myCal func(x, y int) (int, error)

func genMyCalculator(op operate1) myCal {
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New("abc")
		}
		return op(x, y), nil
	}
}
