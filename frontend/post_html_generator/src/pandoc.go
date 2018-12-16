package main

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
)

func generateHTML(path string, slug string) error {
	fmt.Println(slug)
	cmd := exec.Command("pandoc",
		path, "-f", "markdown",
		"-t", "html", "-s", "-o",
		PostDir+slug+".html", "--mathjax")

	err := cmd.Run()
	if err != nil {
		log.Println("[pandoc] ", err)
		return errors.New("Fail to generate HTML")
	}
	return nil
}
