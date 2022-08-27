package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
)

// go get -u go.elastic.co/apm
// https://adityarama1210.medium.com/golang-application-performance-monitoring-with-elastic-apm-92b02ca33374

// cd "C:\Users\chris\Desktop\CMS GoLang\cms.golang.teste.outros\cms.golang.teste.elastic"

// go mod init github.com/chrismarsilva/cms.golang.teste.elastic
// go get github.com/elastic/go-elasticsearch/v8
// go get github.com/elastic/go-elasticsearch/esapi
// go mod tidy

// go run main.go

func main() {
	Teste01()
	// Teste02()
	Teste03()
}

func Teste03() {
	log.Println("Teste.03.Ini")
	log.SetFlags(0)

	var r map[string]interface{}

	// Initialize a client with the default settings. 
	// An `ELASTICSEARCH_URL` environment variable will be used when exported.
	// es, err := elasticsearch.NewDefaultClient()

	// cert, err := ioutil.ReadFilte("c:\\temp\\teste.crt")

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
			// "https://localhost:9200",
		},
		// CACert: cert,
		// Username: 'elastic',
		// Password: 'elastic',
	}

	es, err := elasticsearch.NewClient(cfg) // client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info() // 1. Get cluster info
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() { // Check response status
		log.Fatalf("Error: %s", res.String())
	}

	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	// Print client and server version numbers.
	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	log.Println(strings.Repeat("~", 37))

	var wg sync.WaitGroup

	// 2. Index documents concurrently
	//
	for i, title := range []string{"Test One", "Test Two", "Test Tree"} {
		wg.Add(1)

		go func(i int, title string) {
			defer wg.Done()

			// Build the request body.
			var b strings.Builder
			b.WriteString(`{"title" : "`)
			b.WriteString(title)
			b.WriteString(`"}`)

			// user := User{Name: "Cms", Age: 10}
			// requestBytes, err := json.Marshal(user)

			// Set up the request object.
			req := esapi.IndexRequest{
				Index:      "test",
				DocumentID: strconv.Itoa(i + 1),
				Body:       strings.NewReader(b.String()), // bytes.NewReader(requestBytes)
				Refresh:    "true",
			}

			// Perform the request with the client.
			res, err := req.Do(context.Background(), es) // res, err := req.Do(context.Background(), client)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
			} else {
				// Deserialize the response into a map.
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Printf("Error parsing the response body: %s", err)
				} else {
					// Print the response status and indexed document version.
					log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
				}
			}
		}(i, title)
	}
	wg.Wait()

	// cfg := esapi.GetRequest{Index: "test", DocumentID: "1"}
	// res, err := cfg.Do(context.Background(), es) // res, err := cfg.Do(context.Background(), client)
	// if err != nil {
	// 	log.Fatalf("Error getting response: %s", err)
	// }

	log.Println(strings.Repeat("-", 37))

	// 3. Search for the indexed documents
	// Build the request body.
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
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
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

	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)

	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	log.Println(strings.Repeat("=", 37))
	log.Println("Teste.03.Fim")
}

func Teste02() {
	log.Println("Teste.02.Ini")

	// Permitir a formatação personalizada da saída de log
	log.SetFlags(0)

	// Crie um objeto de contexto para as chamadas de API
	ctx := context.Background()

	// Declare uma configuração do Elasticsearch
	// cfg := elasticsearch.Config{Addresses: []string{"http://localhost:9200"}, Username: "user", Password: "pass"}

	// Instantiate a new Elasticsearch client object instance
	client, err := elasticsearch.NewDefaultClient() // NewClient(cfg)
	if err != nil {
		log.Fatalf("Elasticsearch connection error:", err)
	}

	// Instantiate a mapping interface for API response
	var mapResp map[string]interface{}
	var buf bytes.Buffer

	// Query for filtering a boolean value
	var query = `{"query": {"term": {"SomeStr" : "Some Value"}}, "size": 2}`

	// More example query strings to read and pass to Search()
	//query = `{"query": {"match_all" : {}},"size": 2}`
	//query = `{"query": {"term" : {"SomeBool": true}},"size": 3}`
	query = `{"query": {"match" : {"SomeStr": "Value"}},"size": 3}`

	// Concatenate a string from query for reading
	var b strings.Builder
	b.WriteString(query)
	read := strings.NewReader(b.String())

	fmt.Println("read:", read)
	fmt.Println("read TYPE:", reflect.TypeOf(read))
	fmt.Println("JSON encoding:", json.NewEncoder(&buf).Encode(read))

	// Attempt to encode the JSON query and look for errors
	if err := json.NewEncoder(&buf).Encode(read); err != nil {
		log.Fatalf("Error encoding query: %s", err)
		// Query is a valid JSON object
	} else {
		fmt.Println("json.NewEncoder encoded query:", read, "\n")

		// Pass the JSON query to the Golang client's Search() method
		res, err := client.Search(
			client.Search.WithContext(ctx),
			client.Search.WithIndex("test"),
			client.Search.WithBody(read),
			client.Search.WithTrackTotalHits(true),
			client.Search.WithPretty(),
		)

		// Check for any errors returned by API call to Elasticsearch
		if err != nil {
			log.Fatalf("Elasticsearch Search() API ERROR:", err)
			// If no errors are returned, parse esapi.Response object
		} else {
			fmt.Println("res TYPE:", reflect.TypeOf(res))

			// Close the result body when the function call is complete
			defer res.Body.Close()

			// Decode the JSON response and using a pointer
			if err := json.NewDecoder(res.Body).Decode(&mapResp); err != nil {
				log.Fatalf("Error parsing the response body: %s", err)
				// If no error, then convert response to a map[string]interface
			} else {
				fmt.Println("mapResp TYPE:", reflect.TypeOf(mapResp), "\n")

				// Iterate the document "hits" returned by API call
				for _, hit := range mapResp["hits"].(map[string]interface{})["hits"].([]interface{}) {

					// Parse the attributes/fields of the document
					doc := hit.(map[string]interface{})

					// The "_source" data is another map interface nested inside of doc
					source := doc["_source"]
					fmt.Println("doc _source:", reflect.TypeOf(source))

					// Get the document's _id and print it out along with _source data
					docID := doc["_id"]
					fmt.Println("docID:", docID)
					fmt.Println("_source:", source, "\n")
				} // end of response iteration

			}
		}
	}

	log.Println("Teste.02.Fim")
}

func Teste01() {
	log.Println("Teste.01.Ini")

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	log.Println(elasticsearch.Version)

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	log.Println(res)

	log.Println("Teste.01.Fim")
}
