package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/elastic/go-elasticsearch/v7"
	// "github.com/elastic/go-elasticsearch/v7/esapi"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/DB", DB)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}

func DB(c echo.Context) error {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://elasticsearch:9200",
		},
	}
	// elasticsearchのクライアント生成
	es, err := elasticsearch.NewClient(cfg)
	// 接続できなかった場合
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	// elasticsearchの構成情報を取得できなかった場合
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	return c.JSON(http.StatusOK, res)

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
	return c.JSON(http.StatusOK, prof)
}
