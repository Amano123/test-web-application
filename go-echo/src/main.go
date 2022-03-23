package main

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "encoding/json"
  "fmt"
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
  prof := map[string]interface{}{
		"name":   "Yamada Tarou",
		"age":    18,
		"height": 178.5,
	}

	profJson, _ := json.Marshal(prof)
  fmt.Println(string(profJson))
  return  c.JSON(http.StatusOK, prof)
}