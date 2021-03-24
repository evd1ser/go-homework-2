package store

import (
	_ "github.com/lib/pq"
)

//Instance of store
type Store struct {
	equationRepository *EquationRepository
}

// Constructor for store
func New() *Store {
	return &Store{}
}

//Public for UserRepo
func (s *Store) Equation() *EquationRepository {
	if s.equationRepository != nil {
		return s.equationRepository
	}
	s.equationRepository = &EquationRepository{}
	return s.equationRepository
}
