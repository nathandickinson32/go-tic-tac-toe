# Go Tic-Tac-Toe

A command-line Tic-Tac-Toe game written in Go with three game modes and an unbeatable AI opponent using the minimax algorithm.

## Features

- **Three Game Modes:**
  - Human vs Human
  - Human vs AI
  - AI vs AI

- **Unbeatable AI:** Implements the minimax algorithm with depth-based scoring for optimal play

## Installation

### Prerequisites

- Go 1.25.5 or higher

### Setup

```bash
# Clone the repository
git clone https://github.com/nathandickinson32/go-tic-tac-toe.git
cd go-tic-tac-toe

# Run the game
go run .
```

## How to Play

1. Run the game and select a mode (1-3)
2. Players alternate turns
3. Enter a number (1-9) to place your mark on the board:

```
 1 | 2 | 3 
-----------
 4 | 5 | 6 
-----------
 7 | 8 | 9 
```

4. First player to get three in a row (horizontally, vertically, or diagonally) wins
5. If all spaces are filled with no winner, the game is a draw


### Running Tests


```bash
# Run all tests
go test -v

# Run tests with coverage
go test -cover .

# Run tests excluding long-running simulations
go test -short .

# Run specific test
go test -run TestAIPlayer_Simulation
```


