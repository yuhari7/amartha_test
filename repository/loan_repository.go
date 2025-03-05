package repository

import (
	"errors"

	"github.com/yuhari7/amartha_test/domain"
)

type LoanRepository struct {
	loans map[int]*domain.Loan
}

func NewLoanRepository() *LoanRepository {
	return &LoanRepository{
		loans: make(map[int]*domain.Loan),
	}
}

func (r *LoanRepository) GetLoanByID(id int) (*domain.Loan, error) {
	loan, exists := r.loans[id]
	if !exists {
		return nil, errors.New("loan not found")
	}

	return loan, nil
}

func (r *LoanRepository) UpdateLoan(loan *domain.Loan) error {
	r.loans[loan.ID] = loan
	return nil
}
