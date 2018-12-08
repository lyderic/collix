package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func list(basedir string) (err error) {
	var listing []string
	err = filepath.Walk(basedir, func(path string, fifo os.FileInfo, err error) error {
		c(err)
		if fifo.IsDir() {
			return err
		}
		if strings.HasSuffix(path, ".epub") {
			listing = append(listing, path)
		}
		return err
	})
	fmt.Printf("%d epubs found.\n", len(listing))
	return
}
