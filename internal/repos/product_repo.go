package repos

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	*Repository
}

func NewProductRepository(collection *mongo.Collection) *ProductRepository {
	return &ProductRepository{
		Repository: NewRepository(collection),
	}
}
