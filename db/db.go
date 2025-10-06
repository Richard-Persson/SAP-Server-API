package db
import (
  "log"
  "os"
  "time"

  "github.com/jmoiron/sqlx"
  _ "github.com/jackc/pgx/v5/stdlib"
  "github.com/joho/godotenv"
)

var DB *sqlx.DB

func Init() {
  _ = godotenv.Load()
  dsn := os.Getenv("DATABASE_URL")
  if dsn == "" {
    log.Fatal("DATABASE_URL not set")
  }

  db, err := sqlx.Connect("pgx", dsn)
  if err != nil {
    log.Fatalf("db connect: %v", err)
  }

  db.SetMaxOpenConns(25)
  db.SetMaxIdleConns(5)
  db.SetConnMaxIdleTime(5 * time.Minute)
  db.SetConnMaxLifetime(30 * time.Minute)

  DB = db
}

