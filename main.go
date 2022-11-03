package main

import (
	"github.com/lowryxiao/go-web-fwk/service"
	"net/http"
)

func main() {
	server := &http.Server{
		Handler: service.NewService(),
		Addr:    ":8080",
	}

	server.ListenAndServe()
}
