package provider

import (
	"time"

	"github.com/arif-dyned/sample-code/sum/delivery/http"
	"github.com/arif-dyned/sample-code/sum/repository"
	"github.com/arif-dyned/sample-code/sum/usecase"
	"github.com/gorilla/mux"
)

// RegisterSumService register and initialize sum's services
// The service includes delivery, repository and usecase
// ussually set connection to DB also in here
// @param r *mux.Router
func RegisterSumService(r *mux.Router, timeout time.Duration) {
	http.NewSumHandler(r, usecase.NewSumUsecase(repository.NewSumRepository(), timeout))
}
