package main

import (
	"GourseAPI/internal/platform/server"
	"log"
)

const (
	host = "localhost"
	port = 8080
)

func main() {
	srv := server.New(host, port)
	if err := srv.Run(); err != nil {
		log.Fatalf("could not run the server: %v", err)
	}
}
