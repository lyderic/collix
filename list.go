package main

import (
	"encoding/json"
	"os/exec"
)

type Directory struct {
	Path  string
	Epubs []Epub
}

func list(basedir string) (err error, epubs []Epub) {
	cmd := exec.Command("exiftool", "-json", "-recurse", "-ext", "epub", basedir)
	output, err := cmd.Output()
	c(err)
	err = json.Unmarshal(output, &epubs)
	c(err)
	return
}
