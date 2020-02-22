package elasticsearch

import (
	"log"
	"os"

	"github.com/bungysheep/news-api/pkg/configs"
	"github.com/olivere/elastic/v7"
)

var (
	// ESClient - Elasticsearch client
	ESClient *elastic.Client
)

// CreateESClient - Creates elasticsearch client
func CreateESClient() error {
	log.Printf("Creating elasticsearch client...")

	esURL, err := resolveElasticsearchURL()
	if err != nil {
		return err
	}
	client, err := elastic.NewClient(elastic.SetURL(esURL), elastic.SetSniff(false), elastic.SetHealthcheck(false))

	ESClient = client

	return err
}

func resolveElasticsearchURL() (string, error) {
	esURL := os.Getenv("ELASTICSEARCH_URL")
	if esURL != "" {
		return esURL, nil
	}

	return configs.ELASTICSEARCHURL, nil
}
