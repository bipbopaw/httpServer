package main

import (
	httpapi "httpServer/internal/http"
	"httpServer/internal/repository"
	"httpServer/internal/service"
	"log"
	"net/http"
)

func main() {
	repo := repository.NewInMemoryNoteRepository()
	service := service.NewNoteService(repo)
	handler := httpapi.NewHandler(service)
	router := httpapi.NewRouter(handler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Panicln("server started on :8080")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
