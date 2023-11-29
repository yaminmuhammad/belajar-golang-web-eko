package belajar_golang_web_eko

import (
	"log"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
