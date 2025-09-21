# godu

[![Go Report Card](https://goreportcard.com/badge/github.com/erniebrodeur/godu)](https://goreportcard.com/report/github.com/erniebrodeur/godu)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Fast disk usage analyzer for Go. A modern replacement for `du` with concurrent scanning and cross-platform support.

## Installation

```bash
go install github.com/erniebrodeur/godu/cmd/godu@latest
```

Or build from source:

```bash
git clone https://github.com/erniebrodeur/godu.git
cd godu && go build ./cmd/godu
```

## Usage

```bash
godu [flags] [directory]
```

**Flags:**
- `-h, --human` - Human readable sizes (1.2K, 3.4M)  
- `-d, --depth N` - Limit depth to N levels
- `-v, --verbose` - Debug output
- `--version` - Show version

**Examples:**

```bash
godu                    # current directory  
godu /var/log           # specific path
godu -h -d 2 ~/src      # human readable, max 2 levels deep
```

## Output

Matches traditional `du` format:

```
4.0K    ./cmd
8.0K    ./internal  
136K    .
```

## Performance

- Concurrent file processing
- Block-aligned size calculation (4KB blocks)
- Memory efficient for large trees
- Graceful permission error handling

## License

MIT