package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func setup() (basedir, jsonfile, dbfile string) {
	var err error
	jsonfile = fmt.Sprintf("/tmp/%s.json", PROGNAME)
	dbfile = fmt.Sprintf("/tmp/%s.db", PROGNAME)
	flag.StringVar(&jsonfile, "json", jsonfile, "JSON output `file`")
	flag.StringVar(&dbfile, "db", dbfile, "Database `file`")
	flag.Parse()
	if len(flag.Args()) == 0 {
		usage()
	}
	basedir, err = filepath.Abs(flag.Args()[0])
	c(err)
	if _, err := os.Stat(basedir); os.IsNotExist(err) {
		fmt.Printf("Directory %q not found!\n", basedir)
		usage()
	}
	fmt.Println("Directory:", basedir)
	return
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
