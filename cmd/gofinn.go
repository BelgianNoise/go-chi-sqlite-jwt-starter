package main

import (
	"gofinn/config"
	"gofinn/internal/provider"
	"gofinn/internal/server"
	"log"
	"net/http"
)

func main() {
	config.InitializeVariables()
	provider.Initialize()
	r := server.Initialize()

	port := getPort()
	log.Printf("Server running on port %v", port)
	http.ListenAndServe(":"+getPort(), r)
}

func getPort() string {
	return config.Variables.Port
}
