package input

import (
	"errors"

	"github.com/atrariksa/command-pattern-golang/constants"
	errorhandler "github.com/atrariksa/command-pattern-golang/errors"
)

var output []int
var inputParam FibonacciInput

func Fibonacci(input FibonacciInput) ([]int, errorhandler.Err) {
	inputParam = input
	output := []int{}
	if input.X == 0 {
		return output, errors.New(constants.FibonacciInputError)
	}
	if input.IsRecursive {
		return recursiveFibonacci(output), nil
	} else {
		return loopingFibonacci(), nil
	}
}
func recursiveFibonacci(input []int) []int {
	size := len(input)
	if size == inputParam.X {
		return input
	}
	if size == 0 {
		input = append(input, 0)
		return recursiveFibonacci(input)
	} else if size == 1 {
		input = append(input, 1)
		return recursiveFibonacci(input)
	} else {
		x := input[size-2]
		y := input[size-1]
		input = append(input, x+y)
		return recursiveFibonacci(input)
	}
}
func loopingFibonacci() []int {
	if inputParam.X == 1 {
		output = append(output, 0)
		return output
	} else if inputParam.X == 2 {
		output = append(output, 0)
		output = append(output, 1)
		return output
	} else {
		output = append(output, 0)
		output = append(output, 1)
		for {
			if len(output) == inputParam.X {
				break
			}
			size := len(output)
			x := output[size-2]
			y := output[size-1]
			output = append(output, x+y)
		}
		return output
	}
}
