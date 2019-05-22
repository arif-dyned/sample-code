package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/arif-dyned/sample-code/domain"
	"github.com/gorilla/mux"
)

// FiboHandler represent HTTP handler for Fibo
type FiboHandler struct {
	FiboUsecase domain.FiboUsecase
}

// NewFiboHandler returns HTTP delivery instance for Fibo
func NewFiboHandler(r *mux.Router, u domain.FiboUsecase) {
	h := &FiboHandler{
		FiboUsecase: u,
	}
	r.HandleFunc("/fibo", h.Index).Methods("GET")
	// TODO: add more routes
}

// Index returns all Fibos into JSON
// TODO: handle when Fibo repository is empty
func (h *FiboHandler) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	num, _ := strconv.Atoi(r.FormValue("index"))

	fibo, err := h.FiboUsecase.Index(r.Context(), num)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fibo)

}

// TODO: implement more handler
