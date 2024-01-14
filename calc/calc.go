package calc

import "fmt"

func invertMatrix(matrix []float64) []float64 {
	var matrixR []float64
	for i := len(matrix) - 1; i >= 0; i-- {
		if matrix[i] == 0 && i == len(matrix)-1 {
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

func ValidateEquation(exponent int, indexes []float64) bool {
	var matrix1 []float64
	var matrix2 []float64
	var j float64
	var estable bool
	for k := 0; k <= exponent; k++ {
		if k == 0 {
			matrix1 = indexes
			matrix2 = invertMatrix(matrix1)
		} else if k == exponent {
			matrix1 = calcMatrix(matrix1, matrix2, j)
			if matrix1[len(matrix1)-1] == 0 {
				matrix1 = matrix1[:len(matrix1)-1]
			}
			matrix2 = invertMatrix(matrix1)
		} else {
			matrix1 = calcMatrix(matrix1, matrix2, j)
			matrix2 = invertMatrix(matrix1)
		}
		if matrix1[0] != 0 && k != exponent {
			j = matrix2[0] / matrix1[0]
		}
		if j < 0 {
			j *= -1
		}
		if j < 1 {
			estable = true
		} else {
			estable = false
		}
		fmt.Printf("===========================================\n")
		fmt.Printf("K = %d           %.4g               \n", k, matrix1)
		fmt.Printf("                %.4g                \n", matrix2)
		fmt.Printf("-------------------------------------------\n")
		fmt.Printf("J = %.4g                            \n", j)
		if k > 0 && matrix1[len(matrix1)-1] == 0 {
			matrix1 = matrix1[:len(matrix1)-1]
		}
	}
	fmt.Printf("===========================================\n")
	if len(matrix1) > 1 || matrix1[0] == 0 {
		estable = false
	}
	return estable
}
