# WSL Snappy Backup

![Version](https://img.shields.io/badge/version-1.0.0-green.svg)

WSL Snappy Backup is a utility tool for backing up Windows Subsystem for Linux (WSL) distributions. It supports both sequential and parallel backup modes, providing flexibility and efficiency.

## Features

- Backup WSL distributions to a specified directory.
- Run backups either sequentially or in parallel ("rocket mode") for increased speed.
- Environment variable configuration for easy setup.
- Notification alerts on backup completion or failure.

## Prerequisites

Before you start using WSL Snappy Backup, make sure you have:

- Windows 10 or higher with WSL installed.
- Go programming language environment.
- WSL distributions that you want to back up.

## Installation

1. Clone the repository or download the source code.
    
2. Navigate to the source code directory.
    
3. Compile the program:
    
    bash
    

1. `GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui" -o bin/WSLSnappy.exe cmd/main.go`
    

## Configuration

Set the following environment variables according to your system setup:

- `WSL_BACKUP_DIR`: The directory where the WSL distributions will be backed up.
- `WSL_DISTRIBUTIONS`: The list of WSL distribution names to back up, separated by commas (e.g., `Ubuntu,Debian`).

You can set these variables using the Windows Command Prompt:

```bash

`setx WSL_BACKUP_DIR "C:\path\to\your\backup\directory" setx WSL_DISTRIBUTIONS "Ubuntu,Debian"`
```

## Usage

Run the compiled executable from the command line:

- To perform a sequential backup:
    
```bash
    

- `.\bin\WSLSnappy.exe`
```

- To perform backups in parallel (rocket mode):
    
```bash
    

- `.\bin\WSLSnappy.exe --rocket`
    
```

The program will start the backup process based on your configuration and display the time taken to complete the backups.

## Notifications

Upon completion or failure of a backup, a desktop notification will be shown, informing you of the status.

## Contributing

Contributions, issues, and feature requests are welcome. 
Please feel free to submit issues and pull requests.

## Buy me a coffee

Whether you use this project, have learned something from it, or just like it, please consider supporting it by buying me a coffee, so I can dedicate more time on open-source projects like this 

<a href="https://www.buymeacoffee.com/andrearapoA" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>

## License

This project is licensed under the [MIT License](https://github.com/andrearaponi/wslSnappy/blob/main/LICENSE).