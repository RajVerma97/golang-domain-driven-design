package mongo

import (
	"context"
	"github.com/RajVerma97/golang-domain-driven-design/aggregate"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoRepository struct {
	db       *mongo.Database
	customer *mongo.Collection
}

type mongoCustomer struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

func New(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))

	if err != nil {
		return &MongoRepository{}, err
	}

	db := client.Database("domain-driven-design")
	customer := db.Collection("customers")

	return &MongoRepository{
		db:       db,
		customer: customer,
	}, nil

}

// NewFromCustomer takes in a aggregate and converts into internal structure
func NewFromCustomer(c aggregate.Customer) mongoCustomer {
	return mongoCustomer{
		ID:   c.GetID(),
		Name: c.GetName(),
	}
}

// ToAggregate converts into a aggregate.Customer
// this could validate all values present etc
func (m mongoCustomer) ToAggregate() aggregate.Customer {
	c := aggregate.Customer{}

	c.SetID(m.ID)
	c.SetName(m.Name)

	return c

}
func (mr *MongoRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := mr.customer.FindOne(ctx, bson.M{"id": id})

	var c mongoCustomer
	err := result.Decode(&c)

	if err != nil {
		return aggregate.Customer{}, err
	}

	return c.ToAggregate(), nil

}

func (mr *MongoRepository) Add(c aggregate.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	internal := NewFromCustomer(c)
	_, err := mr.customer.InsertOne(ctx, internal)
	if err != nil {
		return err
	}
	return nil
}

func (mr *MongoRepository) Update(c aggregate.Customer) error {
	panic("to implement")
}
