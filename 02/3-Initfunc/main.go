// init initializes the database connection and sets up the default HTTP handlers.
//
// The database connection is initialized using the DB_DNS environment variable.
// If there is an error connecting to the database, the function will panic.
//
// The HTTP handlers are set up to redirect requests to the root path ("/") for
// the "/blog" and "/blog/" paths, and to serve static files for the "/favicon.ico",
// "/fonts.css", "/fonts/", and "/lib/godoc/" paths.
package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
)

var DB *sql.DB

// Bad init function, won't allow users to handle errors

func init() {
	dns := os.Getenv("DB_DNS")
	d, err := sql.Open("mysql", dns)

	if err != nil {
		log.Panic(err)
	}

	err = d.Ping()
	if err != nil {
		log.Panic(err)
	}
	DB = d
}

// Good init function, won't panic, just setup the default HTTP handlers

func init() {
	redirect := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	http.HandleFunc("/blog", redirect)
	http.HandleFunc("/blog/", redirect)

	static := http.FileServer(http.Dir("static"))
	http.Handle("/favicon.ico", static)
	http.Handle("/fonts.css", static)
	http.Handle("/fonts/", static)
	http.Handle("/lib/godoc/", http.StripPrefix("/lib/godoc", static))

}

func main() {

}
