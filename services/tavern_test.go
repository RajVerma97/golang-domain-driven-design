package services

import (
	"testing"

	"github.com/RajVerma97/golang-domain-driven-design/aggregate"
	"github.com/google/uuid"
)

func Test_Tavern(t *testing.T) {
	// Create OrderService
	products := init_products(t)

	os, err := NewOrderService(
		WithMongoCustomerRepository("mongodb+srv://rajneeshkumar2545:2O2JAZ7kfaHOjkbT@stock-glimpse-db.ypstj.mongodb.net/golang-domain-driven-design?retryWrites=true&w=majority"),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	cust, err := aggregate.NewCustomer("Percy")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	// Execute Order
	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}

}
