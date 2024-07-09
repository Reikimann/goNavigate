/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package db

import (
  "database/sql"
  "log"
  "os"
  "path/filepath"

  "github.com/adrg/xdg"
  "github.com/mattn/go-sqlite3"
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

  stmt, err := tx.Prepare("INSERT INTO directories(path, recurse) VALUES(?,?)")
  if err != nil {
    return err
  }
  defer stmt.Close()

  for _, p := range paths {

    // Gets the absolute file path from input
    p, err = filepath.Abs(p)
    if err != nil {
      log.Fatal(err)
    }

    // Checks if the file exists
    _, err = os.Stat(p)
    if err != nil {
      log.Printf("Path %s doesn't exist.", p)
      continue
    }

    // If the filepath exists, then commit to db.
    _, err = stmt.Exec(p, recurse)
    if err != nil {
      if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
        log.Printf("Path %s already exists in the database.", p)
        continue
      }
      return err
    }
  }

  return nil
}

func ListDirectories() ([]Directory, error) {
  row, err := database.Query("SELECT path, recurse, last_navigation, times_navigated FROM directories")
  if err != nil {
    return nil, err
  }
  defer row.Close()

  var directories []Directory

  for row.Next() {
    var directory Directory

    err := row.Scan(&directory.Path, &directory.Recurse, &directory.LastNavigation, &directory.TimesNavigated)
    if err != nil {
      return nil, err
    }

    directories = append(directories, directory)
  }

  return directories, nil
}
