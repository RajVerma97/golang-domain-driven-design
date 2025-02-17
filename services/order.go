package services

import (
	"github.com/RajVerma97/golang-domain-driven-design/domain/customer"
	"github.com/RajVerma97/golang-domain-driven-design/domain/customer/memory"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	for _, cfg := range cfgs {
		err := cfg(os)

		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// WithCustomerRepository aplies a customer repo to the OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}

}

func WithMemoryRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}
