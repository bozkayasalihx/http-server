package main

import (
	"fmt"
	"log"

	envconfig "github.com/bozkayasalih01x/proj/config"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	validate := validator.New()

	app.Use(cors.New())

	app.Get("/test", func(ctx *fiber.Ctx) error {
		type User struct {
			ID        uint   `validate:"required,omitempty"`
			Firstname string `validate:"required"`
			Password  string `validate:"gte=10"` // gte = Greater than or equal
		}

		user := User{
			ID:        1,
			Firstname: "Fiber",
			Password:  "FiberPassword123",
		}

		if err := validate.Struct(user); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON("success time")
	})

	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}
	log.Fatal(app.Listen(fmt.Sprintf(":%v", envconfig.LoadEnv("PORT"))))
}
