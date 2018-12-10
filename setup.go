package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func setup() (configuration Configuration, command string) {
	file := fmt.Sprintf("%s.json", PROGNAME)
	flag.StringVar(&file, "f", file, "Configuration `file`")
	flag.Parse()
	configuration = parseConfiguration(file)
	if len(flag.Args()) == 0 {
		usage()
	}
	command = flag.Args()[0]
	if _, err := os.Stat(configuration.Directory); os.IsNotExist(err) {
		fmt.Printf("Directory %q not found!\n", configuration.Directory)
		usage()
	}
	fmt.Println(configuration)
	return
}

func parseConfiguration(file string) (configuration Configuration) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, &configuration)
	if err != nil {
		fmt.Println("Invalid JSON!")
		log.Fatal(err)
	}
	return
}

func usage() {
	fmt.Printf("%s v.%s - (c) Lyderic Landry, London 2018\n", PROGNAME, VERSION)
	fmt.Println("Usage: collix <option> command")
	fmt.Println("Commands:")
	fmt.Println("  init    initialize new database")
	fmt.Println("  info    show information on database")
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}

func c(err error) {
	if err != nil {
		panic(err)
	}
}
