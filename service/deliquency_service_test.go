package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuhari7/amartha_test/domain"
)

func TestCheckDelinquency(t *testing.T) {
	service := NewDelinquencyService()
	loan := &domain.Loan{
		ID:          100,
		Amount:      5000000,
		Outstanding: 5500000,
		Schedule:    []domain.Payment{{Week: 1, Amount: 110000}, {Week: 2, Amount: 110000}},
	}

	// Simulate no missed payments, should not be delinquent
	isDelinquent := service.CheckDelinquency(loan)
	assert.False(t, isDelinquent)

	// Simulate missed payments
	loan.Schedule = []domain.Payment{{Week: 1, Amount: 0}, {Week: 2, Amount: 0}}
	loan.Outstanding = 5500000

	// Check if delinquent
	isDelinquent = service.CheckDelinquency(loan)
	assert.True(t, isDelinquent)
}
