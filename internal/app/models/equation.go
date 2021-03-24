package models

import (
	"math"
)

type Equation struct {
	A      float64 `json:"A"`
	B      float64 `json:"B"`
	C      float64 `json:"C"`
	Nroots int     `json:"Nroots"`
}

func (e *Equation) Calculate() {

	D := math.Pow(e.B, 2.0) - (4 * e.A * e.C)

	if (D == 0 && e.B != 0) || (e.A == 0 && e.B != 0) {
		e.Nroots = 1
		return
	}

	if D > 0 {
		e.Nroots = 2
		return
	}

	e.Nroots = 0
}
