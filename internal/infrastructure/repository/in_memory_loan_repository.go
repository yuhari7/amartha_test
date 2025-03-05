package repository

import (
	"errors"

	"github.com/yuhari7/amartha_test/internal/domain/entities"
)

type InMemoryLoanRepository struct {
	Loans map[int]entities.Loan
}

func NewInMemoryLoanRepository() *InMemoryLoanRepository {
	return &InMemoryLoanRepository{
		Loans: make(map[int]entities.Loan),
	}
}

func (r *InMemoryLoanRepository) GetLoanByID(id int) (entities.Loan, error) {
	loan, exists := r.Loans[id]
	if !exists {
		return entities.Loan{}, errors.New("loan not found")
	}
	return loan, nil
}

func (r *InMemoryLoanRepository) Save(loan entities.Loan) error {
	r.Loans[loan.ID] = loan
	return nil
}
