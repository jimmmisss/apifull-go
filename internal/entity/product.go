package entity

import (
	"errors"
	"fmt"
	"time"

	"github.com.br/jimmmisss/api/pkg/entity"
)

var (
	ErrIDIsRequired    = errors.New("id is required")
	ErrInvalidID       = errors.New("invalid id")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, []error) {
	product := &Product{
		ID:        entity.NewId(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	errors := product.Validate()
	if errors != nil {
		return nil, errors
	}
	return product, nil
}

func (p *Product) Validate() []error {
	var allErrors []error

	if p.ID.String() == "" {
		allErrors = append(allErrors, ErrIDIsRequired)
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		allErrors = append(allErrors, ErrInvalidID)
	}
	if p.Name == "" {
		allErrors = append(allErrors, ErrNameIsRequired)
	}
	if p.Price == 0 {
		allErrors = append(allErrors, ErrPriceIsRequired)
	}
	if p.Price < 0 {
		allErrors = append(allErrors, ErrInvalidPrice)
	}

	if len(allErrors) > 0 {
		fmt.Printf("%d errors were found\n", len(allErrors))
		return allErrors
	}

	return nil
}