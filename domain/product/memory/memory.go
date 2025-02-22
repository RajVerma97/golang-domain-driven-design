package memory

import (
	"github.com/RajVerma97/golang-domain-driven-design/aggregate"
	"github.com/RajVerma97/golang-domain-driven-design/domain/product"
	"github.com/google/uuid"
	"sync"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (mpr *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {

	var products []aggregate.Product

	for _, product := range mpr.products {
		products = append(products, product)
	}
	return products, nil

}
func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := mpr.products[id]; ok {
		return product, nil
	}
	return aggregate.Product{}, product.ErrProductNotFound
}

func (mpr *MemoryProductRepository) Add(newprod aggregate.Product) error {

	mpr.Lock()
	defer mpr.Unlock()
	id := newprod.GetID()
	if _, ok := mpr.products[id]; ok {
		return product.ErrProductAlreadyExists
	}

	mpr.products[id] = newprod

	return nil
}

func (mpr *MemoryProductRepository) Update(update aggregate.Product) error {

	mpr.Lock()
	defer mpr.Unlock()

	id := update.GetID()
	if _, ok := mpr.products[id]; !ok {
		return product.ErrProductNotFound
	}

	mpr.products[id] = update
	return nil
}

func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {

	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return product.ErrProductNotFound
	}

	delete(mpr.products, id)
	return nil
}
