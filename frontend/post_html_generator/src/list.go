package main

import (
	"encoding/json"
	"os"
)

// Generate posts' metadata JSON file
func generateJSON(postMetaList []*PostMeta) {
	wrapper := struct {
		Posts []*PostMeta `json:"posts"`
	}{
		postMetaList,
	}
	f, err := os.Create(HTMLDir + "list.json")
	check(err)
	defer f.Close()

	b, err := json.Marshal(wrapper)
	check(err)

	_, err = f.Write(b)
	check(err)
}
