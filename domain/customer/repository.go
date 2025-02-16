package customer

import (
	"errors"
	"github.com/RajVerma97/golang-domain-driven-design/aggregate"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound = errors.New("customer was not found")
	ErrAddCustomer      = errors.New("failed to add the customer")
	ErrUpdateCustomer   = errors.New("failed to update the customer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
