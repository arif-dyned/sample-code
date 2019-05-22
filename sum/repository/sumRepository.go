package repository

import (
	"github.com/arif-dyned/sample-code/domain"
)

type sumRepository struct {
	// Conn *sql.DB
}

// NewSumRepository will create an object that represent the domain interface
// ussualy this repository stores data into MySQL database
// ...
func NewSumRepository() domain.SumRepository {
	return &sumRepository{}
}

// Index a func that represent the usecase interface
func (r *sumRepository) Index(n domain.Num) (int, error) {

	return n.Num1 + n.Num2, nil
}
