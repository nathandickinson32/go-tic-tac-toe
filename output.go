package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type ConsoleOutput struct {
	writer *bufio.Writer
}

func NewConsoleOutput(writer io.Writer) *ConsoleOutput {
	return &ConsoleOutput{
		writer: bufio.NewWriter(writer),
	}
}

func (co *ConsoleOutput) ShowWelcome() {
	fmt.Fprintln(co.writer, "Welcome to Tic-Tac-Toe!")
	co.writer.Flush()
}

func (co *ConsoleOutput) ShowBoard(board Board) {
	fmt.Fprintln(co.writer, "")
	fmt.Fprintln(co.writer, co.formatBoard(board))
	fmt.Fprintln(co.writer, "")
	co.writer.Flush()
}

func (co *ConsoleOutput) ShowPlayerTurn(player string) {
	fmt.Fprintf(co.writer, "Player %s's turn\n", player)
	co.writer.Flush()
}

func (co *ConsoleOutput) ShowPrompt() {
	fmt.Fprint(co.writer, "Enter your move (1-9): ")
	co.writer.Flush()
}

func (co *ConsoleOutput) ShowInvalidInput(err error) {
	fmt.Fprintf(co.writer, "Invalid input: %v\n", err)
	co.writer.Flush()
}

func (co *ConsoleOutput) ShowPositionTaken() {
	fmt.Fprintln(co.writer, "Position already taken, try again")
	co.writer.Flush()
}

func (co *ConsoleOutput) ShowWinner(player string) {
	fmt.Fprintf(co.writer, "Player %s wins!\n", player)
	co.writer.Flush()
}

func (co *ConsoleOutput) ShowDraw() {
	fmt.Fprintln(co.writer, "Game Over! Board is full.")
	co.writer.Flush()
}

func (co *ConsoleOutput) ShowPlayerTypeSelection(player string) {
	fmt.Fprintf(co.writer, "Select Player %s type:\n", player)
	fmt.Fprintln(co.writer, "1. Human")
	fmt.Fprintln(co.writer, "2. AI")
	fmt.Fprintf(co.writer, "Enter choice (1-2): ")
	co.writer.Flush()
}

func (co *ConsoleOutput) ShowFirstPlayerSelection() {
	fmt.Fprintln(co.writer, "Who goes first?")
	fmt.Fprintln(co.writer, "1. Player X")
	fmt.Fprintln(co.writer, "2. Player O")
	fmt.Fprint(co.writer, "Enter choice (1-2): ")
	co.writer.Flush()
}

func (co *ConsoleOutput) ShowPlayAgainPrompt() {
	fmt.Fprintln(co.writer, "")
	fmt.Fprint(co.writer, "Play again? (y/n): ")
	co.writer.Flush()
}

func (co *ConsoleOutput) ShowGoodbye() {
	fmt.Fprintln(co.writer, "Thanks for playing!")
	co.writer.Flush()
}

func (co *ConsoleOutput) ShowNewline() {
	fmt.Fprintln(co.writer, "")
	co.writer.Flush()
}

func (co *ConsoleOutput) formatBoard(board Board) string {
	var display strings.Builder

	for row := range 3 {
		display.WriteString(" " + board[row][0] + " | " + board[row][1] + " | " + board[row][2] + " ")

		if row < 2 {
			display.WriteString("\n-----------\n")
		}
	}

	return display.String()
}
