<div align="center">

# goNavigate (WIP)

goNavigate is a tool for **quickly** navigating specified directories

It is a CLI application, that allows a user to add directories
to a list and quickly navigate them using fuzzy search.

[Getting Started](#getting-started) •
[Installation](#installation) •
[Configuration](#configuration)

</div>

## Getting started

## Installation
GoNavigate is installed in just 2 steps:

1. **Compilation**
This program needs `gcc` to be built. It also the `CGO_ENABLED=1` environment variable if not set by default.

The following command will install the program binary into to your `GOPATH/bin` or `GOBIN` directory.
```sh
go install github.com/Reikimann/goNavigate@latest
```

2. **Setup goNavigate in your shell**
To start using goNavigate, add it to your shell.
<details>
<summary>Zsh</summary>

> Add this to the <ins>**end**<ins> of your config file (usually `~/.zshrc` or `~/.config/zsh/.zshrc`):
> ```sh
> eval $(goNavigate init zsh)
> ```

</details>

## Configuration

### Flags

When calling the command `goNavigate init`, the following flags are available:

- `--cmd`
    - Changes the prefix of the `g` command.
    - `--cmd j` will change the command to `j`.
- `--no-cmd`
    - Prevents goNavigate from defining the `g` command.
    - The function `__goNavigate_g` will still be available in your shell as, should you choose to alias or keybind it.
