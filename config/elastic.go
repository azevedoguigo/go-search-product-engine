package config

import (
	"log"

	"github.com/olivere/elastic/v7"
)

func ConnectElasticsearch() *elastic.Client {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		log.Fatalf("Error to connect Elasticsearch: %v", err)
	}
	log.Println("Connected to Elasticsearch!")
	return client
}
