package main

import (
	"fmt"
	"log"
	"net"
)

const maxBufferSize = 1024

func main() {
	udpAddress, err := net.ResolveUDPAddr("udp", "127.0.0.1:6969")
	if err != nil {
		log.Panic(err)
	}

	connection, err := net.ListenUDP("udp", udpAddress)
	if err != nil {
		log.Panic(err)
	}
	defer connection.Close()

	buffer := make([]byte, maxBufferSize)

	for {
		n, _, err := connection.ReadFromUDP(buffer)
		if err != nil {
			log.Panic(err)
		}
		if n > 0 {
			fmt.Printf("Key received: %s\n", string(buffer[:n]))
		}
	}
}
