package calc

import "estabilidade-equacoes/entity"

func invertMatrix(matrix []float64, k int) []float64 {
	var matrixR []float64
	for i := len(matrix) - 1; i >= 0; i-- {
		if matrix[i] == 0 && i == len(matrix)-1 && k != 0 {
		} else if len(matrix) == 1 {
			matrixR = matrixR[:]
		} else {
			matrixR = append(matrixR, matrix[i])
		}
	}
	return matrixR
}

func calcMatrix(matrix1 []float64, matrix2 []float64, j float64) []float64 {
	var matrix []float64
	for i := 0; i < len(matrix1); i++ {
		element := matrix1[i] - matrix2[i]*j
		matrix = append(matrix, element)
	}
	return matrix
}

func ValidateEquation(input entity.Input) entity.Result {
	var calc entity.Calc
	var result entity.Result
	var j float64
	var count int
	for calc.K = 0; calc.K <= len(input.Indexes)-1; calc.K++ {
		if calc.K == 0 {
			calc.Matrix1 = input.Indexes
			calc.Matrix2 = invertMatrix(calc.Matrix1, calc.K)
		} else if calc.K == len(input.Indexes)-1 {
			calc.Matrix1 = calcMatrix(calc.Matrix1, calc.Matrix2, j)
			if calc.Matrix1[len(calc.Matrix1)-1] == 0 {
				calc.Matrix1 = calc.Matrix1[:len(calc.Matrix1)-1]
			}
			calc.Matrix2 = invertMatrix(calc.Matrix1, calc.K)
		} else {
			calc.Matrix1 = calcMatrix(calc.Matrix1, calc.Matrix2, j)
			calc.Matrix2 = invertMatrix(calc.Matrix1, calc.K)
		}
		if calc.Matrix1[0] != 0 && calc.K != len(input.Indexes)-1 {
			j = calc.Matrix2[0] / calc.Matrix1[0]
			calc.JModule = j
		}
		if calc.JModule < 0 {
			calc.JModule *= -1
		}
		if calc.JModule < 1 && calc.JModule > 0 {
			result.Estable = true
		} else {
			count += 1
		}
		result.Calcs = append(result.Calcs, calc)
		if calc.K > 0 && calc.Matrix1[len(calc.Matrix1)-1] == 0 && calc.K < len(calc.Matrix1) {
			calc.Matrix1 = calc.Matrix1[:len(calc.Matrix1)-1]
		}
	}
	if len(calc.Matrix1) > 1 || calc.Matrix1[0] == 0 || count != 0 {
		result.Estable = false
	}
	return result
}
