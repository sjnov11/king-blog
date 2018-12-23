package main

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
)

type MetaData struct {
	Title string   `json:"title"`
	Slug  string   `json:"slug"`
	Date  string   `json:"date"`
	Tags  []string `json:"tags"`
	Uri   string   `json:"uri"`
}

func (m *MetaData) setUri(uri string) {
	m.Uri = uri
	return
}

func buildMetaData(post []byte, fileName string) (*MetaData, error) {
	metaDataJSON, err := parse(post)
	if err != nil {
		log.Println("[metadata] ", err)
		return nil, errors.New("Fail to build metadata structure")
	}
	uri := WebDir + fileName

	metaData := new(MetaData)
	if err := json.Unmarshal(metaDataJSON, metaData); err != nil {
		log.Println("[metadata]", err)
		return nil, errors.New("Fail to build metadata structure")
	}
	uri = strings.Replace(uri, fileName, metaData.Slug, 1)
	//uri = uri[:strings.LastIndex(uri, ".")]
	metaData.setUri(uri)
	log.Println(uri)

	return metaData, nil
}
