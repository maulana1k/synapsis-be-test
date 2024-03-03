package handlers

import (
	"synapsis-be-test/db"
	"synapsis-be-test/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RegisterBody struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func Register(c *fiber.Ctx) error {
	body := &RegisterBody{}

	if err := c.BodyParser(body); err != nil {
		return ResponseBadRequest(c, err.Error())
	}
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return ResponseBadRequest(c, "Validation error: "+err.Error())
	}
	existingUser := &models.User{}
	if err := db.DB.Where("email = ?", body.Email).First(existingUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			user := &models.User{
				ID:       uuid.New(),
				Name:     body.Name,
				Email:    body.Email,
				Password: body.Password,
			}
			if err := db.DB.Create(user).Error; err != nil {
				return ResponseError(c, "Fail to register")
			}
			claims := jwt.MapClaims{
				"ID": user.ID,
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			t, err := token.SignedString([]byte("synapsis"))
			if err != nil {
				return ResponseError(c, err.Error())
			}
			return ResponseSuccess(c, "User registered successfully", fiber.Map{"token": t})
		}
		return ResponseError(c, "Fail to check user")
	}
	return ResponseBadRequest(c, "User is already exist")
}

func Login(c *fiber.Ctx) error {
	body := &LoginBody{}

	if err := c.BodyParser(body); err != nil {
		return ResponseBadRequest(c, "Invalid request")
	}

	user := &models.User{}
	if err := db.DB.First(&user, "email = ?", body.Email).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ResponseBadRequest(c, "User not found")
		}
		return ResponseError(c, "Fail to check user")
	}

	if user.Password != body.Password {
		return ResponseBadRequest(c, "Wrong email or password")
	}
	claims := jwt.MapClaims{
		"ID": user.ID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("synapsis"))
	if err != nil {
		return ResponseError(c, err.Error())
	}
	return ResponseSuccess(c, "success", fiber.Map{"token": t})
}
