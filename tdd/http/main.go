package main

import (
	"log"
	"net/http"

	"github.com/cpustejovsky/ready-steady-go/tdd/http/server"
)

func main() {
	server := &server.PlayerServer{server.NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
