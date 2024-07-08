package db

import (
  "database/sql"
  "log"
  "os"
  "path/filepath"

  "github.com/adrg/xdg"
  _ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

type Directory struct {
  Path string
  Recurse bool
  LastNavigation int64
  TimesNavigated int
}

func OpenDatabase() error {
  // TODO: Add an environment variable to load other database locations
  // or use a config file (use also viper?)
  // This shall be prioritized over xdg.DataHome

  dataDir := filepath.Join(xdg.DataHome, "goNavigate")
  err := os.MkdirAll(dataDir, os.ModePerm)
  if err != nil {
    log.Fatal(err)
  }

  dbPath := filepath.Join(dataDir, "goNavigate.db")

  database, err = sql.Open("sqlite3", dbPath)
  if err != nil {
    log.Fatal(err)
  }

  return database.Ping()
}


func InitDB() {
  initStatement := `
  CREATE TABLE IF NOT EXISTS directories (
    id INTEGER PRIMARY KEY,
    path TEXT UNIQUE NOT NULL,
    recurse BOOLEAN DEFAULT 0,
    last_navigation INTEGER DEFAULT 0,
    times_navigated INTEGER DEFAULT 0
  );`

  _, err := database.Exec(initStatement)
  if err != nil {
    log.Fatalf("%q: %s\n", err, initStatement)
  }

  log.Println("Directory database created!")
}

func AddDirectories(paths []string, recurse bool) error {
  tx, err := database.Begin()
  if err != nil {
    return err
  }
  // Rolls back db changes if AddDirectories() returns an error,
  // otherwise it commits the transactions.
  defer func() {
    if err != nil {
      tx.Rollback()
      return
    }
    err = tx.Commit()
  }()

  stmt, err := tx.Prepare("INSERT OR IGNORE INTO directories(path, recurse) VALUES(?,?)")
  if err != nil {
    return err
  }
  defer stmt.Close()

  // TODO: What if the recurse isn't passed? will it fail?
  for _, p := range paths {
    _, err = stmt.Exec(p, recurse)
    if err != nil {
      return err
    }
  }

  return nil
}
