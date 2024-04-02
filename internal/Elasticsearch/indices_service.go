package Elasticsearch

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
)

const artistMapping = `{
  "mappings": {
    "properties": {
      "artist_id": { "type": "long" },
      "name": { 
        "type": "text",
        "fields": {
          "keyword": { "type": "keyword" }
        }
      },
      "bio": { "type": "text" },
      "followers_count": { "type": "integer" },
      "likes_count": { "type": "integer" }
    }
  }
}`

const playlistMapping = `{
  "mappings": {
    "properties": {
      "playlist_id": { "type": "long" },
      "name": { 
        "type": "text",
        "fields": {
          "keyword": { "type": "keyword" }
        }
      },
      "artist_id": { "type": "long" },
      "release_date": { "type": "date" },
      "songs_count": { "type": "integer" },
      "likes_count": { "type": "integer" }
    }
  }
}`

const songMapping = `{
  "mappings": {
    "properties": {
      "song_id": { "type": "long" },
      "name": { 
        "type": "text",
        "fields": {
          "keyword": { "type": "keyword" }
        }
      },
      "artist_id": { "type": "long" },
      "playlist_id": { "type": "long" },
      "duration": { "type": "integer" },
      "release_date": { "type": "date" },
      "likes_count": { "type": "integer" }
    }
  }
}`

const userMapping = `{
  "mappings": {
    "properties": {
      "user_id": { "type": "long" },
      "username": { "type": "keyword" },
      "password": { "type": "keyword" },
      "email": { "type": "keyword" },
      "likes_given": { "type": "integer" },
      "follows_given": { "type": "integer" }
    }
  }
}`

// InitializeIndices creates the indices in Elasticsearch
func InitializeIndices() {
	var wg sync.WaitGroup
	indices := map[string]string{
		"artists":   artistMapping,
		"playlists": playlistMapping,
		"songs":     songMapping,
		"users":     userMapping,
	}

	for indexName, mapping := range indices {
		wg.Add(1)
		go func(indexName string, mapping string) {
			defer wg.Done()
			err := createIndex(indexName, mapping)
			if err != nil {
				log.Printf("Error creating index %s: %s", indexName, err)
			}
		}(indexName, mapping)
	}

	wg.Wait()
}

func createIndex(indexName string, mapping string) error {
	res, err := ESClient.Indices.Exists([]string{indexName})
	if err != nil {
		return err
	}
	if res.StatusCode == 200 {
		//fmt.Printf("Index %s already exists.\n", indexName)
		return nil
	}

	res, err = ESClient.Indices.Create(
		indexName,
		ESClient.Indices.Create.WithBody(strings.NewReader(mapping)),
		ESClient.Indices.Create.WithContext(context.Background()),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error creating index %s: %s", indexName, res.String())
	}

	fmt.Printf("Index %s created successfully.\n", indexName)
	return nil
}
