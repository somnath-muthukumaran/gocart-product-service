package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/niklaus-mikael/gocart/product-service/internal/database"
	"github.com/niklaus-mikael/gocart/product-service/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(app *fiber.App, client *mongo.Client) {
	fmt.Println("registered")
	productService := services.NewProductService(database.GetProductCollection(client))
	app.Get("/api/products", productService.FindProductById)
}
