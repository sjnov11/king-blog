package main

import (
	"os/exec"
)

const (
	TemplatePath = "./assets/template/post.html"
	HTMLDir      = "../home/public/blog/"
)

func generatePostHTML(path string, slug string) {
	cmd := exec.Command("pandoc",
		path, "-f", "markdown",
		"-t", "html", "-s", "-o",
		HTMLDir+slug+".html", "--mathjax",
		"--template="+TemplatePath)
	err := cmd.Run()
	check(err)
}
