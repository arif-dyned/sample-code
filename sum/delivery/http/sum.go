package http

import (
	"encoding/json"
	"net/http"

	"github.com/arif-dyned/sample-code/domain"
	"github.com/gorilla/mux"
	"github.com/mholt/binding"
)

// SumHandler represent HTTP handler for Sum
type SumHandler struct {
	SumUsecase domain.SumUsecase
}

// NewSumHandler returns HTTP delivery instance for Sum
func NewSumHandler(r *mux.Router, u domain.SumUsecase) {
	h := &SumHandler{
		SumUsecase: u,
	}
	r.HandleFunc("/sum", h.Index).Methods("POST")
	// TODO: add more routes
}

// Index returns all Sums into JSON
// TODO: handle when Sum repository is empty
func (h *SumHandler) Index(w http.ResponseWriter, r *http.Request) {
	NewSum := new(domain.Num)
	//validate data=
	errs := binding.Bind(r, NewSum)
	if errs.Handle(w) {
		return
	}
	validateData := NewSum.Validate(r, errs)
	if validateData.Error() != "" {
		json.NewEncoder(w).Encode(map[string]string{"error": validateData.Error()})
		return
	}
	Sum, err := h.SumUsecase.Index(*NewSum)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Sum)

}

// TODO: implement more handler
