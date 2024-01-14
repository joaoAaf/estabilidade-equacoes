package data

import "fmt"

func ReceiveData() (int, []float64) {
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
