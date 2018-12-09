package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func index(basedir string) (epubs []Epub, err error) {
	start := time.Now()
	fmt.Printf("Indexing, please wait...")
	epubs, err = indexByTab(basedir)
	if err != nil {
		return
	}
	fmt.Printf("\r                          \r")
	fmt.Printf("%d epubs indexed in %s.\n", len(epubs), time.Since(start))
	return
}

func indexByTab(basedir string) (epubs []Epub, err error) {
	cmd := exec.Command("exiftool", "-table", "-recurse", "-ext", "epub",
		"-FileName", "-Directory",
		"-Title", "-Creator", "-Language",
		"-Publisher", "-Description", basedir)
	output, err := cmd.Output()
	if err != nil {
		return
	}
	var epub Epub
	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()
		bits := strings.Split(line, "\t")
		ln := len(bits)
		if ln != 7 {
			skipping(line, fmt.Sprintf("Incorrect number of fields (has %d, expecting 7)", ln))
			continue
		}
		epub.FileName = bits[0]
		epub.Directory = bits[1]
		epub.Title = bits[2]
		if epub.Title == "" {
			skipping(line, "'Title' metadata not found")
			continue
		}
		epub.Author = bits[3]
		epub.Language = bits[4]
		epub.Publisher = bits[5]
		epub.Description = bits[6]
		epubs = append(epubs, epub)
	}
	return
}

func skipping(line, reason string) {
	fmt.Println(strings.Repeat("*", 80))
	fmt.Println("Skipping the following line. Reason:", reason)
	fmt.Println(line)
	fmt.Println(strings.Repeat("*", 80))
}
