package main

import (
	"github.com/shackit/golandapi/internal/db"
	"github.com/shackit/golandapi/internal/todo"
	"github.com/shackit/golandapi/internal/transport"
	"log"
)

func main() {

	d, err := db.New("postgres", "example", "localhost", 5432, "postgres")
	if err != nil {
		log.Fatal(err)
	}

	svc := todo.NewService(d)
	server := transport.NewServer(svc)

	err = server.Serve()
	if err != nil {
		return
	}
}
