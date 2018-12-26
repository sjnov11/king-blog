package main

import (
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

	var postMetaList []*PostMeta

	for _, file := range files {
		filePath := MarkDownDir + file.Name()
		if filepath.Ext(filePath) != ".md" {
			continue
		}

		content, err := ioutil.ReadFile(filePath)
		check(err)

		// Get metadata of post
		postMeta, err := buildPostMeta(content, file.Name())
		check(err)
		postMetaList = append(postMetaList, postMeta)

		// Generate HTML using markdown
		err = generatePostHTML(filePath, postMeta.Slug)
		check(err)
	}
	generateJSON(postMetaList)
}
