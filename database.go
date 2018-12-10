package main

import (
	"database/sql"
	"fmt"
	"os/exec"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func createDb(dbfile string, epubs []Epub) {
	start := time.Now()
	db, err := sql.Open("sqlite3", dbfile)
	c(err)
	if db == nil {
		panic("db nil")
	}
	defer db.Close()
	c(err)
	_, err = db.Exec(schema)
	c(err)
	_, err = db.Exec(fmt.Sprintf("INSERT INTO meta VALUES('%s', '%s');",
		time.Now().UTC().Format("2006-01-02 15:04:05"),
		VERSION))
	c(err)
	tx, err := db.Begin()
	c(err)
	stmt, err := tx.Prepare(insert)
	c(err)
	defer stmt.Close()
	for _, epub := range epubs {
		_, err := stmt.Exec(epub.FileName,
			epub.Directory,
			epub.Title,
			epub.Author,
			epub.Language,
			epub.Publisher,
			epub.Description, "")
		if err != nil {
			fmt.Println("THIS EPUB IS NOT VALID FOR DB:", epub)
			tx.Rollback()
		}
	}
	tx.Commit()
	fmt.Printf("Database written to %q in %s\n", dbfile, time.Since(start))
}

func addText(dbfile string, epubs []Epub) {
	db, err := sql.Open("sqlite3", dbfile)
	c(err)
	if db == nil {
		panic("db nil")
	}
	defer db.Close()
	tx, err := db.Begin()
	c(err)
	stmt, err := tx.Prepare("UPDATE epubs SET TextContent = ? WHERE Directory = ? AND FileName = ?;")
	c(err)
	defer stmt.Close()
	for _, epub := range epubs {
		path := epub.Directory + "/" + epub.FileName
		fmt.Printf("Processing text of %q\n", path)
		output, err := exec.Command("epub2txt", path).Output()
		if err != nil {
			fmt.Println("Skipping", path)
			continue
		}
		text := string(output)
		fmt.Println(len(text))
		_, err = stmt.Exec(text, epub.Directory, epub.FileName)
		c(err)
	}
	tx.Commit()
}

var schema = `
CREATE TABLE IF NOT EXISTS epubs(
	FileName    TEXT NOT NULL,
	Directory   TEXT NOT NULL,
	Title       TEXT NOT NULL,
	Author      TEXT NOT NULL,
	Language    TEXT NOT NULL,
	Publisher   TEXT NOT NULL,
	Description TEXT NOT NULL,
	TextContent TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS meta(
	Timestamp   TEXT NOT NULL,
	Version     TEXT NOT NULL
);
`
var insert = `INSERT INTO epubs(
	FileName,
	Directory,
	Title,
	Author,
	Language,
	Publisher,
	Description,
	TextContent
) VALUES (?, ?, ?, ?, ?, ?, ?, ?);
`
