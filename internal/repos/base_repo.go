package repos

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) *Repository {
	return &Repository{collection}
}

// InsertOne inserts a single document into the collection.
func (r *Repository) InsertOne(document interface{}) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(context.Background(), document)
}

// FindOne finds a single document by filter.
func (r *Repository) FindOne(filter interface{}) *mongo.SingleResult {
	return r.collection.FindOne(context.Background(), filter)
}

// UpdateOne updates a single document by filter.
func (r *Repository) UpdateOne(filter, update interface{}) (*mongo.UpdateResult, error) {
	return r.collection.UpdateOne(context.Background(), filter, update)
}

// FindAll finds all documents that match the filter.
func (r *Repository) FindAll(filter interface{}) (*mongo.Cursor, error) {
	return r.collection.Find(context.Background(), filter)
}
