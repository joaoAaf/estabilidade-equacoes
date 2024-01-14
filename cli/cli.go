package cli

import (
	"estabilidade-equacoes/calc"
	"fmt"
)

func receiveData() (int, []float64) {
	var exponent int
	var indexes []float64
	fmt.Scanf("%d", &exponent)
	for i := 0; i <= exponent; i++ {
		var index float64
		fmt.Scanf("%g", &index)
		indexes = append(indexes, index)
	}
	return exponent, indexes
}

func Menu() {
	var opt string
	for opt != "0" {
		fmt.Println()
		fmt.Println("===========================================")
		fmt.Println("====CALCULO DE ESTABILIDADE DE EQUAÇÕES====")
		fmt.Println("===========================================")
		fmt.Println("Digite opção desejada:")
		fmt.Println("1 - Iniciar Calculo")
		fmt.Println("0 - Sair")
		fmt.Scanln(&opt)
		switch opt {
		case "1":
			fmt.Println()
			fmt.Println("Digite o grau da equação e os indices em sequencia, separados por espaços, conforme o exemplo abaixo.")
			fmt.Println("Exemplo: 8z³+z²+z+4 --> 3 8 1 1 4")
			result := calc.ValidateEquation(receiveData())
			calc.ResultCalc(result)
		case "0":
		default:
			fmt.Println()
			fmt.Println("Opção Inválida!")
		}
	}
}
