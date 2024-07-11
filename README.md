# goNavigate (WIP)

goNavigate is a CLI application, written in Go, that allows a user to add
directories to a list and quickly navigate them using fuzzy search.

## Compilation
This program needs `gcc` to be built. It also the `CGO_ENABLED=1` environment variable if not set by default.

The following command will install the program binary into to your `GOPATH/bin` or `GOBIN` directory.
```sh
go install github.com/Reikimann/goNavigate
```
