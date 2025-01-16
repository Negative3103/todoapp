package main

import (
	"log"

	"github.com/Negative3103/todoapp"
	"github.com/Negative3103/todoapp/pkg/handlers"
)

func main() {
	handlers := new(handlers.Handler)
	srv := new(todoapp.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %v", err.Error())
	}
}
