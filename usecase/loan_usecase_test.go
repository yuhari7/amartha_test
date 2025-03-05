package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuhari7/amartha_test/domain"
	"github.com/yuhari7/amartha_test/repository"
)

func TestGetOutstanding(t *testing.T) {
	loanRepo := repository.NewLoanRepository()
	loan := &domain.Loan{
		ID:          100,
		Amount:      5000000,
		Outstanding: 5500000,
	}

	loanRepo.UpdateLoan(loan)
	loanUsecase := NewLoanUsecase(loanRepo)

	outstanding, err := loanUsecase.GetOutstanding(100)
	assert.NoError(t, err)
	assert.Equal(t, 5500000.0, outstanding)
}
