package io

import (
	"fmt"
	"io"
	"strings"
	"ttt/boards"
)

const (
	PositionRange   = "1-9"
	PlayerTypeRange = "1-2"
	GridSeparator   = "-----------"
	GridDivider     = " | "
	CellPadding     = " "
	NewlineChar     = "\n"
	RowsPerBoard    = 3
	ColumnsPerRow   = 3
)

func ShowWelcome(writer io.Writer) {
	fmt.Fprintln(writer, "Welcome to Tic-Tac-Toe!")
}

func ShowBoard(writer io.Writer, board boards.Board) {
	fmt.Fprintln(writer, "")
	fmt.Fprintln(writer, formatBoard(board))
	fmt.Fprintln(writer, "")
}

func ShowPlayerTurn(writer io.Writer, player string) {
	fmt.Fprintf(writer, "Player %s's turn\n", player)
}

func ShowPrompt(writer io.Writer) {
	fmt.Fprintf(writer, "Enter your move (%s): ", PositionRange)
}

func ShowInvalidInput(writer io.Writer, err error) {
	fmt.Fprintf(writer, "Invalid input: %v\n", err)
}

func ShowPositionTaken(writer io.Writer) {
	fmt.Fprintln(writer, "Position already taken, try again")
}

func ShowWinner(writer io.Writer, player string) {
	fmt.Fprintf(writer, "Player %s wins!\n", player)
}

func ShowDraw(writer io.Writer) {
	fmt.Fprintln(writer, "Game Over! Board is full.")
}

func ShowPlayerTypeSelection(writer io.Writer, player string) {
	fmt.Fprintf(writer, "Select Player %s type:\n", player)
	fmt.Fprintln(writer, "1. Human")
	fmt.Fprintln(writer, "2. AI")
	fmt.Fprintf(writer, "Enter choice (%s): ", PlayerTypeRange)
}

func ShowPlayAgainPrompt(writer io.Writer) {
	fmt.Fprintln(writer, "")
	fmt.Fprint(writer, "Play again? (y/n): ")
}

func ShowGoodbye(writer io.Writer) {
	fmt.Fprintln(writer, "Thanks for playing!")
}

func ShowNewline(writer io.Writer) {
	fmt.Fprintln(writer, "")
}

func formatBoard(board boards.Board) string {
	var display strings.Builder

	for row := range RowsPerBoard {
		display.WriteString(CellPadding + board[row][0] + GridDivider + board[row][1] + GridDivider + board[row][2] + CellPadding)

		if row < RowsPerBoard-1 {
			display.WriteString(NewlineChar + GridSeparator + NewlineChar)
		}
	}

	return display.String()
}
