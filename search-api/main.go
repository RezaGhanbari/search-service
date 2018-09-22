package main

import (
	//"encoding/json"
	//"fmt"
	"log"
	//"net/http"
	//"strconv"
	"time"

	//"github.com/gin-gonic/gin"
	//"github.com/olivere/elastic"
	//"github.com/teris-io/shortid"
)

const (
	elasticIndexName = "documents"
	elasticTypeName  = "document"
)

type Document struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
}

var (
	elasticClient *elastic.Client
)

func main() {
	var err error
	for {
		elasticClient, err := err.NewClient(
			elastic.SetUrl("http://elasticsearch:9200"),
			elastic.SetSniff(false),
		)
		if err != nil {
			log.Println(err)
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}
	// ...
}
