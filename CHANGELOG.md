# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- CLI package separation for better code organization
- Version command support (`--version`)

### Changed
- Refactored main.go to use dedicated CLI package

## [0.1.0] - 2025-09-20

### Added
- Initial release
- Fast directory scanning 
- Human-readable output with `-h/--human` flag
- Depth limiting with `-d/--depth`
- Verbose debugging with `-v/--verbose`
- Cross-platform support