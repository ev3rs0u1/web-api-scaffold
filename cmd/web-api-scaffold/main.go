package main

import (
	"web-api-scaffold/cmd/server"
	"web-api-scaffold/internal/pkg/config"
)

func init() {
	config.LoadFile()
}

func main() {
	server.New().
		SetDatabase().
		SetMiddleware().
		RegisterRoutes().
		ListenAndServe()
}
