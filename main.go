package main

import (
	"estabilidade-equacoes/calc"
	"estabilidade-equacoes/cli"
	"estabilidade-equacoes/data"
)

func main() {
	validate := calc.ValidateEquation(data.ReceiveData())
	cli.ResultCalc(validate)
}
