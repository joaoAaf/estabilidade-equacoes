package calc

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

func ValidateEquation(exponent int, indexes []float64) bool {
	var matrix1 []float64
	var matrix2 []float64
	var j float64
	var jModule float64
	var estable bool
	var count int
	for k := 0; k <= exponent; k++ {
		if k == 0 {
			matrix1 = indexes
			matrix2 = invertMatrix(matrix1, k)
		} else if k == exponent {
			matrix1 = calcMatrix(matrix1, matrix2, j)
			if matrix1[len(matrix1)-1] == 0 {
				matrix1 = matrix1[:len(matrix1)-1]
			}
			matrix2 = invertMatrix(matrix1, k)
		} else {
			matrix1 = calcMatrix(matrix1, matrix2, j)
			matrix2 = invertMatrix(matrix1, k)
		}
		if matrix1[0] != 0 && k != exponent {
			j = matrix2[0] / matrix1[0]
			jModule = j
		}
		if jModule < 0 {
			jModule *= -1
		}
		if jModule < 1 && jModule > 0 {
			estable = true
		} else {
			count += 1
		}
		tableResponse(k, exponent, matrix1, matrix2, jModule)
		if k > 0 && matrix1[len(matrix1)-1] == 0 {
			matrix1 = matrix1[:len(matrix1)-1]
		}
	}
	if len(matrix1) > 1 || matrix1[0] == 0 || count != 0 {
		estable = false
	}
	return estable
}
