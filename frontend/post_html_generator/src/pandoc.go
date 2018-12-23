package main

import (
	"errors"
	"log"
	"os/exec"
)

const (
	PostTemplatePath = "./assets/template/post.html"
	PostHTMLDir      = "../home/public/blog/"
)

func generatePostHTML(path string, slug string) error {
	cmd := exec.Command("pandoc",
		path, "-f", "markdown",
		"-t", "html", "-s", "-o",
		PostHTMLDir+slug+".html", "--mathjax",
		"--template="+PostTemplatePath)
	err := cmd.Run()
	if err != nil {
		log.Println("[pandoc] ", err)
		return errors.New("Fail to generate HTML")
	}
	return nil
}
