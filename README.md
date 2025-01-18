# HLCP - Highlight Copier

HLCP (Highlight Copier) is a Windows application that automatically copies highlighted text from any application (e.g., text editors, browsers) into the clipboard. It runs in the background with a system tray icon and can be terminated easily.

## Features
- Automatically copies selected text to the clipboard without requiring manual `Ctrl+C`.
- Runs silently in the system tray.
- Simple and lightweight.
- Easy termination via the system tray menu.

## Requirements
- **Operating System:** Windows (10/11 recommended)
- **Go Version:** 1.20+ (for building the application)

## Installation
### Precompiled Executable
1. Download the precompiled `HLCP.exe` file from the [Releases](#) section.
2. Place the executable in your desired directory.
3. Run the program to start monitoring clipboard activity.

### Build from Source
1. Clone the repository:
   ```bash
   git clone https://github.com/p2zbar/HLCP.git
   cd HLCP
   ```
2. Initialize the Go module (if not already initialized):
   ```bash
   go mod init 
   go mod tidy
   ```
3. Install required dependencies:
   ```bash
   go get github.com/atotto/clipboard
   go get github.com/getlantern/systray
   ```
4. Ensure you have Go installed. Then, build the executable:
   ```bash
   go build -ldflags="-H windowsgui" -o HLCP.exe main.go
   ```
5. Run the `HLCP.exe` file to start the program.

## Usage
1. Start the program by running `HLCP.exe`.
2. Highlight any text in a supported application (e.g., text editor, browser).
3. The text will be automatically copied to the clipboard.
4. Access the system tray icon for options:
   - **Quit:** Exits the application and removes the icon from the tray.

## System Tray Icon
The application displays an icon in the system tray while running. The icon provides quick access to the "Quit" option.

## How It Works
HLCP uses Windows API calls to:
1. Monitor the system clipboard.
2. Simulate the `Ctrl+C` command when text is selected, ensuring it is copied to the clipboard.


## License
This project is licensed under the https://creativecommons.org/licenses/by-nc/4.0.

You can share,fork,re-use but not for commercial use.


