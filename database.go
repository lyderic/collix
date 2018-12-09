package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func writeDb(dbfile string, epubs []Epub) {
	db, err := sql.Open("sqlite3", dbfile)
	c(err)
	if db == nil {
		panic("db nil")
	}
	defer db.Close()
	_, err = db.Exec(create)
	c(err)
	stmt, _ := db.Prepare(insert)
	c(err)
	defer stmt.Close()
	for _, epub := range epubs {
		_, err = stmt.Exec(epub.FileName, epub.Directory, epub.Title, epub.Author, epub.Language, epub.Publisher, epub.Description)
		c(err)
	}
	fmt.Println("Database written to", dbfile)
}

var create = `CREATE TABLE IF NOT EXISTS epubs(
	FileName    TEXT NOT NULL,
	Directory   TEXT NOT NULL,
	Title       TEXT NOT NULL,
	Author      TEXT NOT NULL,
	Language    TEXT NOT NULL,
	Publisher   TEXT,
	Description TEXT
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
