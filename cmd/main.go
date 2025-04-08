package main

import (
	"github.com/xumoyun005/todo-app"
	"github.com/xumoyun005/todo-app/pkg/handlers"
	"log"
)

func main() {
	handlers := new(handlers.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
