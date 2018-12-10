/*
 * TO DO:
 *
 * -- Add possibility to insert or update one row only in DB (not recreating the whole db everytime there is a little change)
 * -- Improve performance of insert full text (-text option) with goroutines
 */

package main

/* Globals */
const (
	VERSION  = "0.2.1"
	PROGNAME = "collix"
)

func main() {
	configuration, command := setup()
	switch command {
	case "init":
		initialize(configuration)
	case "info":
		info(configuration)
	default:
		usage()
	}
}
