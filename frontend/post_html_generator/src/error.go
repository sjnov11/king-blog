package main

import (
	"log"
)

// Error check helper
func check(err error) {
	if err != nil {
		log.Fatalf("[Error Fatal] ====== %+v ======", err)
	}
}
