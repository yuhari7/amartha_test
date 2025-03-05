package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuhari7/amartha_test/internal/domain/entities"
	"github.com/yuhari7/amartha_test/internal/domain/usecases"
	"github.com/yuhari7/amartha_test/internal/infrastructure/repository"
)

func initializeLoan(loanRepo *repository.InMemoryLoanRepository, id int, payments []entities.Payment) {
	loanRepo.Save(entities.Loan{
		ID:            id,
		Amount:        5000000,
		InterestRate:  10,
		TermWeeks:     50,
		WeeklyPayment: 110000,
		Payments:      payments,
	})
}

func TestLoanService_GetOutstanding(t *testing.T) {
	loanRepo := repository.NewInMemoryLoanRepository()
	loanRepo.Loans[100] = entities.Loan{
		ID: 100, Amount: 5000000, InterestRate: 10,
		Payments: []entities.Payment{{Week: 1, Amount: 110000}, {Week: 2, Amount: 110000}},
	}

	service := usecases.LoanService{LoanRepository: loanRepo}
	outstanding, err := service.GetOutstanding(100)

	assert.NoError(t, err)
	assert.Equal(t, 5280000.0, outstanding)
}

func TestLoanService_IsDelinquent(t *testing.T) {
	loanRepo := repository.NewInMemoryLoanRepository()
	loanRepo.Loans[100] = entities.Loan{
		ID: 100, Payments: []entities.Payment{{Week: 1, Amount: 0}, {Week: 2, Amount: 0}},
	}

	service := usecases.LoanService{LoanRepository: loanRepo}
	delinquent, err := service.IsDelinquent(100)

	assert.NoError(t, err)
	assert.True(t, delinquent)
}

func TestLoanService_MakePayment_ValidAmount(t *testing.T) {
	loanRepo := repository.NewInMemoryLoanRepository()
	initializeLoan(loanRepo, 100, []entities.Payment{{Week: 1, Amount: 110000}})

	service := usecases.LoanService{LoanRepository: loanRepo}

	// Valid amount
	message, err := service.MakePayment(100, 110000)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(loanRepo.Loans[100].Payments))
	assert.Equal(t, "You paid for week 2", message)

	// Payment after loan term completion
	initializeLoan(loanRepo, 101, []entities.Payment{})

	for i := 0; i < 50; i++ {
		message, err = service.MakePayment(101, 110000)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assert.Equal(t, fmt.Sprintf("You paid for week %d", i+1), message)
	}

	message, err = service.MakePayment(101, 110000)
	assert.Error(t, err)
	assert.Equal(t, "your loan is 0: loan term completed, no more payments allowed", err.Error())
	assert.Equal(t, "", message)
}

func TestLoanService_MakePayment_InvalidAmount(t *testing.T) {
	loanRepo := repository.NewInMemoryLoanRepository()
	initializeLoan(loanRepo, 100, []entities.Payment{{Week: 1, Amount: 110000}})

	service := usecases.LoanService{LoanRepository: loanRepo}

	// Invalid amount
	message, err := service.MakePayment(100, 100000)
	assert.Error(t, err)
	assert.Equal(t, "invalid payment amount, should be exactly the weekly payment amount", err.Error())
	assert.Equal(t, "", message)

	// Non-positive amount
	message, err = service.MakePayment(100, -110000)
	assert.Error(t, err)
	assert.Equal(t, "payment amount must be positive", err.Error())
	assert.Equal(t, "", message)

	// Additional invalid large amount check
	initializeLoan(loanRepo, 101, []entities.Payment{})

	message, err = service.MakePayment(101, 110001)
	assert.Error(t, err)
	assert.Equal(t, "invalid payment amount, should be exactly the weekly payment amount", err.Error())
	assert.Equal(t, "", message)
}
