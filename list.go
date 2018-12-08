package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Directory struct {
	Path  string
	Epubs []Epub
}

func list(basedir string) (err error, epubs []Epub) {
	var directories []Directory
	err = filepath.Walk(basedir, func(path string, fifo os.FileInfo, err error) error {
		c(err)
		if !fifo.IsDir() {
			return err
		}
		var directory Directory
		directory.Path = path
		err = process(path, &directory)
		directories = append(directories, directory)
		return err
	})
	for _,directory := range directories {
		for  _,epub := range directory.Epubs {
			epubs = append(epubs, epub)
		}
	}
	return
}

func process(path string, directory *Directory) (err error) {
	cmd := exec.Command("exiftool", "-json", "-ext", "epub", path)
	output, err := cmd.Output()
	c(err)
	if len(output) == 0 {
		return
	}
	if debug {
		fmt.Println("*************************************************")
		fmt.Println(string(output))
		fmt.Println("*************************************************")
		fmt.Println()
	}
	err = json.Unmarshal(output, &directory.Epubs)
	if err != nil {
		fmt.Printf("ERROR in %q: %s\n", path ,string(output))
		c(err)
	}
	return
}
