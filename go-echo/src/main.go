package main

import (
  "net/http"
  "bytes"
  "context"
  "encoding/json"
  "log"
  // "strconv"
  // "strings"
  // "sync"
  // "time"
  // "net"
  // "crypto/tls"
  
  "fmt"

   "github.com/elastic/go-elasticsearch/v7"
  // "github.com/elastic/go-elasticsearch/v7/esapi"
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
  // log.SetFlags(0)

  // var (
  //   r  map[string]interface{}
  //   wg sync.WaitGroup
  // )

  // Initialize a client with the default settings.
  //
  // An `ELASTICSEARCH_URL` environment variable will be used when exported.
  //
  
  //es, err := elasticsearch.NewDefaultClient()

  cfg := elasticsearch.Config{
    Addresses: []string{
      "http://elasticsearch:9200",
    },
  }
  es, err := elasticsearch.NewClient(cfg)
  fmt.Println("hello")
  fmt.Println(es)
  // エラー処理
  if err != nil {
        log.Fatalf("Error creating the client: %s", err)
      }
  res, err := es.Info()
  if err != nil {
    log.Fatalf("Error getting response: %s", err)
  }
  // fmt.Println(res)

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
  // fmt.Println(res)

  // log.Printf(err)
  //return err

//   cfg := elasticsearch.Config{
//     Addresses: []string{
//       "elasticsearch:9200",
//     },
//     Username: "foo",
//     Password: "bar",
//     RetryOnStatus: []int{429, 502, 503, 504},
//     //CACert: cert,
//     Transport: &http.Transport{
//       MaxIdleConnsPerHost:   10,
//       ResponseHeaderTimeout: time.Second,
//       DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
//       TLSClientConfig: &tls.Config{
//         MinVersion:         tls.VersionTLS11,
//       },
//     }, 
//   }
//   es, err := elasticsearch.NewClient(cfg)

// elasticsearch.NewClient(cfg)

//   if err != nil {
//     log.Fatalf("Error creating the client: %s", err)
//   }

//   // 1. Get cluster info
//   //
//   res, err := es.Info()
//   if err != nil {
//     log.Fatalf("Error getting response: %s", err)
//   }
//   defer res.Body.Close()
//   // Check response status
//   if res.IsError() {
//     log.Fatalf("Error: %s", res.String())
//   }
//   // Deserialize the response into a map.
//   if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
//     log.Fatalf("Error parsing the response body: %s", err)
//   }
//   // Print client and server version numbers.
//   log.Printf("Client: %s", elasticsearch.Version)
//   log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
//   log.Println(strings.Repeat("~", 37))

//   // 2. Index documents concurrently
//   //
//   for i, title := range []string{"Test One", "Test Two"} {
//     wg.Add(1)

//     go func(i int, title string) {
//       defer wg.Done()

//       // Build the request body.
//       var b strings.Builder
//       b.WriteString(`{"title" : "`)
//       b.WriteString(title)
//       b.WriteString(`"}`)

//       // Set up the request object.
//       req := esapi.IndexRequest{
//         Index:      "test",
//         DocumentID: strconv.Itoa(i + 1),
//         Body:       strings.NewReader(b.String()),
//         Refresh:    "true",
//       }

//       // Perform the request with the client.
//       res, err := req.Do(context.Background(), es)
//       if err != nil {
//         log.Fatalf("Error getting response: %s", err)
//       }
//       defer res.Body.Close()

//       if res.IsError() {
//         log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
//       } else {
//         // Deserialize the response into a map.
//         var r map[string]interface{}
//         if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
//           log.Printf("Error parsing the response body: %s", err)
//         } else {
//           // Print the response status and indexed document version.
//           log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
//         }
//       }
//     }(i, title)
//   }
//   wg.Wait()

//   log.Println(strings.Repeat("-", 37))

//   // 3. Search for the indexed documents
//   //
//   // Build the request body.
//   var buf bytes.Buffer
//   query := map[string]interface{}{
//     "query": map[string]interface{}{
//       "match": map[string]interface{}{
//         "title": "test",
//       },
//     },
//   }
//   if err := json.NewEncoder(&buf).Encode(query); err != nil {
//     log.Fatalf("Error encoding query: %s", err)
//   }

//   // Perform the search request.
//   res, err = es.Search(
//     es.Search.WithContext(context.Background()),
//     es.Search.WithIndex("test"),
//     es.Search.WithBody(&buf),
//     es.Search.WithTrackTotalHits(true),
//     es.Search.WithPretty(),
//   )
//   if err != nil {
//     log.Fatalf("Error getting response: %s", err)
//   }
//   defer res.Body.Close()

//   if res.IsError() {
//     var e map[string]interface{}
//     if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
//       log.Fatalf("Error parsing the response body: %s", err)
//     } else {
//       // Print the response status and error information.
//       log.Fatalf("[%s] %s: %s",
//         res.Status(),
//         e["error"].(map[string]interface{})["type"],
//         e["error"].(map[string]interface{})["reason"],
//       )
//     }
//   }

//   if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
//     log.Fatalf("Error parsing the response body: %s", err)
//   }
//   // Print the response status, number of results, and request duration.
//   log.Printf(
//     "[%s] %d hits; took: %dms",
//     res.Status(),
//     int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
//     int(r["took"].(float64)),
//   )
//   // Print the ID and document source for each hit.
//   for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
//     log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
//   }

  // log.Println(strings.Repeat("=", 37))
  // //return err
  // users = map[string]user{
  //   "1": user{
  //       ID:   "1",
  //       Name: "ジョナサン・ジョースター",
  //       Age:  22,
  //   },
  //   "2": user{
  //       ID:   "2",
  //       Name: "ディオ・ブランドー",
  //       Age:  25,
  //       },
  //   }
  // return c.JSON(http.StatusOK, users)

  json, err := json.Marshal(res)
  if err != nil {
    return err
  }

  fmt.Printf("%+v\n", string(json))

  return c.JSON(http.StatusOK, string(json))
}

// func addindex (name string){

//   cfg := elasticsearch.Config{
//     Addresses: []string{
//       "http://elasticsearch:9200",
//     },
//   }
//   es, err := elasticsearch.NewClient(cfg)

//   mapping := `{
//     "settings":{
//       "number_of_shards":1,
//       "number_of_replicas":0
//     },
//     "mappings":{
//       "properties":{
//         "tags":{
//           "type":"keyword"
//         },
//         "location":{
//           "type":"txt"
//         },
//       }
//     }
//   }`
  
//   ctx := context.Background()
//   createIndex, err := es.CreateIndex(name).BodyString(mapping).Do(ctx)
//   if err != nil {
//     // Handle error
//     panic(err)
//   }
//   if !createIndex.Acknowledged {
//     // Not acknowledged
//   }
// }

// func delindex (name string) {
//   ctx := context.Background()
//   deleteIndex, err := client.DeleteIndex(name).Do(ctx)
//   if err != nil {
//     // Handle error
//     panic(err)
//   }
//   if !deleteIndex.Acknowledged {
//     // Not acknowledged
//   }
// }

// func seaindex (name string) {
//   exists, err := client.IndexExists(name).Do(context.Background())
//   if err != nil {
//     // Handle error
//   }
//   if !exists {  
//   }
//}

