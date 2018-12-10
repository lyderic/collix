package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func initialize(configuration Configuration) {
	if _, err := os.Stat(configuration.Database); os.IsNotExist(err) {
		epubs, err := index(configuration.Directory)
		c(err)
		createDb(configuration.Database, epubs)
	} else {
		log.Fatalf("Database %q already exists!\n", configuration.Database)
	}
	return
}

func info(configuration Configuration) {
	fmt.Println("Epubs in database:")
	cmd := exec.Command("sqlite3", configuration.Database, "SELECT COUNT(*) FROM epubs;")
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	cmd.Run()
	return
}
