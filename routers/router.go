package routers

import (
	repositories "go-discord-clone/repositories"
	"os"

	"go-discord-clone/handlers"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	auth := v1.Group("/auth")

	auth.Post("/register", handlers.RegisterUser)
	auth.Post("/login", handlers.LoginUser)

	// JWT Middleware
	v1.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
	}))

	user := v1.Group("/user")

	user.Get("/:username", func(c *fiber.Ctx) error {
		user, err := repositories.GetUser(c.Params("username"))

		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(user)
	})

	user.Get("/", func(c *fiber.Ctx) error {
		users, err := repositories.GetUsers()

		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(users)
	})
}
