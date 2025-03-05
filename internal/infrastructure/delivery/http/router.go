package http

import (
	"github.com/gorilla/mux"
)

func NewRouter(loanController LoanController) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/loans/{id}/outstanding", loanController.GetOutstanding).Methods("GET")
	r.HandleFunc("/api/loans/{id}/delinquency", loanController.IsDelinquent).Methods("GET")
	r.HandleFunc("/api/loans/{id}/payments", loanController.MakePayment).Methods("POST")

	return r
}
