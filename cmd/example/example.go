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
