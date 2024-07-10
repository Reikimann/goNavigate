/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package db

import (
  "database/sql"
  "log"
  "os"
  "fmt"
  // "time"
  "path/filepath"

  "github.com/adrg/xdg"
  "github.com/mattn/go-sqlite3"
  _ "github.com/mattn/go-sqlite3"
)

type Directory struct {
  ID             uint
  Path           string
  Recurse        bool
  LastNavigation uint64
  // LastNavigation time.Time
  TimesNavigated uint
}

type dirDB struct {
  Database *sql.DB
  DataDir  string
}

func setupDataPath() string {
  // TODO: Add an environment variable to load other database locations
  // or use a config file (use also viper?)
  // This shall be prioritized over xdg.DataHome

  dataDir := filepath.Join(xdg.DataHome, "goNavigate")
  err := os.MkdirAll(dataDir, os.ModePerm)
  if err != nil {
    log.Fatal(err)
  }

  return dataDir
}

func OpenDatabase() (*dirDB, error) {
  dataPath := setupDataPath()

  database, err := sql.Open("sqlite3", filepath.Join(dataPath, "goNavigate.db"))
  if err != nil {
    return nil, err
  }
  db := dirDB{database, dataPath}

  return &db, nil
}


func (d *dirDB) InitDB() {
  initStatement := `
  CREATE TABLE IF NOT EXISTS directories (
    id INTEGER PRIMARY KEY,
    path TEXT UNIQUE NOT NULL,
    recurse BOOLEAN DEFAULT 0,
    last_navigation INTEGER DEFAULT 0,
    times_navigated INTEGER DEFAULT 0
  );`

  _, err := d.Database.Exec(initStatement)
  if err != nil {
    log.Fatalf("%q: %s\n", err, initStatement)
  }

  log.Println("Directory database created!")
}


func (d *dirDB) AddDirectories(paths []string, recurse bool) error {
  tx, err := d.Database.Begin()
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

  var pathsAdded uint
  for _, p := range paths {
    _, err = stmt.Exec(p, recurse)
    if err != nil {
      if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
        log.Printf("Path %s already exists in the database.", p)
        continue
      }
      return err
    }
    pathsAdded++
  }

  // User feedback detailing the amount of added paths
  switch pathsAdded {
  case 0:
    fmt.Println("No directories were added.")
  case 1:
    fmt.Println("One directory added successfully.")
  default:
    fmt.Println("Directories added successfully.")
  }

  return nil
}

func (d *dirDB) ListDirectories() ([]Directory, error) {
  row, err := d.Database.Query("SELECT * FROM directories")
  if err != nil {
    return nil, err
  }
  defer row.Close()

  var directories []Directory

  for row.Next() {
    var directory Directory

    err := row.Scan(&directory.ID, &directory.Path, &directory.Recurse, &directory.LastNavigation, &directory.TimesNavigated)
    if err != nil {
      return nil, err
    }

    directories = append(directories, directory)
  }

  return directories, nil
}
