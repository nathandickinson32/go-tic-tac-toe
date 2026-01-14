package main

import (
	"bufio"
	"fmt"
	"strings"
)

func displayWelcome(writer *bufio.Writer) {
	fmt.Fprintln(writer, "Welcome to Tic-Tac-Toe!")
	writer.Flush()
}

func displayCurrentBoard(board Board, writer *bufio.Writer) {
	fmt.Fprintln(writer, displayBoard(board))
	fmt.Fprintln(writer)
	writer.Flush()
}

func displayPlayerTurn(player string, writer *bufio.Writer) {
	fmt.Fprintf(writer, "Player %s's turn\n", player)
	writer.Flush()
}

func displayGameOver(writer *bufio.Writer) {
	fmt.Fprintln(writer, "Game Over! Board is full.")
	writer.Flush()
}

func displayPrompt(writer *bufio.Writer) {
	fmt.Fprint(writer, "Enter your move (1-9): ")
	writer.Flush()
}

func displayInvalidInput(err error, writer *bufio.Writer) {
	fmt.Fprintf(writer, "Invalid input: %v\n", err)
	writer.Flush()
}

func displayPositionTaken(writer *bufio.Writer) {
	fmt.Fprintln(writer, "Position already taken, try again")
	writer.Flush()
}

func displayBoard(board Board) string {
	var display strings.Builder

	for row := range 3 {
		display.WriteString(" " + board[row][0] + " | " + board[row][1] + " | " + board[row][2] + " ")

		if row < 2 {
			display.WriteString("\n-----------\n")
		}
	}

	return display.String()
}
