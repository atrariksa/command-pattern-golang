package multiply

import (
	"github.com/atrariksa/command-pattern-golang/input"
)

type ICommand interface {
	Execute(input input.MultiplyInput) MultiplyOutput
}

func Execute(inp input.MultiplyInput) MultiplyOutput {
	var res MultiplyOutput
	res.Value = input.Multiply(inp)
	return res
}
