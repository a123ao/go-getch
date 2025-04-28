package main

import (
	"fmt"

	"github.com/a123ao/go-getch/getch"
)

func main() {
	fmt.Println("Press `Enter` to quit")
	for {
		c, _ := getch.Getch()
		if c.VirtualKeyCode == getch.VK_ENTER {
			break
		}

		fmt.Println("Key pressed:", c.UnicodeChar, c.VirtualKeyCode)
	}
}
