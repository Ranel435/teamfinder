package main

import (
	"teamfinder/backend/internal/pkg/server"
	"teamfinder/backend/internal/pkg/storage"
)

func main() {
	store, err := storage.NewStorage()
	if err != nil {
		panic(err)
	}

	s := server.New(":8090", &store)
	s.Start()
}
