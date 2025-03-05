package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yuhari7/amartha_test/usecase"
)

type LoanHandler struct {
	Usecase *usecase.LoanUsecase
}

func NewLoanHandler(uc *usecase.LoanUsecase) *LoanHandler {
	return &LoanHandler{Usecase: uc}
}

func (h *LoanHandler) GetOutstanding(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	loanID, _ := strconv.Atoi(vars["loanID"])

	outstanding, err := h.Usecase.GetOutstanding(loanID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]float64{"outstanding": outstanding})
}

func (h *LoanHandler) MakePayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	loanID, _ := strconv.Atoi(vars["loanID"])
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)

	err := h.Usecase.MakePayment(loanID, amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
