# rotzaplog

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue)](https://go.dev/)
[![Test Status](https://github.com/alex-cos/rotzaplog/actions/workflows/test.yml/badge.svg)](https://github.com/alex-cos/rotzaplog/actions/workflows/test.yml)
[![Lint Status](https://github.com/alex-cos/rotzaplog/actions/workflows/lint.yml/badge.svg)](https://github.com/alex-cos/rotzaplog/actions/workflows/lint.yml)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/alex-cos/rotzaplog)](https://goreportcard.com/report/github.com/alex-cos/rotzaplog)

Package Go providing zap logger initialization with daily file rotation.

## Installation

```bash
go get github.com/alex-cos/rotzaplog
```

## Usage

### File Logger

```go
logger := rotzaplog.InitFileLogger(
    "logs/app.log",           // log file path
    zapcore.InfoLevel,        // log level
    false,                   // use UTC time
    true,                    // verbose (also log to console)
)
defer logger.Sync()

logger.Info("message")
```

### Console Logger

```go
logger := rotzaplog.InitConsoleLogger(zapcore.InfoLevel, false)
defer logger.Sync()

logger.Info("message")
```

## Features

- Daily file rotation (files named `app_2026-01-04.log`)
- Optional UTC timestamps
- Verbose mode for console output
- JSON encoder for files, console encoder for terminal
- Caller info and stacktrace on errors
- ISO8601 timestamp format
