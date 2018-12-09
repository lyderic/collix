package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

func index(basedir string) (err error, epub []Epub) {
	return indexByTab(basedir)
	//return indexByJson(basedir)
}

func indexByTab(basedir string) (err error, epubs []Epub) {
	cmd := exec.Command("exiftool", "-table", "-recurse", "-ext", "epub",
		"-FileName", "-Directory",
		"-Title", "-Creator", "-Language",
		"-Publisher", "-Description", basedir)
	output, err := cmd.Output()
	c(err)
	var epub Epub
	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()
		bits := strings.Split(line, "	") /* TAB */
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

func indexByJson(basedir string) (err error, epubs []Epub) {
	cmd := exec.Command("exiftool", "-json", "-sep", ",", "-recurse", "-ext", "epub", basedir)
	output, err := cmd.Output()
	c(err)
	err = json.Unmarshal(output, &epubs)
	c(err)
	return
}