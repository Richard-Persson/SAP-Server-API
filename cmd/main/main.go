
package main

import (
  "context"
  "log"
  "net/http"
  "os"
  "os/signal"
  "time"

  "github.com/Richard-Persson/SAP-Server-API/db"
  "github.com/Richard-Persson/SAP-Server-API/internal/handlers"
)

func main() {
  db.Init()

	ctx := context.Background()

	// Dropper tabeller 
	if err := db.RollbackFromFiles(ctx, "migrations"); err != nil {
		log.Fatalf("Rollback failed: %v", err)
	}




	//Lager tabeller
	if err := db.MigrateFromFiles(ctx, "migrations"); err != nil {
		log.Fatalf("migrate failed: %v", err)
	}



  srv := &http.Server{
    Addr:    ":8080",
    Handler: handlers.Router(),
  }

  go func() {
    log.Println("starting server on :8080")
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      log.Fatalf("server error: %v", err)
    }
  }()

  quit := make(chan os.Signal, 1)
  signal.Notify(quit, os.Interrupt)
  <-quit

  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  if err := srv.Shutdown(ctx); err != nil {
    log.Fatalf("shutdown error: %v", err)
  }
  if err := db.DB.Close(); err != nil {
    log.Printf("db close: %v", err)
  }
}

