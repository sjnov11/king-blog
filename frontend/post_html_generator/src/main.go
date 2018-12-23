package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

const (
	// PostDir : location of blog posts md
	PostDir = "../home/public/blog/markdown/"
	HTMLDir = "../home/public/blog/"
	WebDir  = "/blog/"
)

func main() {
	files, err := ioutil.ReadDir(PostDir)
	if err != nil {
		log.Println("[main] ", err)
		log.Fatal(err)
	}
	var metaDataList []*MetaData

	for _, file := range files {
		path := PostDir + file.Name()
		if filepath.Ext(path) != ".md" {
			continue
		}
		content, err := ioutil.ReadFile(path)
		if err != nil {
			log.Println("[main] ", err, "(", path, ")")
			log.Fatal(err)
		}

		metaData, err := buildMetaData(content, file.Name())
		if err != nil {
			log.Println("[main] ", err, "(", path, ")")
			log.Fatal(err)
		}

		err = generatePostHTML(path, metaData.Slug)
		if err != nil {
			log.Println("[main] ", err, "(", path, ")")
			log.Fatal(err)
		}
		metaDataList = append(metaDataList, metaData)
	}
	generateJSON(metaDataList)
}
