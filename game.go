package main

import (
	"bufio"
	"io"
)

type GameStatus int

const (
	InProgress GameStatus = iota
	XWins
	OWins
	Draw
)

type Game struct {
	board     Board
	bufReader *bufio.Reader
	bufWriter *bufio.Writer
}

func newGame(reader io.Reader, writer io.Writer) *Game {
	return &Game{
		board:     initBoard(),
		bufReader: bufio.NewReader(reader),
		bufWriter: bufio.NewWriter(writer),
	}
}

func switchPlayer(currentPlayer string) string {
	if currentPlayer == "X" {
		return "O"
	}
	return "X"
}

func getGameStatus(board Board) GameStatus {
	if winner := checkWinner(board); winner != "" {
		if winner == "X" {
			return XWins
		}
		return OWins
	}

	if len(getAvailableMoves(board)) == 0 {
		return Draw
	}

	return InProgress
}

func (game *Game) takeTurns() {
	currentPlayer := "X"

	for {
		displayPlayerTurn(currentPlayer, game.bufWriter)

		position := getValidUserMove(&game.board, game.bufReader, game.bufWriter)
		makeMove(&game.board, position, currentPlayer)

		displayCurrentBoard(game.board, game.bufWriter)

		status := getGameStatus(game.board)

		if status != InProgress {
			game.displayEndResult(status)
			break
		}

		currentPlayer = switchPlayer(currentPlayer)
	}
}

func (game *Game) start() {
	displayWelcome(game.bufWriter)
	displayCurrentBoard(game.board, game.bufWriter)
	game.takeTurns()
}

func startGame(reader io.Reader, writer io.Writer) {
	game := newGame(reader, writer)
	game.start()
}
