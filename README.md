# Go Tic-Tac-Toe

This is a Tic-Tac-Toe game written in Go that lets you challenge an unbeatable AI.

## Installation

### Prerequisites

Verify that Go is installed on your system (version 1.25.5 or newer recommended):

```bash
go version
```

If Go isn't installed, choose one of these methods:

**Option 1: Using Homebrew (macOS/Linux)**

```bash
brew install go
```

**Option 2: Download from go.dev**

Visit [go.dev/dl](https://go.dev/dl) and follow the installation instructions for your operating system.

### Get the Game

Clone the repository:

```bash
git clone https://github.com/nathandickinson32/go-tic-tac-toe.git
cd go-tic-tac-toe
go mod download
```

## Usage

To start playing:

```bash
go run .
```

### How to Play

1. **Pick your players**: Select whether X and O are controlled by humans or AI
2. **Decide who starts**: Choose which player makes the first move
3. **Take your turn**: Enter a number from 1-9 to place your mark
4. **Rematch?**: When the game ends, you can start a new round or quit

## Running Tests

Execute all tests:

```bash
go test ./... -v
```

Run tests with coverage report:

```bash
go test -cover ./...
```

Test a specific file:

```bash
go test -v board_test.go board.go
```