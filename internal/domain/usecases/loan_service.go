package usecases

import (
	"errors"
	"fmt"

	"github.com/yuhari7/amartha_test/internal/domain/entities"
	"github.com/yuhari7/amartha_test/internal/domain/interfaces"
)

type LoanService struct {
	LoanRepository interfaces.LoanRepository
}

func (s *LoanService) GetOutstanding(loanID int) (float64, error) {
	loan, err := s.LoanRepository.GetLoanByID(loanID)
	if err != nil {
		return 0, err
	}

	outstanding := loan.Amount + (loan.Amount * loan.InterestRate / 100)
	for _, payment := range loan.Payments {
		outstanding -= payment.Amount
	}

	return outstanding, nil
}

func (s *LoanService) IsDelinquent(loanID int) (bool, error) {
	loan, err := s.LoanRepository.GetLoanByID(loanID)
	if err != nil {
		return false, err
	}

	missedPayments := 0
	for _, payment := range loan.Payments {
		if payment.Amount == 0 {
			missedPayments++
		} else {
			missedPayments = 0
		}

		if missedPayments >= 2 {
			return true, nil
		}
	}

	return false, nil
}

func (s *LoanService) MakePayment(loanID int, amount float64) (string, error) {
	loan, err := s.LoanRepository.GetLoanByID(loanID)
	if err != nil {
		return "", err
	}

	// validation if less or equal to zero
	if amount <= 0 {
		return "", errors.New("payment amount must be positive")
	}

	// Calculate the total outstanding balance
	totalPaid := float64(0)
	for _, payment := range loan.Payments {
		totalPaid += payment.Amount
	}

	totalWithInterest := loan.Amount + (loan.Amount * loan.InterestRate / 100)
	outstandingBalance := totalWithInterest - totalPaid

	// Validate if loan is fully paid
	if outstandingBalance == 0 {
		return "", errors.New("your loan is 0: loan term completed, no more payments allowed")
	}

	// Validate that the loan term is not completed
	if len(loan.Payments) >= loan.TermWeeks {
		return "", errors.New("loan term completed, no more payments allowed")
	}

	// Validate that payment is not equal the weekly payment
	if amount != loan.WeeklyPayment {
		return "", errors.New("invalid payment amount, should be exactly the weekly payment amount")
	}

	currentWeek := len(loan.Payments) + 1
	loan.Payments = append(loan.Payments, entities.Payment{Week: currentWeek, Amount: amount})
	err = s.LoanRepository.Save(loan)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("You paid for week %d", currentWeek), nil
}
