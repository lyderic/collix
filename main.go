package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

/* Globals */
const (
	VERSION  = "0.1.3"
	PROGNAME = "collix"
)

func main() {
	basedir, jsonfile, dbfile := setup()
	epubs, err := index(basedir)
	c(err)
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
