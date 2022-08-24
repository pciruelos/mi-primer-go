package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
)

type User struct {
	Id       string
	Name     string
	Lastname string
}

func handleUser(c *fiber.Ctx) error {
	user := User{
		Name:     "juan",
		Lastname: "jose",
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func createUser(c *fiber.Ctx) error {
	user := User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	user.Id = uuid.NewString()
	return c.Status(fiber.StatusOK).JSON(user)
}

func main() {
	app := fiber.New()

	//middlewares
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hola")
	})

	userGroup := app.Group("/user")

	userGroup.Get("", handleUser)

	app.Post("/create", createUser)

	app.Listen(":3000")
}
