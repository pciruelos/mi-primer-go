package main

import "github.com/gofiber/fiber/v2"

type User struct {
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
	user := User{
		c.BodyParser(&user); err != nil
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hola")
	})
	app.Get("/user", handleUser)
	app.Post("/create", createUser)

	app.Listen(":3000")
}
