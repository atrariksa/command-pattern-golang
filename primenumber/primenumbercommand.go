package primenumber

import (
	"log"

	"github.com/atrariksa/command-pattern-golang/input"
)

type ICommand interface {
	Execute(input input.PrimeNumberInput) PrimeNumberOutput
}

func Execute(inp input.PrimeNumberInput) PrimeNumberOutput {
	var res PrimeNumberOutput
	r, err := input.PrimeNumber(inp)
	if err != nil {
		log.Fatal(err)
	}
	res.Value = r
	return res
}
