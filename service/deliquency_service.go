package service

import "github.com/yuhari7/amartha_test/domain"

type DelinquencyService struct{}

func NewDelinquencyService() *DelinquencyService {
	return &DelinquencyService{}
}

func (s *DelinquencyService) CheckDelinquency(loan *domain.Loan) bool {
	missedPayments := 0
	for _, payment := range loan.Schedule {
		if payment.Amount == 0 {
			missedPayments++
		}
	}

	return missedPayments >= 2
}
