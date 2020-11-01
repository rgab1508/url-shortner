package link

import (
  "os"
  "log"
  "context"

  "github.com/gofiber/fiber"
  "github.com/rgab1508/url-shortner/utils"

  firebase "firebase.google.com/go"
  "google.golang.org/api/option"
)

type Link struct {
  Slug string `json:"slug"`
  Url string  `json:"url"`
}

func (l *Link) Save(){
  ctx := context.Background()
  sa := option.WithCredentialsFile(os.Getenv("PATH_TO_SERVICE_FILE"))
  conf := &firebase.Config{ProjectID: os.Getenv("PROJECT_ID")}
  app, err := firebase.NewApp(ctx, conf, sa)
  if err != nil {
    log.Fatalln(err)
  }

  client, err := app.Firestore(ctx)
  if err != nil {
    log.Fatalln(err)
  }
  defer client.Close()

  _, _, err = client.Collection("links").Add(ctx, map[string]interface{}{
    "slug": l.Slug,
    "url":  l.Url,
  })
  if err != nil {
    log.Fatalf("Failed adding Link : %v", err)
  }
}

/* ROUTES */

func GetLink(c *fiber.Ctx){
  id := c.Params("id")
  c.Send(id)
}

func NewLink(c *fiber.Ctx){
  link := new(Link)
  if err := c.BodyParser(link); err != nil {
    c.Status(500).Send(err)
    return
  }
  link.Slug = utils.RandomSlug(6)
  link.Save()
  c.JSON(link)
}
