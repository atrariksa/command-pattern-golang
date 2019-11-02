package sum

import (
	"github.com/atrariksa/command-pattern-golang/input"
)

type ICommand interface {
	Execute(input input.SumInput) SumOutput
}

func Execute(inp input.SumInput) SumOutput {
	var res SumOutput
	res.Value = input.Sum(inp)
	return res
}
