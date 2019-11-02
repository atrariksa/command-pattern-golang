package input

import errorhandler "github.com/atrariksa/command-pattern-golang/errors"

type Process interface {
	Sum(input SumInput) int
	Multiply(input MultiplyInput) int
	PrimeNumber(input PrimeNumberInput) ([]int, errorhandler.Err)
	Fibonacci(input FibonacciInput) ([]int, errorhandler.Err)
}
