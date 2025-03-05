package domain

type Loan struct {
	ID           int
	Amount       float64
	InterestRate float64
	Term         int
	Schedule     []Payment
	Outstanding  float64
	Borrower     Borrower
}
