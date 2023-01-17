package main

import (
	"log"
	"net/http"
	"time"
	"toy-siwe/poc"
)

func main() {
	srv := poc.WebServerStartUp()

	// Attach SIWE functionality
	srv.WebServerSIWE()

	concept := &http.Server{
		Addr:           "127.0.0.1:3000",
		Handler:        srv.Conn,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(concept.ListenAndServe())
}
