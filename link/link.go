package link

import (
  "github.com/gofiber/fiber"
  "github.com/rgab1508/url-shortner/utils"
)

type Link struct {
  Slug string `json:"slug"`
  Url string  `json:"url"`
}

func GetLink(c *fiber.Ctx){
  c.Send("This is a Link")
}

func NewLink(c *fiber.Ctx){
  //var link Link
  link := new(Link)
  if err := c.BodyParser(link); err != nil {
    c.Status(500).Send(err)
    return
  }
  link.Slug = utils.RandomSlug(6)
    c.JSON(link)
}
