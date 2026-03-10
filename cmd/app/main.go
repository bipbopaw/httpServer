package main

import (
	httpapi "calendar-notes-api/internal/http"
	"calendar-notes-api/internal/notifier"
	"calendar-notes-api/internal/repository"
	"calendar-notes-api/internal/scheduler"
	"calendar-notes-api/internal/service"
	"context"
	"log"
	"net/http"
)

func main() {
	repo := repository.NewInMemoryNoteRepository()
	service := service.NewNoteService(repo)
	handler := httpapi.NewHandler(service)
	router := httpapi.NewRouter(handler)
	consoleNotifier := notifier.NewConsoleNotifier()
	sched := scheduler.NewScheduler(repo, consoleNotifier)

	go sched.Start(context.Background())

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("server started on :8080")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
