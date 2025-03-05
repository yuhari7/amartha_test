package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuhari7/amartha_test/domain"
)

func TestMakePayment(t *testing.T) {
	service := NewPaymentService()
	loan := &domain.Loan{
		ID:          100,
		Amount:      5000000,
		Outstanding: 5500000,
		Schedule:    []domain.Payment{{Week: 1, Amount: 110000}, {Week: 2, Amount: 110000}},
	}

	err := service.MakePayment(loan, 110000)
	assert.NoError(t, err)
	assert.Equal(t, 5390000.0, loan.Outstanding)
	assert.Equal(t, 1, len(loan.Schedule))
}
