A lightweight, fast, and colorful command-line task manager built with Go.

## Features
- **Fast & Simple**: Minimalist design for quick task management.
- **Local Storage**: Data is saved in a `tasks.json` file.

## Installation

### Prerequisites
- [Go](https://go.dev/doc/install) 1.21 or higher installed on your system.

### Steps
1. **Clone the repository:**
   ```bash
   git clone https://github.com/brul1ka/todoListGo.git](https://github.com/brul1ka/todoListGo.git
   cd todoListGo
2. **Download dependencies:**
   ```bash
    go mod download
3. **Build the binary:**
   ```bash
   go build -o todo
4. **(Optional) Install system-wide (Linux/macOS):**
   ```bash
   sudo mv todo /usr/local/bin/

## Usage

### Once installed, use the following commands:

| Command | Description | Example |
| :--- | :--- | :--- |
| add | Create a new task | `todo add "Buy coffee"` |
| list | Show all tasks | `todo list` |
| done | Complete a task | `todo done 1` |
| delete | Remove a task | `todo delete 2` |
| help | Show help menu | `todo help` |
