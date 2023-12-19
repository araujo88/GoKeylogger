# GoKeylogger

This repository contains a simple Go program that listens for keystrokes and prints them to the console. It uses the `gohook` library to capture keyboard events.

## Overview

The program is a basic demonstration of using the `gohook` library to capture and handle keyboard events in Go. When a key is pressed, the program prints a human-readable representation of the key to the console. It specifically identifies some common special keys like Space, Enter, and Escape.

## Getting Started

### Prerequisites

- Go (Version 1.13 or later)
- Dependencies: `gohook` library
- Libraries: libx11-dev libxcb1-dev libx11-xcb-dev libxkbcommon-x11-dev libxtst-dev

### Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/araujo88/GoKeylogger.git
   cd GoKeylogger
   ```

2. **Install Dependencies:**

   ```bash
   go get -u github.com/robotn/gohook
   ```

   ```bash
   sudo apt-get update
   sudo apt-get install libx11-dev libxcb1-dev libx11-xcb-dev libxkbcommon-x11-dev libxtst-dev
   ```

### Running the Program

Execute the program by running:

```bash
go run main.go
```

The program will start and begin listening for keystrokes. Press various keys to see their representation printed to the console.

## Usage

The program is straightforward to use:

- Run the program.
- Once running, it will listen for and output any key pressed.
- Special keys (Space, Enter, Escape) are identified explicitly.

## Customization

You can extend the switch-case in the `main.go` file to handle more special keys or add additional functionality as per your requirements.

## Contributing

Contributions to this project are welcome. Please feel free to fork the repository, make changes, and submit pull requests.

## License

This project is licensed under the GPL License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to the creators and maintainers of the `gohook` library.
