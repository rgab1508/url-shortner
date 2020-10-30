package main

import (
  "fmt"
  "github.com/gofiber/fiber"
  "github.com/rgab1508/url-shortner/link"
)


func SetupRoutes(app *fiber.App){
  app.Get("/api/v1/:id", link.GetLink)
  app.Post("/api/v1/new", link.NewLink)
}

func main(){
  app := fiber.New()
  SetupRoutes(app)
  fmt.Println("Server Running......")
	app.Listen(3000)
}
