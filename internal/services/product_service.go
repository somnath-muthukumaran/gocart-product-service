package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/niklaus-mikael/gocart/product-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	*BaseService
}

func NewProductService(collection *mongo.Collection) *ProductService {
	return &ProductService{
		NewBaseService(collection),
	}
}

func (ps *ProductService) CreateProduct(ctx *fiber.Ctx) error {
	dto := new(models.Product)
	if err := ctx.BodyParser(dto); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"Error": "Unable to parse body"})
	}
	if dto.ProductName == "" {
		return ctx.Status(400).JSON(fiber.Map{"Error": "Body is mandatory"})
	}

	if dto.ProductName == "" {
		return ctx.Status(400).JSON(fiber.Map{"Error": "Product name is mandatory"})
	}

	product := &models.Product{
		ProductName:        dto.ProductName,
		Categories:         dto.Categories,
		DiscountedPrice:    dto.DiscountedPrice,
		ActualPrice:        dto.ActualPrice,
		DiscountPercentage: dto.DiscountPercentage,
		Rating:             dto.Rating,
		RatingCount:        dto.RatingCount,
	}
	product.Initialize()
	insertResult, err := ps.repository.InsertOne(product)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"Error": "Unable to add todo"})
	}
	product.Id = insertResult.InsertedID.(primitive.ObjectID)
	return ctx.Status(201).JSON(product)
}

func (ps *ProductService) UpdateProduct(ctx *fiber.Ctx) (*mongo.UpdateResult, error) {
	var _id string
	var update interface{}
	filter := bson.M{"_id": _id}
	return ps.repository.UpdateOne(filter, update)
}

func (ps *ProductService) FindProductById(ctx *fiber.Ctx) error {
	var _id string
	filter := bson.M{"_id": _id}
	var result bson.M
	err := ps.repository.FindOne(filter).Decode(&result)
	if err != nil {
		return err
	}
	return ctx.JSON(result)
}
