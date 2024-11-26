package main

import (
	"teamfinder/backend/internal/pkg/server"
)

func main() {
	s := server.New(":8080", &store)
	s.Start()
}
