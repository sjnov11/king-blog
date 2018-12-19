package main

import (
	"encoding/json"
	"html/template"
	"log"
	"os"
)

const (
	ListTemplatePath = "./assets/template/list.html"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func generateJSON(metaDataList []*MetaData) {
	f, err := os.Create("result.json")
	check(err)
	defer f.Close()
	b, err := json.Marshal(metaDataList)
	check(err)
	_, err = f.Write(b)
	check(err)
}

func generateListHTML(metaDataList []*MetaData) {

	t, err := template.ParseFiles(ListTemplatePath)
	check(err)

	f, err := os.Create("result.html")
	check(err)
	defer f.Close()

	data := struct {
		List []*MetaData
	}{
		List: metaDataList,
	}
	err = t.Execute(f, data)
	check(err)

}
