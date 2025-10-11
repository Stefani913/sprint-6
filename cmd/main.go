package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(log.Writer(), "http: ", log.LstdFlags)
	newServ := server.Router(logger)
	if err := newServ.Server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
