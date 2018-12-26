package main

import (
	"encoding/json"
	"os"
)

// Generate posts' metadata JSON file
func generateJSON(postMetaSlice []*PostMeta) {
	wrapper := struct {
		Index []*PostMeta `json:"index"`
	}{
		postMetaSlice,
	}
	f, err := os.Create(HTMLDir + "list.json")
	defer f.Close()
	check(err)

	b, err := json.Marshal(wrapper)
	check(err)

	_, err = f.Write(b)
	check(err)
}
