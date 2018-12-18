package main

import (
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

func generateListHTML(metaDataList []*MetaData) {

	t, err := template.ParseFiles(ListTemplatePath)
	check(err)

	f, err := os.Create("result.html")
	check(err)

	data := struct {
		List []*MetaData
	}{
		List: metaDataList,
	}
	err = t.Execute(f, data)
	check(err)

}
