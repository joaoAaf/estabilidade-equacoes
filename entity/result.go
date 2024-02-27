package entity

type Result struct {
	Calcs   []Calc `json:"calcs"`
	Estable bool   `json:"estable"`
}
