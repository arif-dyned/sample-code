package usecase

import (
	"context"
	"time"

	"github.com/arif-dyned/sample-code/domain"
)

type fiboUsecase struct {
	fiboRepository domain.FiboRepository
	timeout        time.Duration
}

// NewFiboUsecase will create new an object representation of domain interface
func NewFiboUsecase(or domain.FiboRepository, timeout time.Duration) domain.FiboUsecase {
	return &fiboUsecase{
		fiboRepository: or,
		timeout:        timeout,
	}
}

// Fetch returns questions with it's related repository
// Note: when using inmem repo, the usecase is a bit of useless
func (u *fiboUsecase) Index(ctx context.Context, num int) ([]int, error) {
	return u.fiboRepository.Index(ctx, num)
}
