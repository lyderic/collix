package main

import "fmt"

type Configuration struct {
	Database  string `json:"dbfile"`
	Directory string `json:"basedir"`
}

func (configuration Configuration) String() string {
	return fmt.Sprintf("Database:  %s\nDirectory: %s\n", configuration.Database, configuration.Directory)
}
