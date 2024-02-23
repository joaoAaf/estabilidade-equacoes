package calc

import "fmt"

func tableResponse(k, exponent int, matrix1, matrix2 []float64, j float64) {
	if k < exponent {
		fmt.Printf("===========================================\n")
		fmt.Printf("K = %d           %.4g               \n", k, matrix1)
		fmt.Printf("                %.4g                \n", matrix2)
		fmt.Printf("-------------------------------------------\n")
		fmt.Printf("|J| = %.4g                            \n", j)
	} else {
		fmt.Printf("===========================================\n")
		fmt.Printf("K = %d           %.4g               \n", k, matrix1)
		fmt.Printf("===========================================\n")
	}
}

func ResultCalc(validate bool) {
	if validate {
		fmt.Println("A equação é estavel.")
	} else {
		fmt.Println("A equação é instavel.")
	}
}
