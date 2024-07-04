package main

import (
	db "go-discord-clone/configs"
	"go-discord-clone/routers"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	db.Connect()

	app := fiber.New()

	routers.SetupRouter(app)

	app.Listen(":3000")
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		os.Exit(1)
	}
}
