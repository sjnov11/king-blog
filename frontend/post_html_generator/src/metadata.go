package main

import (
	"encoding/json"
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
func buildPostMeta(post []byte, fileName string) *PostMeta {
	uri := WebDir + fileName
	postMetaJSON := parse(post)

	postMeta := new(PostMeta)
	err := json.Unmarshal(postMetaJSON, postMeta)
	check(err)
	uri = strings.Replace(uri, fileName, postMeta.Slug, 1)
	postMeta.setURI(uri)

	return postMeta
}
