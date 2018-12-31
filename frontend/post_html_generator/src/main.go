package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

const (
	MarkDownDir = "../home/public/blog/markdown/"
	WebDir      = "/blog/"
)

func main() {
	files, err := ioutil.ReadDir(MarkDownDir)
	check(err)

	var postMetaSlice []*PostMeta

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".md" {
			continue
		}
		filePath := MarkDownDir + file.Name()
		fmt.Println("[Processing] " + file.Name())
		content, err := ioutil.ReadFile(filePath)
		check(err)

		// Get metadata of post
		postMeta := buildPostMeta(content, file.Name())
		postMetaSlice = append(postMetaSlice, postMeta)

		// Generate HTML using markdown
		generatePostHTML(filePath, postMeta.Slug)
	}
	generateJSON(postMetaSlice)
}
