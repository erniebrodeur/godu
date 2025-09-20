# godu

A fast, modern disk usage analyzer written in Go that mimics the behavior of the traditional Unix `du` command.

## Features

- **Fast directory scanning** with concurrent file processing
- **Human-readable output** with `-h/--human` flag (B, K, M, G, T, P, E)
- **Depth limiting** with `-d/--depth` to control how deep to scan
- **Verbose debugging** with `-v/--verbose` to see detailed scan progress
- **Cross-platform** support (Linux, macOS, Windows)
- **Memory efficient** for large directory trees
- **Permission-aware** - gracefully handles unreadable directories

## Installation

### From Source

```bash
go install github.com/erniebrodeur/godu/cmd/godu@latest
```

### Clone and Build

```bash
git clone https://github.com/erniebrodeur/godu.git
cd godu
go build -o godu cmd/godu/main.go
```

## Usage

```bash
godu [OPTIONS] [DIRECTORY]
```

### Options

- `-h, --human`: Display sizes in human readable format (e.g., 1.2K, 3.4M)
- `-d, --depth N`: Limit output to directories at most N levels deep
- `-v, --verbose`: Enable verbose output showing all files and directories being processed

### Examples

```bash
# Scan current directory
godu

# Scan specific directory with human-readable sizes
godu -h /var/log

# Limit depth to 2 levels
godu -d 2 ~/Projects

# Verbose output for debugging
godu -v -h /tmp

# Combine flags
godu --depth 1 --human --verbose /usr/local
```

## Output Format

The output format matches traditional `du`:

```
4.0K    ./cmd
4.0K    ./internal
124K    ./.git
136K    .
```

- Subdirectories are listed first in alphabetical order
- The root directory total appears last
- Sizes are displayed in KB by default, or human-readable format with `-h`

## How It Works

`godu` walks the directory tree and calculates disk usage by:

1. **Block-aligned calculation**: File sizes are rounded up to the nearest 4KB block to match actual disk usage
2. **Accumulative sizing**: Each file's size is added to all parent directories up to the scan root
3. **Permission handling**: Unreadable directories are silently skipped rather than causing errors
4. **Depth filtering**: Only directories within the specified depth limit are included in output

## Performance

`godu` is designed to be fast and memory-efficient:

- Single-pass directory traversal
- Minimal memory allocation
- Efficient path handling
- Graceful error recovery

## Compatibility

`godu` aims to be compatible with standard Unix `du` behavior:

- Same output format and ordering
- Similar size calculations (block-aligned)
- Comparable command-line interface
- Cross-platform support

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -am 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by the classic Unix `du` command
- Built with Go's excellent `filepath` package
- Thanks to the Go community for excellent tooling and documentation