package main

import (
	"log"

	"github.com/ekimovvd/todo"
	"github.com/ekimovvd/todo/pkg/handler"
)

func main() {
	srv := new(todo.Server)
	handlers := new(handler.Handler)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
