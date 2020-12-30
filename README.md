<div align="center">

# Freedom

Give your computer the freedom to do more by freeing its resources.

[![build](https://github.com/Justintime50/freedom/workflows/build/badge.svg)](https://github.com/Justintime50/freedom/actions)
[![Licence](https://img.shields.io/github/license/justintime50/GitHub-archive)](LICENSE)

<img src="assets/showcase.png" alt="Showcase">

Freedom lets you free different resources of your machine so it can do more. Prune Docker, close all Finder windows on macOS, kill processes attached to ports, and more.

**Note:** This package is built for macOS and Linux.

</div>

## Install

```bash
# Setup the tap
brew tap justintime50/formulas

# Install Freedom
brew install freedom
```

## Usage

```
Usage:
    free --port 8000

Options:
    -docker
        Prune your Docker instance
    -finder
        Close all your macOS Finder Windows.
    -port
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
```

## Attribution

* Based on [nuke](https://github.com/Matt-Gleich/nuke)
* Icons made by <a href="https://www.flaticon.com/authors/freepik" title="Freepik">Freepik</a> from <a href="https://www.flaticon.com/" title="Flaticon">www.flaticon.com</a>
