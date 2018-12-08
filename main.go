package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

/* Globals */
var (
	verbose bool
)

func main() {
	flag.BoolVar(&verbose, "v", false, "be verbose")
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
	err = list(basedir)
	c(err)
}

func usage() {
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
