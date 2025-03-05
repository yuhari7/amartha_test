package main

import (
	"fmt"

	"github.com/yuhari7/amartha_test/domain"
	"github.com/yuhari7/amartha_test/repository"
	"github.com/yuhari7/amartha_test/usecase"
)

func main() {
	loanRepo := repository.NewLoanRepository()

	loan := &domain.Loan{
		ID:           100,
		Amount:       5000000,
		InterestRate: 0.10,
		Term:         50,
		Outstanding:  5500000,
		Borrower:     domain.Borrower{ID: 1, Name: "John Doe"},
		Schedule:     generateSchedule(110000, 50),
	}

	loanRepo.UpdateLoan(loan)

	loanUsecase := usecase.NewLoanUsecase(loanRepo)

	fmt.Println("Loan Billing System")

	outstanding, _ := loanUsecase.GetOutstanding(100)
	fmt.Printf("Outstanding: %.2f\n", outstanding)

	err := loanUsecase.MakePayment(100, 110000)
	if err != nil {
		fmt.Println("Payment error:", err)
	}

	outstanding, _ = loanUsecase.GetOutstanding(100)
	fmt.Printf("Outstanding: %.2f\n", outstanding)

	isDelinquent, _ := loanUsecase.IsDelinquent(100)
	fmt.Printf("Is Delinquent: %v\n", isDelinquent)
}

func generateSchedule(amount float64, term int) []domain.Payment {
	schedule := make([]domain.Payment, term)
	for i := 0; i < term; i++ {
		schedule[i] = domain.Payment{Week: i + 1, Amount: amount}
	}
	return schedule
}
