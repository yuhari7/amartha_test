package main

import (
	"fmt"
	"net/http"

	"github.com/yuhari7/amartha_test/internal/domain/entities"
	usecases "github.com/yuhari7/amartha_test/internal/domain/usecases"
	delivery "github.com/yuhari7/amartha_test/internal/infrastructure/delivery/http"
	"github.com/yuhari7/amartha_test/internal/infrastructure/repository"
)

func main() {
	loanRepo := repository.NewInMemoryLoanRepository()
	initializeSampleData(loanRepo)

	loanService := usecases.LoanService{LoanRepository: loanRepo}
	loanController := delivery.LoanController{LoanService: loanService}

	r := delivery.NewRouter(loanController)

	fmt.Println("Server is running and connected on port 8080")
	http.ListenAndServe(":8080", r)
}

func initializeSampleData(loanRepo *repository.InMemoryLoanRepository) {
	loanRepo.Save(entities.Loan{
		ID:            100,
		Amount:        5000000,
		InterestRate:  10,
		TermWeeks:     50,
		WeeklyPayment: 110000,
		Payments:      []entities.Payment{},
	})
}
