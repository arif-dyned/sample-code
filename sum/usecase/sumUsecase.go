package usecase

import (
	"time"

	"github.com/arif-dyned/sample-code/domain"
)

type sumUsecase struct {
	sumRepository domain.SumRepository
	timeout       time.Duration
}

// NewSumUsecase will create new an object representation of domain interface
func NewSumUsecase(or domain.SumRepository, timeout time.Duration) domain.SumUsecase {
	return &sumUsecase{
		sumRepository: or,
		timeout:       timeout,
	}
}

// Fetch returns questions with it's related repository
// Note: when using inmem repo, the usecase is a bit of useless
func (u *sumUsecase) Index(num domain.Num) (int, error) {
	return u.sumRepository.Index(num)
}
