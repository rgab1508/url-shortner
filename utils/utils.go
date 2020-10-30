package utils

import (
  "time"
  "math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandomSlug(n int) string {
  rand.Seed(time.Now().UnixNano())

  b := make([]rune, n)
  for i := range b {
      b[i] = letters[rand.Intn(len(letters))]
  }
  return string(b)
}

