package main

import (
  "net/http"
  "bytes"
  "context"
  "encoding/json"
  "strings"
  "log"  
  // "fmt"

  "github.com/elastic/go-elasticsearch/v7"
  "github.com/elastic/go-elasticsearch/v7/esapi"

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
  var (
    r  map[string]interface{}
  )
  request := esapi.IndexRequest{
    Index:      "test1",                                 // Index name
    Body:       strings.NewReader(`{"title" : "Test"}`), // Document body
    DocumentID: "3",                                     // Document id
    Refresh:    "true",                                  // Refresh
  }

  res, err = request.Do(context.Background(), es)
  if err != nil {
    log.Fatalf("Error getting response: %s", err)
  }
  defer res.Body.Close()
  
  if res.IsError() {
    log.Printf("[%s] Error indexing document ID=%d", res.Status(), 1)
  } else {
    // mapを使って処理してる
    // mapの勉強が必要
    var r map[string]interface{}
    if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
      log.Printf("Error parsing the response body: %s", err)
    } else {
      // 応答とバージョンを表示
      log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
    }
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

  // 検索してる場所
  res, err = es.Search(
    es.Search.WithContext(context.Background()),
    es.Search.WithIndex("test1"),
    es.Search.WithBody(&buf),
    es.Search.WithTrackTotalHits(true),
    es.Search.WithPretty(),
  )
  if err != nil {
    log.Fatalf("Error getting response: %s", err)
  }
  defer res.Body.Close()

  if res.IsError() {
    var e map[string]interface{}
    if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
      log.Fatalf("Error parsing the response body: %s", err)
    } else {
      // レスポンスとエラー情報の表示
      log.Fatalf("[%s] %s: %s",
        res.Status(),
        e["error"].(map[string]interface{})["type"],
        e["error"].(map[string]interface{})["reason"],
      )
    }
  }
  

  if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
    log.Fatalf("Error parsing the response body: %s", err)
  }
  // 検索に引っかかった数とレスポンス時間表示
  log.Printf(
    "[%s] %d hits; took: %dms",
    res.Status(),
    int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
    int(r["took"].(float64)),
  )
  // 検索に引っかかったID、ドキュメントの表示
  for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
    log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
  }


  return c.JSON(http.StatusOK, r)
}