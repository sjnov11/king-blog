package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

const (
	// PostDir : location of blog posts
	PostDir = "../home/public/posts/"
)

func main() {
	files, err := ioutil.ReadDir(PostDir)
	if err != nil {
		log.Println("[main] ", err)
		log.Fatal(err)
	}

	for _, file := range files {
		path := PostDir + file.Name()
		if filepath.Ext(path) != ".md" {
			continue
		}
		content, err := ioutil.ReadFile(path)
		if err != nil {
			log.Println("[main] ", err)
			log.Fatal(err)
		}

		metaData, err := buildMetaData(content)
		if err != nil {
			log.Println("[main] ", err)
			log.Fatal(err)
		}
		fmt.Println(metaData.Slug)
		generateHTML(path, metaData.Slug)
	}
}
