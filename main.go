package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

/* Globals */
const (
	VERSION = "0.0.3"
)

var (
	debug bool
)

func main() {
	var outfile string
	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.StringVar(&outfile, "o", "/tmp/listing.json", "JSON output `file`")
	flag.Parse()
	if len(flag.Args()) == 0 {
		usage()
	}
	basedir, err := filepath.Abs(flag.Args()[0])
	c(err)
	if _, err := os.Stat(basedir); os.IsNotExist(err) {
		fmt.Printf("Base directory %q not found!\n", basedir)
		return
	}
	fmt.Println("Base directory:", basedir)
	fmt.Printf("Please wait...")
	err, epubs := list(basedir)
	fmt.Printf("\r                     \r")
	fmt.Printf("%d epubs found.\n", len(epubs))
	jsonOutput, err := json.Marshal(epubs)
	c(err)
	err = ioutil.WriteFile(outfile, jsonOutput, 0644)
	c(err)
}

func usage() {
	fmt.Printf("collix v.%s - (c) Lyderic Landry, London 2018\n", VERSION)
	fmt.Println("Usage: collix <option> <directory>")
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}

func c(err error) {
	if err != nil {
		panic(err)
	}
}
