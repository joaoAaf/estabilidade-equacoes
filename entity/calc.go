package entity

type Calc struct {
	K       int       `json:"k"`
	Matrix1 []float64 `json:"matriz1"`
	Matrix2 []float64 `json:"matriz2"`
	JModule float64   `json:"jmodule"`
}
