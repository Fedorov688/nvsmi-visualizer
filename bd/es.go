// created by D. "Mordok" Fedorov

package bd

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"time"
)

const sleepTime = 5

type ES struct {
	Address string
	client  *elasticsearch.Client
}

func (es *ES) Init() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			es.Address,
		},
	}
	var err error
	es.client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
}

func (es *ES) reconnect(index string, payload []byte) {
	fmt.Printf("Sleep %v sec.", sleepTime)
	time.Sleep(time.Duration(sleepTime) * time.Second)
	fmt.Println("Try reconnect to sendJson.")
	es.Init()
	es.SendJson(index, payload)
}

func (es *ES) SendJson(index string, payload []byte) {
	docID := fmt.Sprintf("%x", md5.Sum(payload))
	// Set up the request object.
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: docID,
		Body:       bytes.NewReader(payload),
		Refresh:    "true",
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), es.client)
	if err != nil {
		log.Printf("Error getting response: %s", err)
		es.reconnect(index, payload)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf(
			"[%s] Error indexing document ID=%d",
			res.Status(),
		)
		es.reconnect(index, payload)
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf(
				"[%s] %s; version=%d",
				res.Status(),
				r["result"],
				int(r["_version"].(float64)),
			)
		}
	}
}
