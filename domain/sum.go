package domain

import (
	"net/http"

	"github.com/mholt/binding"
)

// Num SumMap represent Sum's data model
type Num struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

// FieldMap binds HTTP request
func (o *Num) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&o.Num1: "num1",
		&o.Num2: "num2",
	}
}

// Validate perform data validation
// ...
func (o *Num) Validate(req *http.Request, errs binding.Errors) error {
	if o.Num1 < 0 {
		errs = append(errs, binding.Error{
			FieldNames:     []string{"message"},
			Classification: "ComplaintError",
			Message:        "The Num 1 field is required!",
		})
	}
	if o.Num2 < 0 {
		errs = append(errs, binding.Error{
			FieldNames:     []string{"message"},
			Classification: "ComplaintError",
			Message:        "The Num 2 field is required!",
		})
	}

	return errs
}

// SumRepository represent the sum's repository contract
type SumRepository interface {
	Index(num Num) (int, error)
}

// SumUsecase represent the sum's usecases
type SumUsecase interface {
	Index(um Num) (int, error)
}
