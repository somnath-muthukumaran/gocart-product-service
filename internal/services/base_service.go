package services

import (
	"github.com/niklaus-mikael/gocart/product-service/internal/repos"
	"go.mongodb.org/mongo-driver/mongo"
)

type BaseService struct {
	repository repos.Repository
}

func NewBaseService(collection *mongo.Collection) *BaseService {
	repo := repos.NewRepository(collection)
	return &BaseService{
		repository: *repo,
	}
}
