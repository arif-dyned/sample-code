package repository

import (
	"context"
	"testing"
)

func (r *fiboRepository) TestFibonacci(t *testing.T) {
	data := []struct {
		n int
	}{
		{0}, {1}, {2}, {3}, {4}, {5}, {6}, {10}, {42},
	}

	ctx := context.Background()

	for _, d := range data {
		if got, err := r.Index(ctx, d.n); err != nil {
			t.Errorf("Invalid Fibonacci value for N: %d, got: %d, want: %d", d.n, got, err)
		}
	}
}
