package entities

type Loan struct {
	ID            int
	Amount        float64
	InterestRate  float64
	TermWeeks     int
	WeeklyPayment float64
	Payments      []Payment
}
