package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

/* Globals */
const (
	VERSION  = "0.1.2"
	PROGNAME = "collix"
)

func main() {
	jsonfile := fmt.Sprintf("/tmp/%s.json", PROGNAME)
	dbfile := fmt.Sprintf("/tmp/%s.db", PROGNAME)
	flag.StringVar(&jsonfile, "json", jsonfile, "JSON output `file`")
	flag.StringVar(&dbfile, "db", dbfile, "Database `file`")
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
	start := time.Now()
	fmt.Printf("Indexing, please wait...")
	err, epubs := index(basedir)
	c(err)
	fmt.Printf("\r                          \r")
	fmt.Printf("%d epubs indexed in %s.\n", len(epubs), time.Since(start))
	writeJson(jsonfile, epubs) /* see below */
	writeDb(dbfile, epubs)     /* in database.go */
}

func writeJson(jsonfile string, epubs []Epub) {
	start := time.Now()
	output, err := json.MarshalIndent(epubs, "", "  ")
	c(err)
	err = ioutil.WriteFile(jsonfile, output, 0644)
	c(err)
	fmt.Printf("JSON written to %q in %s\n", jsonfile, time.Since(start))
}

func usage() {
	fmt.Printf("%s v.%s - (c) Lyderic Landry, London 2018\n", PROGNAME, VERSION)
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
