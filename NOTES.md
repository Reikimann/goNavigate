# Note to self
No need to look at this mess. It's just my noobish notes xD

## TUI rendering
Links regarding TUI rendering:
- https://github.com/charmbracelet/bubbletea/issues/860 (See the last to comments to fix coloring with lipgloss and termenv)
- https://github.com/junegunn/fzf/issues/3741
- https://github.com/junegunn/fzf/wiki/Language-bindings

Does the method below mean that I have to use tea.WithInput(os.OpenFile("/dev/tty", os.O_RDONLY, 0)) to get user input and at the
same time use STDIN to use pipes?
The junegunn/fzf discussions link have some interesting code featured in a patch.

Maybe use stderr to render UI. fzf now uses the following
- Input list: Read from STDIN
- User input: Read from /dev/tty - read
- UI: Print to STDERR or /dev/tty - write
- Selected item: Print to STDOUT

It seems the link below changed how UI rendering worked. FZF now uses the /dev/tty
https://github.com/junegunn/fzf/discussions/3792#user-content-fn-1-73dfe85d976f2420e77f78eb8d76f128

Old way?
```go
    p := tea.NewProgram(tui.NewModel(), tea.WithOutput(os.Stderr))
```

Im trying this out
```go
    tty, err := os.OpenFile("/dev/tty", os.O_WRONLY, 0)
    if err != nil {
      panic(err)
    }
    defer tty.Close()

    p := tea.NewProgram(tui.NewModel(), tea.WithOutput(tty))
    finalModel, err := p.Run()
    if err != nil {
      fmt.Printf("There has been an error: %v", err)
      os.Exit(1)
    }
```

## Exit codes makes terminal quit
When returning exit codes the terminal closes. Investigate further.
