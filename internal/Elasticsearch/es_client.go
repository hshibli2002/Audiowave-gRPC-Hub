package Elasticsearch

import (
	"Audiowave-gRPC-Hub/config"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"io"
	"log"
)

var ESClient *elasticsearch.Client

func InitElasticsearch() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %s", err)
	}

	esConfig := elasticsearch.Config{
		CloudID: cfg.ES_CLOUD_ID,
		APIKey:  cfg.ES_API_KEY,
	}

	ESClient, err = elasticsearch.NewClient(esConfig)
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}

	//infoResponse, err := ESClient.Info()
	//if err != nil {
	//	log.Fatalf("Error getting response: %s", err)
	//}
	//
	//fmt.Println(infoResponse)

	res, err := ESClient.Ping()
	if err != nil {
		log.Fatalf("Error pinging Elasticsearch: %s", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Error closing the response body: %s", err)
		}
	}(res.Body)

	fmt.Println("Elasticsearch client successfully connected, response status:", res.StatusCode)
}
