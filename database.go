package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func writeDb(dbfile string, epubs []Epub) {
	start := time.Now()
	os.Remove(dbfile)
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
	stmt, _ := tx.Prepare(insert)
	c(err)
	defer stmt.Close()
	for idx, epub := range epubs {
		result, err := stmt.Exec(epub.FileName,
			epub.Directory,
			epub.Title,
			epub.Author,
			epub.Language,
			epub.Publisher,
			epub.Description)
		if err != nil {
			fmt.Println("THIS EPUB IS NOT VALID FOR DB:", epub)
		}
		lastinsert, err := result.LastInsertId()
		c(err)
		fmt.Printf("Insert #%d: row id #%d", idx+1, lastinsert)
		fmt.Printf("\r                                       \r")
	}
	tx.Commit()
	fmt.Printf("Database written to %q in %s\n", dbfile, time.Since(start))
}

var schema = `
CREATE TABLE IF NOT EXISTS epubs(
	FileName    TEXT NOT NULL,
	Directory   TEXT NOT NULL,
	Title       TEXT NOT NULL,
	Author      TEXT NOT NULL,
	Language    TEXT NOT NULL,
	Publisher   TEXT NOT NULL,
	Description TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS meta(
	Timestamp   TEXT NOT NULL,
	Version     TEXT NOT NULL
);
`
var insert = `INSERT OR REPLACE INTO epubs(
	FileName,
	Directory,
	Title,
	Author,
	Language,
	Publisher,
	Description
) VALUES (?, ?, ?, ?, ?, ?, ?);
`
