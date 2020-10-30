package link

import (
  "github.com/gofiber/fiber"
)

func GetLink(c *fiber.Ctx){
  c.Send("This is a Link")
}

func NewLink(c *fiber.Ctx){
  c.Send("New Link")
}
