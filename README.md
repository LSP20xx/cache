# Cache Module

This Go module provides a generic, thread-safe cache implementation. It is designed to be easily integrated into any Go project, supporting both concurrent and non-concurrent use cases.

## Features

- **Thread-safe caching**: Ensures safe access in concurrent environments, avoiding race conditions.
- **Generic support**: Cache any type of data, making it flexible for various use cases.
- **Easy integration**: Simple API for quick integration into existing projects.
- **Configurable concurrency**: Optionally configure the cache to be thread-safe depending on your application's needs.

## Installation

To install this module, you can run:

```bash
go get github.com/LSP20xx/cache
```

## Usage

### Basic Usage

To create and use a simple cache:

```go
package main

import (
    "github.com/LSP20xx/cache"
)

func main() {
    c := cache.New(false) // Create a non-concurrent cache

    c.Set("key", "value")
    value, found := c.Get("key")
    if found {
        fmt.Println("Cached value:", value)
    }
}
```

### Thread-safe Cache

To create a thread-safe cache:

```go
package main

import (
    "github.com/LSP20xx/cache"
)

func main() {
    c := cache.New(true) // Create a concurrent cache

    c.Set("key", "value")
    value, found := c.Get("key")
    if found {
        fmt.Println("Cached value:", value)
    }
}
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
