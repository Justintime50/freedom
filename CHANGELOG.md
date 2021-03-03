# CHANGELOG

## v0.4.2 (2021-03-03)

* Bumping release to add support for `ARM macOS`
* Adds support for `Go 1.16`

## v0.4.1 (2021-01-16)

* Removed testing framework and replaced with `mockcmd`

## v0.4.0 (2021-01-15)

* Replaces `fmt.Printf` with `fmt.Println`
* Added `os.Exit(1)` when no flag is passed after the error message
* Improved error handling by returning status codes and messages where applicable
* Full unit tests and code coverage

## v0.3.0 (2020-12-30)

* Changes the executable `free` to `freedom` since `free` is already a tool built into Linux

## v0.2.0 (2020-12-30)

* Refactored the `ports` module to properly accept an integer as an argument
* Tool now exits instead of returning on errors
* Added golangci-lint to GitHub Actions
* Added auto-releasing via GoReleaser and GitHub Actions (closes #2)

## v0.1.0 (2020-12-29)

* Initial release
* Prune your Docker instance
* Kill a process on a port
* Close all Finder windows in macOS
