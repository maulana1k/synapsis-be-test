package handlers

import (
	"synapsis-be-test/db"
	"synapsis-be-test/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ProductBody struct {
	Name     string  `json:"name" validate:"required"`
	Price    float64 `json:"price" validate:"required,min=0"`
	Category string  `json:"category" validate:"required"`
}

func AddProduct(c *fiber.Ctx) error {
	body := &ProductBody{}

	if err := c.BodyParser(body); err != nil {
		return ResponseBadRequest(c, "Invalid product data")
	}
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return ResponseBadRequest(c, "Validation error: "+err.Error())
	}
	product := &models.Product{
		Name:     body.Name,
		Price:    body.Price,
		Category: body.Category,
	}
	if err := db.DB.Create(product).Error; err != nil {
		return ResponseError(c, "Failed to add product")
	}
	return ResponseSuccess(c, "Success store new product", product)
}
func GetByCategory(c *fiber.Ctx) error {
	cat := c.Query("category")

	var products []models.Product
	if cat == "" {
		if err := db.DB.Find(&products).Error; err != nil {
			return ResponseError(c, "Failed to get data")
		}
	} else {
		if err := db.DB.Where("category LIKE ?", "%"+cat+"%").Find(&products).Error; err != nil {
			return ResponseError(c, "Failed to get data")
		}
	}
	message := "Success list product"
	if len(products) == 0 {
		message = "Product empty"
	}
	return ResponseSuccess(c, message, products)

}
