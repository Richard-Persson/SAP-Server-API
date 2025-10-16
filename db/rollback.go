package db

// RollbackFromFiles executes all *.down.sql files in the given directory in reverse lexical order.
// Use this to drop or rollback schema changes defined in corresponding .down.sql files.
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

func RollbackFromFiles(ctx context.Context, dir string) error {
  var files []string

  err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
    if err != nil {
      return err
    }
    if d.IsDir() {
      return nil
    }
    if strings.HasSuffix(d.Name(), ".down.sql") {
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
  // reverse order so last-up is rolled back first
  for i, j := 0, len(files)-1; i < j; i, j = i+1, j-1 {
    files[i], files[j] = files[j], files[i]
  }

  for _, f := range files {
    b, err := os.ReadFile(f)
    if err != nil {
      return fmt.Errorf("read rollback %s: %w", f, err)
    }
    sql := string(b)

    // use a context with timeout per rollback
    mctx, cancel := context.WithTimeout(ctx, 30*time.Second)

    tx, err := DB.Beginx()
    if err != nil {
      cancel()
      return fmt.Errorf("begin tx for %s: %w", f, err)
    }

    // try executing whole script first
    if _, err := tx.ExecContext(mctx, sql); err != nil {
      // fallback: try splitting statements
      stmts := splitSQLStatements(sql)
      for _, s := range stmts {
        if strings.TrimSpace(s) == "" {
          continue
        }
        if _, err := tx.ExecContext(mctx, s); err != nil {
          tx.Rollback()
          cancel()
          return fmt.Errorf("exec statement in %s: %w", f, err)
        }
      }
    }

    if err := tx.Commit(); err != nil {
      tx.Rollback()
      cancel()
      return fmt.Errorf("commit rollback %s: %w", f, err)
    }

    cancel()
  }
  return nil
}

