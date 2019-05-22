package domain

import (
	"context"
)

// FiboMap represent Fibo's data model
// ...

// FieldMap binds HTTP request
// ...

// Validate perform data validation
// ...

// FiboRepository represent the fibo's repository contract
type FiboRepository interface {
	Index(ctx context.Context, num int) ([]int, error)
}

// FiboUsecase represent the fibo's usecases
type FiboUsecase interface {
	Index(ctx context.Context, num int) ([]int, error)
}
