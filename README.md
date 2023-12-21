# GoKeylogger

## Overview

This repository contains two Go programs designed to work together for capturing and transmitting keystroke data over a UDP network. The first program (`server.go`) is a UDP server that listens for incoming keystroke data, and the second program (`client.go`) captures keystrokes and sends them to the UDP server.

## Files in the Repository

1. **server.go**: A UDP server program that listens on port 6969 and prints out any received keystroke data.
2. **client.go**: A program that captures keystrokes and sends them to the UDP server running on `127.0.0.1:6969`.

## Requirements

- Go programming language
- `gohook` package (github.com/robotn/gohook) for capturing keystrokes

## Installation

1. Clone the repository.
2. Ensure Go is installed on your system.
3. Install the `gohook` package using `go get github.com/robotn/gohook`.

## Running the Programs

1. **Start the server**: Run `go run server.go` to start the UDP server.
2. **Start the client**: In a separate terminal, run `go run client.go` to start capturing keystrokes.

## Usage

- The **client** program will listen for incoming UDP packets on port 6969. It will print out the received keystroke data to the console.
- The **server** program captures keystrokes and sends the corresponding data to the UDP client. It prints the captured keys to the console and sends this data over the network.
- Special keys like Space, Enter, and Escape are handled and sent as text (e.g., "Space", "Enter", "Escape").

## Notes

- Ensure that both programs are running on the same network.
- For demonstration purposes, the server is set to `127.0.0.1`, but you can modify this to suit your network configuration.
- The programs are designed for educational and demonstration purposes. Be mindful of privacy and security implications when capturing and transmitting keystroke data.

## Contributing

Feel free to fork the repository, make changes, and submit pull requests. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contributing

Contributions to this project are welcome. Please feel free to fork the repository, make changes, and submit pull requests.

## License

This project is licensed under the GPL License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to the creators and maintainers of the `gohook` library.
