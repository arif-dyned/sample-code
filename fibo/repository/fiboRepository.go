package repository

import (
	"context"
	"errors"

	"github.com/arif-dyned/sample-code/domain"
)

type fiboRepository struct {
	// Conn *sql.DB
}

// NewFiboRepository will create an object that represent the domain interface
// ussualy this repository stores data into MySQL database
// ...
func NewFiboRepository() domain.FiboRepository {
	return &fiboRepository{}
}

// Index a func that represent the usecase interface
func (r *fiboRepository) Index(ctx context.Context, n int) ([]int, error) {
	t1 := 0
	t2 := 1
	f := make([]int, n+1, n+2)

	if n <= 1 {
		return nil, errors.New("num must more than 1")
	}

	for i := 1; i <= n; i++ {
		if i == 1 {
			f[i] = t1
			continue
		}
		if i == 2 {
			f[i] = t2
			continue
		}

		f[i] = t1 + t2
		t1 = t2
		t2 = f[i]
	}

	return f, nil
}
