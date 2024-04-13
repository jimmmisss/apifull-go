package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Product 1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Product 1", p.Name)
	assert.Equal(t, 10.0, p.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 10)
	assert.Nil(t, p)
	assert.Equal(t, len(err), 1)
	assert.Equal(t, err[0].Error(), ErrNameIsRequired.Error())
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Product 1", 0)
	assert.Nil(t, p)
	assert.Equal(t, len(err), 1)
	assert.Equal(t, err[0].Error(), ErrPriceIsRequired.Error())
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Product 1", -10)
	assert.Nil(t, p)
	assert.Equal(t, len(err), 1)
	assert.Equal(t, err[0].Error(), ErrInvalidPrice.Error())
}

func TestProdictvalidade(t *testing.T) {
	p, err := NewProduct("Product 1", 10)
	assert.Nil(t, err)
	assert.Nil(t, p.Validate())
}
