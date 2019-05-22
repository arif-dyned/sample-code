package repository

import (
	"testing"

	"github.com/arif-dyned/sample-code/domain"
)

func (r *sumRepository) TestSum(t *testing.T) {
	data := domain.Num{}
	data.Num1 = 1
	data.Num2 = 2

	if _, err := r.Index(data); err != nil {
		t.Errorf("Invalid Sum value")
	}

}
