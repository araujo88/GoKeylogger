package main

import (
	"fmt"
	"log"
	"net"

	hook "github.com/robotn/gohook"
)

func main() {
	fmt.Println("Starting to listen for keystrokes...")
	evChan := hook.Start()
	//defer hook.End()

	udpAddress, err := net.ResolveUDPAddr("udp", "127.0.0.1:6969")
	if err != nil {
		log.Panic(err)
	}

	connection, err := net.DialUDP("udp", nil, udpAddress)
	if err != nil {
		log.Panic(err)
	}
	defer connection.Close()

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
			connection.Write([]byte(keyChar))
		}
	}
}
