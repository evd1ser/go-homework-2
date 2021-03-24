package store

import (
	"github.com/evd1ser/go-homework-2/internal/app/models"
)

type EquationRepository struct {
	db *models.Equation
}

//Create user in database
func (er *EquationRepository) Set(u *models.Equation) *models.Equation {
	er.db = u
	return u
}

func (er *EquationRepository) GetWithCalc() *models.Equation {
	er.db.Calculate()
	return er.db
}
