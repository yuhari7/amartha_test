package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yuhari7/amartha_test/internal/domain/usecases"
)

type LoanController struct {
	LoanService usecases.LoanService
}

func (c *LoanController) GetOutstanding(w http.ResponseWriter, r *http.Request) {
	loanID, _ := strconv.Atoi(mux.Vars(r)["id"])
	outstanding, err := c.LoanService.GetOutstanding(loanID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(outstanding)
}

func (c *LoanController) IsDelinquent(w http.ResponseWriter, r *http.Request) {
	loanID, _ := strconv.Atoi(mux.Vars(r)["id"])
	delinquent, err := c.LoanService.IsDelinquent(loanID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(delinquent)
}

func (c *LoanController) MakePayment(w http.ResponseWriter, r *http.Request) {
	loanID, _ := strconv.Atoi(mux.Vars(r)["id"])
	var payment struct {
		Amount float64 `json:"amount"`
	}
	json.NewDecoder(r.Body).Decode(&payment)
	message, err := c.LoanService.MakePayment(loanID, payment.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(message)
}
