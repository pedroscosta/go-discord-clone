package main

import (
	db "go-discord-clone/configs"
	"go-discord-clone/routers"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	db.Connect()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routers.SetupRouter(app, engine)

	app.Listen(":3000")
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		os.Exit(1)
	}
}
