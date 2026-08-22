# pgcode

[![CI](https://github.com/b0r1ssh/pgcode/actions/workflows/go-test.yml/badge.svg)](https://github.com/b0r1ssh/pgcode/actions/workflows/go-test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/b0r1ssh/pgcode.svg)](https://pkg.go.dev/github.com/b0r1ssh/pgcode)

A small Go package exposing PostgreSQL SQLSTATE constants as typed string values.

It is useful when you want to check or compare database error codes without manually retyping the SQLSTATE values from PostgreSQL documentation.

## Features

- PostgreSQL SQLSTATE constants for common and advanced error classes
- Easy-to-use string constants
- Compatible with Go modules
- No runtime dependency beyond the Go standard library

## Installation

```bash
go get github.com/b0r1ssh/pgcode
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/b0r1ssh/pgcode"
)

func main() {
    fmt.Println(pgcode.UniqueViolation) // 23505
    fmt.Println(pgcode.SyntaxError)     // 42601
    fmt.Println(pgcode.ConnectionException)
}
```

## Example constants

```go
pgcode.UniqueViolation
pgcode.NotNullViolation
pgcode.ForeignKeyViolation
pgcode.InvalidPassword
pgcode.TimeoutExpired
pgcode.ConnectionException
pgcode.SyntaxError
```

The package includes constants for the major PostgreSQL SQLSTATE classes and many commonly used codes, including:

- Class 00: successful completion
- Class 08: connection exceptions
- Class 22: data exceptions
- Class 23: integrity constraint violations
- Class 42: syntax errors and invalid object references
- and many more PostgreSQL-specific states

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contributing

Contributions are welcome. If you want to improve the list of SQLSTATE values or documentation, feel free to open a pull request.
