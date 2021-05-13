package main

import (
  "net/http"

  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)

type (
  user struct {
      ID   string `json:"id"`
      Name string `json:"name"`
      Age  int    `json:"age"`
  }
)

var (
  users map[string]user
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", hello)

  // Start server
  e.Logger.Fatal(e.Start(":3000"))
}

// Handler
func hello(c echo.Context) error {
  users = map[string]user{
    "1": user{
        ID:   "1",
        Name: "ジョナサン・ジョースター",
        Age:  22,
    },
    "2": user{
        ID:   "2",
        Name: "ディオ・ブランドー",
        Age:  25,
        },
    }
  return c.JSON(http.StatusOK, users)
}