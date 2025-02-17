package aggregate

import (
	"errors"
	"github.com/RajVerma97/golang-domain-driven-design/entity"
	"github.com/google/uuid"
)

var (
	ErrMissingValues = errors.New("missing product values")
)

type Product struct {
	item     *entity.Item
	price    float64
	quantity int
}

func NewProduct(name, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValues
	}

	item := &entity.Item{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
	}

	return Product{
		item:     item,
		price:    price,
		quantity: 0,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetName() string {
	return p.item.Name
}

func (p Product) GetPrice() float64 {
	return p.price
}
