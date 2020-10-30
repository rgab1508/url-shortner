package main

import (
  "fmt"
  "github.com/gofiber/fiber"
)

func main(){
  app := fiber.New()
  app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
  fmt.Println("Server Running......")
}
