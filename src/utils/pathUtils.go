package utils

import (
  "os"
  "fmt"
  "strings"
  "regexp"
)

// Strips the /home/user/ from a given path. Returns the whole path on error.
func DirPathStripHome(path string) (string, error) {
    homeDir, err := os.UserHomeDir()
    if err != nil {
      return path, err
    }
    if strings.Contains(path, homeDir) {
      regexString := fmt.Sprintf("^%s/", homeDir)
      regex := regexp.MustCompile(regexString)
      path = regex.Split(path, -1)[1]
    }
    return path, nil
}
