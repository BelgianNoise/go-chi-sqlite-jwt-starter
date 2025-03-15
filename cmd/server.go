package main

import (
	"go-chi-sqlite-jwt-starter/config"
	"go-chi-sqlite-jwt-starter/internal/database"
	"go-chi-sqlite-jwt-starter/internal/provider"
	"go-chi-sqlite-jwt-starter/internal/server"
	"log"
	"net/http"
)

func main() {
	config.InitializeVariables()
	database.Initialize()
	provider.Initialize()
	r := server.Initialize()

	port := getPort()
	log.Printf("Server running on port %v", port)
	http.ListenAndServe(":"+getPort(), r)
}

func getPort() string {
	return config.Variables.Port
}
