package services

import (
	"log"

	"github.com/RajVerma97/golang-domain-driven-design/aggregate"
	"github.com/RajVerma97/golang-domain-driven-design/domain/customer"
	"github.com/RajVerma97/golang-domain-driven-design/domain/customer/memory"
	"github.com/RajVerma97/golang-domain-driven-design/domain/product"
	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
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

// CreateOrder will chaintogether all repositories to create a order for a customer
// will return the collected price of all Products
func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	// Get the customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	// Get each Product, Ouchie, We need a ProductRepository
	var products []aggregate.Product
	var total float64
	for _, id := range productIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		total += p.GetPrice()
	}

	// All Products exists in store, now we can create the order
	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))
	// Add Products and Update Customer

	return price, nil
}
