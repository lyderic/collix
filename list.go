package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func list(basedir string) (err error) {
	var epubs []Epub
	err = filepath.Walk(basedir, func(path string, fifo os.FileInfo, err error) error {
		c(err)
		if fifo.IsDir() {
			return err
		}
		if strings.HasSuffix(path, ".epub") {
			var epub Epub
			epub.Path = path
			epubs = append(epubs, epub)
		}
		return err
	})
	fmt.Printf("%d epubs found.\n", len(epubs))
	if verbose {
		fmt.Println(epubs)
	}
	return
}
