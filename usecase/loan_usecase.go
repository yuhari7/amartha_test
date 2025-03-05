package usecase

import (
	"errors"

	"github.com/yuhari7/amartha_test/domain"
)

type LoanUsecase struct {
	LoanRepo LoanRepository
}

type LoanRepository interface {
	GetLoanByID(id int) (*domain.Loan, error)
	UpdateLoan(loan *domain.Loan) error
}

func NewLoanUsecase(repo LoanRepository) *LoanUsecase {
	return &LoanUsecase{
		LoanRepo: repo,
	}
}

func (uc *LoanUsecase) GetOutstanding(loanID int) (float64, error) {
	loan, err := uc.LoanRepo.GetLoanByID(loanID)
	if err != nil {
		return 0, err
	}
	return loan.Outstanding, nil
}

func (uc *LoanUsecase) IsDelinquent(loanID int) (bool, error) {
	loan, err := uc.LoanRepo.GetLoanByID(loanID)
	if err != nil {
		return false, err
	}

	return loan.Borrower.IsDelinquent, nil
}

func (uc *LoanUsecase) MakePayment(loanID int, amount float64) error {
	loan, err := uc.LoanRepo.GetLoanByID(loanID)
	if err != nil {
		return err
	}

	if amount != loan.Schedule[0].Amount {
		return errors.New("payment amount does not match the schedule")
	}

	loan.Outstanding -= amount
	loan.Schedule = loan.Schedule[1:]

	if len(loan.Schedule) == 0 {
		loan.Outstanding = 0
	}

	if loan.Outstanding > amount*2 {
		loan.Borrower.IsDelinquent = true
	} else {
		loan.Borrower.IsDelinquent = false
	}

	return uc.LoanRepo.UpdateLoan(loan)
}
