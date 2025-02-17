package product

import (
	"errors"
	"github.com/RajVerma97/golang-domain-driven-design/aggregate"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound      = errors.New("product not found")
	ErrProductAlreadyExists = errors.New("product already exists")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(uuid.UUID) (aggregate.Product, error)
	Add(aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(uuid.UUID) error
}
