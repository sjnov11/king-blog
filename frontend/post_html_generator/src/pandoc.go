package main

import (
	"errors"
	"log"
	"os/exec"
)

const (
	TemplatePath = "./assets/template/post.html"
	HTMLDir      = "../home/public/blog/"
)

func generatePostHTML(path string, slug string) error {
	cmd := exec.Command("pandoc",
		path, "-f", "markdown",
		"-t", "html", "-s", "-o",
		HTMLDir+slug+".html", "--mathjax",
		"--template="+TemplatePath)
	err := cmd.Run()
	if err != nil {
		log.Println("[pandoc] ", err)
		return errors.New("Fail to generate HTML")
	}
	return nil
}
