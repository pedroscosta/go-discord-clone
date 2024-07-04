package handlers

import (
	"go-discord-clone/models"
	"go-discord-clone/repositories"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func RegisterUser(c *fiber.Ctx) error {
	user := models.User{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}

	validationErrs := validate.Struct(user)

	if validationErrs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationErrs.Error()})
	}

	if userExists, _ := repositories.GetUser(user.Username); userExists != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "username already exists"})
	}

	if err := repositories.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}

func LoginUser(c *fiber.Ctx) error {
	if (c.FormValue("username") == "") || (c.FormValue("password") == "") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "username and password are required"})
	}

	user, _ := repositories.GetUser(c.FormValue("username"))

	if user == nil || !user.CheckPassword(c.FormValue("password")) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid username or password"})
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 3).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
