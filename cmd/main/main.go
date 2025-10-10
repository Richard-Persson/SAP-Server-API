package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/Richard-Persson/SAP-Server-API/db"
	"github.com/Richard-Persson/SAP-Server-API/internal/handlers"
)

func main() {
  db.Init()


	/*
	// Dropper tabeller 
	if err := db.RollbackFromFiles(ctx, "migrations"); err != nil {
		log.Fatalf("Rollback failed: %v", err)
	}

	//Lager tabeller
	if err := db.MigrateFromFiles(ctx, "migrations"); err != nil {
		log.Fatalf("migrate failed: %v", err)
	}
	*/



  server := &http.Server{
    Addr:    ":8080",
    Handler: handlers.Router(),
  }

  go func() {
    log.Println("starting server on :8080")
    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      log.Fatalf("server error: %v", err)
    }
  }()

  quit := make(chan os.Signal, 1)
  signal.Notify(quit, os.Interrupt)
  <-quit

}

