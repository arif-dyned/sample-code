package provider

import (
	"time"

	"github.com/arif-dyned/sample-code/fibo/delivery/http"
	"github.com/arif-dyned/sample-code/fibo/repository"
	"github.com/arif-dyned/sample-code/fibo/usecase"
	"github.com/gorilla/mux"
)

var timeout = time.Second * 5

// RegisterFiboService register and initialize fibo's services
// The service includes delivery, repository and usecase
// ussually set connection to DB also in here
// @param r *mux.Router
func RegisterFiboService(r *mux.Router) {
	http.NewFiboHandler(r, usecase.NewFiboUsecase(repository.NewFiboRepository(), timeout))
}
