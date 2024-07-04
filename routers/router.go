package routers

import (
	"os"

	"go-discord-clone/handlers"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func SetupRouter(app *fiber.App, engine *html.Engine) {
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{}, "layouts/main")
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")

	auth := v1.Group("/auth")

	auth.Post("/register", handlers.RegisterUser)
	auth.Post("/login", handlers.LoginUser)

	// JWT Middleware
	v1.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
	}))

}
