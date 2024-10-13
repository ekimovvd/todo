package main

import (
	"log"

	"github.com/ekimovvd/todo-backend"
	"github.com/ekimovvd/todo-backend/pkg/handler"
	"github.com/ekimovvd/todo-backend/pkg/repository"
	"github.com/ekimovvd/todo-backend/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
