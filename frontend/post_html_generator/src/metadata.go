package main

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
)

type PostMeta struct {
	Title string   `json:"title"`
	Slug  string   `json:"slug"`
	Date  string   `json:"date"`
	Tags  []string `json:"tags"`
	URI   string   `json:"uri"`
}

func (m *PostMeta) setURI(uri string) {
	m.URI = uri
	return
}

// Build PostMeta structure of post
func buildPostMeta(post []byte, fileName string) (*PostMeta, error) {
	uri := WebDir + fileName
	postMetaJSON, err := parse(post)
	// if err != nil {
	// 	log.Println("[metadata] ", err)
	// 	return nil, errors.New("Fail to build metadata structure")
	// }
	check(err)

	postMeta := new(PostMeta)
	if err := json.Unmarshal(postMetaJSON, postMeta); err != nil {
		log.Println("[metadata]", err)
		return nil, errors.New("Fail to build metadata structure")
	}

	uri = strings.Replace(uri, fileName, postMeta.Slug, 1)
	postMeta.setURI(uri)

	return postMeta, nil
}
