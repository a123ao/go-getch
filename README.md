# Go Getch (Windows)

A simple Go library for reading single characters (including special keys like arrows) from the Windows console without echoing, similar to the C `getch()` function. Uses standard library packages (`golang.org/x/sys/windows`) only.

## Installation

```bash
go get github.com/a123ao/go-getch
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/a123ao/go-getch"
)

func main() {
    fmt.Println("Press `Enter` to quit")
    for {
        key, _ := getch.Read()
        if key.VirtualKeyCode == getch.VK_ENTER {
            break
        }
        fmt.Println("Key pressed:", key.UnicodeChar, key.VirtualKeyCode)
    }
}

```
