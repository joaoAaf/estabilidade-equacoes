package main

import (
	"estabilidade-equacoes/calc"
	"estabilidade-equacoes/data"
)

func main() {
	calc.ValidateEquation(data.ReceiveData())
}
