package main

import (
	"encoding/json"
	"errors"
	"log"
)

type MetaData struct {
	Title string   `json:"title"`
	Slug  string   `json:"slug"`
	Date  string   `json:"date"`
	Tags  []string `json:"tags"`
}

func buildMetaData(post []byte) (*MetaData, error) {
	metaDataJSON, err := parse(post)
	if err != nil {
		log.Println("[metadata] ", err)
		return nil, errors.New("Fail to build metadata structure")
	}

	metaData := new(MetaData)
	if err := json.Unmarshal(metaDataJSON, metaData); err != nil {
		log.Println("[metadata]", err)
		return nil, errors.New("Fail to build metadata structure")
	}

	return metaData, nil
}
