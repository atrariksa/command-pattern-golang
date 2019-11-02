package fibonacci

import (
	"log"

	"github.com/atrariksa/command-pattern-golang/input"
)

type ICommand interface {
	Execute(input input.FibonacciInput) FibonacciOutput
}

func Execute(inp input.FibonacciInput) FibonacciOutput {
	var res FibonacciOutput
	r, err := input.Fibonacci(inp)
	if err != nil {
		log.Fatal(err)
	}
	res.Value = r
	return res
}
