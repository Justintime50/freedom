<div align="center">

# Freedom

Give your computer the freedom to do more by freeing its resources.

[![Build Status](https://github.com/Justintime50/freedom/workflows/build/badge.svg)](https://github.com/Justintime50/freedom/actions)
[![Coverage Status](https://coveralls.io/repos/github/Justintime50/freedom/badge.svg?branch=main)](https://coveralls.io/github/Justintime50/freedom?branch=main)
[![Licence](https://img.shields.io/github/license/justintime50/freedom)](LICENSE)

<img src="assets/showcase.png" alt="Showcase">

</div>

Freedom lets you free different resources of your machine so it can do more. Prune Docker, close all Finder windows on macOS, kill processes attached to ports, and more.

## Install

```bash
# Setup the tap
brew tap justintime50/formulas

# Install Freedom
brew install freedom
```

## Usage

**Note:** Single or double dashes can be used with options.

```
Usage:
    freedom --port 8000

Options:
    -help
        Show this output.
    -docker
        Prune your Docker instance.
    -finder
        Close all your macOS Finder Windows.
    -port int
        Free a port of its process (pass a port number as an argument).
```

## Development

```bash
# Build the project
make build

# Install the project globally from source
make install

# Clean the executables
make clean

# Test the project
make test

## Get test coverage
make coverage

# Lint the project (requires golangci-lint be installed)
make lint
```

## Attribution

* Inspired by [nuke](https://github.com/Matt-Gleich/nuke)
* Icons made by <a href="https://www.flaticon.com/authors/freepik" title="Freepik">Freepik</a> from <a href="https://www.flaticon.com/" title="Flaticon">www.flaticon.com</a>
