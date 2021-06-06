package main

import (
  "net/http"
  "bytes"
  "context"
  "encoding/json"
  "strings"
  "log"  
  "fmt"


  "github.com/elastic/go-elasticsearch/v7"
  "github.com/elastic/go-elasticsearch/v7/esapi"

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
  // config setting for elasticsearch 
  cfg := elasticsearch.Config{
    Addresses: []string{
      "http://elasticsearch:9200",

    },
  }
  // elasticsearchのクライアント生成
  es, err := elasticsearch.NewClient(cfg)
  // elasticsearchを取得できなかった場合
  if err != nil {
    log.Fatalf("Error creating the client: %s", err)
  }
  
  res, err := es.Info()
  // elasticsearchの構成情報を取得できなかった場合
  if err != nil {
    log.Fatalf("Error getting response: %s", err)
  }

  var buf bytes.Buffer
  query := map[string]interface{}{
    "query": map[string]interface{}{
      "match": map[string]interface{}{
        "title": "test",
      },
    },
  }
  if err := json.NewEncoder(&buf).Encode(query); err != nil {
    log.Fatalf("Error encoding query: %s", err)
  }

  users = map[string]user{
    "1": user{
        ID:   "1",
        Name: "ジョナサン・ジョースター",
        Age:  22,
    },
  }

  request := esapi.IndexRequest{
    Index:      "test1",                                  // Index name
    Body:       strings.NewReader(`{"title" : "Test2"}`), // Document body
    DocumentID: "3",   
    Refresh:    "true",                                  // Refresh
  }
  res1, err1 := request.Do(context.Background(), es)
  if err1 != nil {
    log.Fatalf("Error getting response: %s", err1)
  }
  // res1.Body.Close()

  log.Println("aaa : %s", res1)
  // Perform the search request.
  res, err = es.Search(
    es.Search.WithContext(context.Background()),
    es.Search.WithIndex("test"),
    es.Search.WithBody(&buf),
    es.Search.WithTrackTotalHits(true),
    es.Search.WithPretty(),
  )
  if err != nil {
    log.Fatalf("Error getting response: %s", err)
  }

  json, err := json.Marshal(res)
  if err != nil {
    return err
  }

  fmt.Printf("%+v\n", string(json))

  return c.JSON(http.StatusOK, res1)
}