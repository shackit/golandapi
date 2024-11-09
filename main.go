package main

import (
	"goland-course-api/internal/db"
	"goland-course-api/internal/todo"
	"goland-course-api/internal/transport"
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
