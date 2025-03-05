package interfaces

import "github.com/yuhari7/amartha_test/internal/domain/entities"

type LoanRepository interface {
	GetLoanByID(id int) (entities.Loan, error)
	Save(loan entities.Loan) error
}
