package handlers

import (
	"errors"
	"synapsis-be-test/db"
	"synapsis-be-test/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CartBody struct {
	ProductID int `json:"productId" validate:"required"`
	Quantity  int `json:"quantity" validate:"required,min=1"`
}
type CartListItems struct {
	Product  models.Product `json:"product"`
	Quantity uint           `json:"quantity"`
}

func AddProductToCart(c *fiber.Ctx) error {
	userId, ok := c.Locals("user").(string)
	if !ok {
		return ResponseError(c, "Failed to get user ID from context")
	}
	body := &CartBody{}

	if err := c.BodyParser(body); err != nil {
		return ResponseBadRequest(c, "Invalid cart data")
	}
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return ResponseBadRequest(c, "Validation error: "+err.Error())
	}
	cart := &models.Cart{
		UserID: userId,
	}
	if err := db.DB.FirstOrCreate(cart).Error; err != nil {
		return ResponseError(c, "Failed to create user cart")
	}
	product := &models.Product{}
	if err := db.DB.First(product, body.ProductID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ResponseNotFound(c, "Product not found")
		}
		return ResponseError(c, "Failed to find product")
	}

	var cartItem models.CartItems
	if err := db.DB.Where("cart_id = ? AND product_id = ? ", cart.ID, product.ID).First(&cartItem).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cartItem = models.CartItems{
				CartID:    cart.ID,
				ProductID: product.ID,
				Quantity:  uint(body.Quantity),
			}
			if err := db.DB.Create(&cartItem).Error; err != nil {
				return ResponseError(c, "Failed to add product to cart")
			}
		}
		return ResponseSuccess(c, "Success add product to cart", nil)
	}
	cartItem.Quantity += uint(body.Quantity)
	if err := db.DB.Save(&cartItem).Error; err != nil {
		return ResponseError(c, "Failed to update product quantity")
	}

	return ResponseSuccess(c, "Success add product to cart", nil)
}

func GetProductInCart(c *fiber.Ctx) error {
	userId, ok := c.Locals("user").(string)
	if !ok {
		return ResponseError(c, "Failed to get user ID from context")
	}

	var cart models.Cart
	if err := db.DB.Where("user_id = ?", userId).First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ResponseNotFound(c, "Cart not found for the user")
		}
		return ResponseError(c, "Failed to find cart for the user")
	}
	var cartItems []models.CartItems
	if err := db.DB.Where("cart_id = ?", cart.ID).Find(&cartItems).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ResponseNotFound(c, "Cart not found for the user")
		}
		return ResponseError(c, "Failed to find cart for the user")
	}
	var cartListItems []CartListItems
	for _, item := range cartItems {
		var product models.Product
		if err := db.DB.First(&product, item.ProductID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				continue
			}
		}
		productItems := CartListItems{
			Product:  product,
			Quantity: item.Quantity,
		}
		cartListItems = append(cartListItems, productItems)
	}
	return ResponseSuccess(c, "Success retrive cart data", cartListItems)
}

func RemoveProductFromCart(c *fiber.Ctx) error {
	userId, ok := c.Locals("user").(string)
	if !ok {
		return ResponseError(c, "Failed to get user ID from context")
	}
	var body CartBody
	if err := c.BodyParser(&body); err != nil {
		return ResponseBadRequest(c, "Invalid cart data "+err.Error())
	}

	if body.ProductID == 0 {
		return ResponseBadRequest(c, "Invalid product id")
	}

	var cart models.Cart
	if err := db.DB.Where("user_id = ?", userId).First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ResponseNotFound(c, "Cart not found for the user")
		}
		return ResponseError(c, "Failed to find cart for the user")
	}
	var cartItems models.CartItems
	if err := db.DB.Where("product_id = ?", body.ProductID).First(&cartItems).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ResponseNotFound(c, "Product not found in cart")
		}
		return ResponseError(c, "Failed to find product int cart")
	}
	if err := db.DB.Delete(&cartItems).Error; err != nil {
		return ResponseError(c, "Failed to remove product item")
	}
	return ResponseSuccess(c, "Item removed from cart", nil)
}

func Checkout(c *fiber.Ctx) error {
	userID, ok := c.Locals("user").(string)
	if !ok {
		return ResponseError(c, "Failed to get user ID from context")
	}

	cart := &models.Cart{}
	if err := db.DB.Where("user_id = ?", userID).First(cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ResponseNotFound(c, "Cart not found for the user")
		}
		return ResponseError(c, "Failed to find cart for the user")
	}

	var cartItems []models.CartItems
	if err := db.DB.Where("cart_id = ?", cart.ID).Find(&cartItems).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ResponseNotFound(c, "Product not found in cart")
		}
		return ResponseError(c, "Failed to find product int cart")
	}
	if len(cartItems) == 0 {
		return ResponseSuccess(c, "Cart is empty", cartItems)
	}
	var totalAmount float64
	for _, item := range cartItems {
		product := &models.Product{}
		if err := db.DB.First(product, item.ProductID).Error; err != nil {
			return ResponseError(c, "Failed to find product")
		}
		totalAmount += float64(item.Quantity) * product.Price
	}
	transaction := &models.Transaction{
		CartID: cart.ID,
		Amount: totalAmount,
		Status: "success",
	}
	if err := db.DB.Create(transaction).Error; err != nil {
		return ResponseError(c, "Failed to create transaction")
	}
	if err := db.DB.Delete(&cartItems).Error; err != nil {
		return ResponseError(c, "Failed to clear cart")
	}
	return ResponseSuccess(c, "Checkout success", transaction)
}
