package main

import (
	"fmt"

	hook "github.com/robotn/gohook"
)

func main() {
	fmt.Println("Starting to listen for keystrokes...")
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		if ev.Kind == hook.KeyDown {
			keyChar := string(ev.Keychar)

			switch ev.Keychar {
			case 32: // Space
				keyChar = "Space"
			case 13: // Enter
				keyChar = "Enter"
			case 27: // Escape
				keyChar = "Escape"
				// Add more cases as needed for other special keys
			}

			fmt.Printf("Key pressed: %s\n", keyChar)
		}
	}
}
