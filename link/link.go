package link

import (
  "os"
  "time"
  //"fmt"
  "log"
  "context"

  "github.com/gofiber/fiber"
  "github.com/rgab1508/url-shortner/utils"

  firebase "firebase.google.com/go"
  "google.golang.org/api/option"
  "google.golang.org/api/iterator"
)

type Link struct {
  ID string `json:"id"`
  Url string  `json:"url"`
  Timestamp time.Time `json:"timestamp"`
}

func (l *Link) Save() error {
  ctx := context.Background()
  sa := option.WithCredentialsFile(os.Getenv("PATH_TO_SERVICE_FILE"))
  conf := &firebase.Config{ProjectID: os.Getenv("PROJECT_ID")}
  app, err := firebase.NewApp(ctx, conf, sa)
  if err != nil {
    return err
  }

  client, err := app.Firestore(ctx)
  if err != nil {
    return err
  }
  defer client.Close()

  _, _, err = client.Collection("links").Add(ctx, map[string]interface{}{
    "id": l.ID,
    "url":  l.Url,
    "timestamp": l.Timestamp,
  })
  if err != nil {
    return err
  }

  return nil
}

/* ROUTES */

func GetLink(c *fiber.Ctx){
  id := c.Params("id")

  ctx := context.Background()
  sa := option.WithCredentialsFile(os.Getenv("PATH_TO_SERVICE_FILE"))
  conf := &firebase.Config{ProjectID: os.Getenv("PROJECT_ID")}
  app, err := firebase.NewApp(ctx, conf, sa)
  if err != nil {
    c.Status(500).Send(err)
  }

  client, err := app.Firestore(ctx)
  if err != nil {
    c.Status(500).Send(err)
  }
  defer client.Close()

  iter := client.Collection("links").Where("id", "==", id).Documents(ctx)
  doc, err := iter.Next()
  if err  == iterator.Done {
    c.Status(404).Send("Link Does not Exists.")
    return
  }
  c.JSON(doc.Data())
}

func NewLink(c *fiber.Ctx){
  link := new(Link)
  if err := c.BodyParser(link); err != nil {
    c.Status(500).Send(err)
    return
  }
  link.ID = utils.RandomSlug(6)
  link.Timestamp = time.Now()
  err := link.Save()
  if err != nil {
    log.Fatalln(err)
    c.Status(500).Send("Error Storing Url")
    return
  }
  c.JSON(link)
}
