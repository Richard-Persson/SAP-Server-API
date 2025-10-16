package db

import (
  "context"
  "fmt"
  "io/fs"
  "os"
  "path/filepath"
  "sort"
  "strings"
  "time"
)

// MigrateFromFiles executes all *.up.sql files in the given directory in lexical order.
// Files should be named with a leading sequence so order is deterministic (e.g. 000001_...).
func MigrateFromFiles(ctx context.Context, dir string) error {
  // gather files
  var files []string
  err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
    if err != nil {
      return err
    }
    if d.IsDir() {
      return nil
    }
    if strings.HasSuffix(d.Name(), ".up.sql") {
      files = append(files, path)
    }
    return nil
  })
  if err != nil {
    return err
  }

  if len(files) == 0 {
    return nil
  }

  sort.Strings(files)

  // execute each file inside its own transaction
  for _, f := range files {
    b, err := os.ReadFile(f)
    if err != nil {
      return fmt.Errorf("read migration %s: %w", f, err)
    }
    sql := string(b)
    // use a context with timeout per migration
    mctx, cancel := context.WithTimeout(ctx, 30*time.Second)
    defer cancel()

    tx, err := DB.Beginx()
    if err != nil {
      return fmt.Errorf("begin tx for %s: %w", f, err)
    }

    // split by semicolon to guard simple multi-statement files; prefer sending whole script when supported
    // attempt to execute whole script first
    if _, err := tx.ExecContext(mctx, sql); err != nil {
      // fallback: try splitting statements
      stmts := splitSQLStatements(sql)
      for _, s := range stmts {
        if strings.TrimSpace(s) == "" {
          continue
        }
        if _, err := tx.ExecContext(mctx, s); err != nil {
          tx.Rollback()
          return fmt.Errorf("exec statement in %s: %w", f, err)
        }
      }
    }

    if err := tx.Commit(); err != nil {
      tx.Rollback()
      return fmt.Errorf("commit migration %s: %w", f, err)
    }
  }

  return nil
}

func splitSQLStatements(sql string) []string {
  parts := strings.Split(sql, ";")
  out := make([]string, 0, len(parts))
  for _, p := range parts {
    trimmed := strings.TrimSpace(p)
    if trimmed != "" {
      out = append(out, trimmed)
    }
  }
  return out
}

