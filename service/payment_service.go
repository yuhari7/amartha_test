package service

import (
	"errors"

	"github.com/yuhari7/amartha_test/domain"
)

type PaymentService struct{}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (s *PaymentService) MakePayment(loan *domain.Loan, amount float64) error {
	if amount != loan.Schedule[0].Amount {
		return errors.New("payment amount does not match the schedule")
	}

	loan.Outstanding -= amount
	loan.Schedule = loan.Schedule[1:]

	if len(loan.Schedule) == 0 {
		loan.Outstanding = 0
	}

	return nil
}
