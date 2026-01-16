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

func (co *ConsoleOutput) ShowModeSelection() {
	fmt.Fprintln(co.writer, "Tic-Tac-Toe Game Modes")
	fmt.Fprintln(co.writer, "1. Human vs Human")
	fmt.Fprintln(co.writer, "2. Human vs AI")
	fmt.Fprintln(co.writer, "3. AI vs AI")
	fmt.Fprint(co.writer, "Select mode (1-3): ")
	co.writer.Flush()
}

func (co *ConsoleOutput) ShowModeStart(message string) {
	fmt.Fprintf(co.writer, "Starting %s mode...\n", message)
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
