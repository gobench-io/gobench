package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/web"
	_ "github.com/mattn/go-sqlite3"
)

// run replay -f testresult.sqlite3 -p 3000

func main() {
	file := flag.String("f", "", "SQLite3 file. Required")
	port := flag.Int("p", 8000, "Listening port of the HTTP server")

	flag.Parse()

	if *file == "" {
		flag.PrintDefaults()
		os.Exit(-1)
	}

	db, err := ent.Open("sqlite3", *file)

	if err != nil {
		log.Panicf("failed opening connection to %s: %v\n", *file, err)
	}

	r := web.New(db)

	portS := fmt.Sprintf(":%d", *port)
	log.Printf("Serving http at %s\n", portS)

	if err := http.ListenAndServe(portS, r); err != nil {
		log.Panicf("failed to start http server at port %s: %v", portS, err)
	}
}
