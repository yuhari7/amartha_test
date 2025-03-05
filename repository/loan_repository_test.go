package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuhari7/amartha_test/domain"
)

func TestLoanRepository_GetLoanByID(t *testing.T) {
	repo := NewLoanRepository()

	loan := &domain.Loan{ID: 100, Outstanding: 5500000}
	repo.UpdateLoan(loan)

	fetchedLoan, err := repo.GetLoanByID(100)
	assert.NoError(t, err)
	assert.Equal(t, loan, fetchedLoan)
}

func TestLoanRepository_GetLoanByID_NotFound(t *testing.T) {
	repo := NewLoanRepository()

	_, err := repo.GetLoanByID(200) // Loan ID 200 doesn't exist
	assert.Error(t, err)
	assert.Equal(t, "loan not found", err.Error())
}

func TestLoanRepository_UpdateLoan(t *testing.T) {
	repo := NewLoanRepository()

	loan := &domain.Loan{ID: 100, Outstanding: 5500000}
	err := repo.UpdateLoan(loan)
	assert.NoError(t, err)

	updatedLoan, err := repo.GetLoanByID(100)
	assert.NoError(t, err)
	assert.Equal(t, 5500000.0, updatedLoan.Outstanding)
}
