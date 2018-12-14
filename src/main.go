package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

const (
	// Root directory
	WebRoot = "./web_root/"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	path := r.URL.Path[len("/"):]
	source, err := ioutil.ReadFile(WebRoot + path)
	if err != nil {
		source, err = ioutil.ReadFile(WebRoot + path + "/index.html")
		if err != nil {
			// Redirect to 404 page
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, err)
			log.Println("(rootHandler) ", err)
			return
		}
	}

	// It doesn't load css if content type is not set
	// Set content type as css if required file's extension is css
	if filepath.Ext(path) == ".css" {
		w.Header().Set("Content-Type", "text/css")
	}

	fmt.Fprint(w, string(source))
	log.Println("(rootHandler) The requested file has been sent: ", WebRoot+path)
}

// Main function
func main() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
