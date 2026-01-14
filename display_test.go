package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestDisplayBoard(t *testing.T) {
	t.Run("empty board", func(t *testing.T) {
		board := initBoard()

		want := " 1 | 2 | 3 \n" +
			"-----------\n" +
			" 4 | 5 | 6 \n" +
			"-----------\n" +
			" 7 | 8 | 9 "

		got := displayBoard(board)

		if got != want {
			t.Errorf("Board display mismatch.\nGot:\n%s\n\nWant:\n%s", got, want)
		}
	})

	t.Run("board with moves", func(t *testing.T) {
		board := Board{
			{"X", "2", "O"},
			{"4", "X", "6"},
			{"O", "8", "9"},
		}

		want := " X | 2 | O \n" +
			"-----------\n" +
			" 4 | X | 6 \n" +
			"-----------\n" +
			" O | 8 | 9 "

		got := displayBoard(board)

		if got != want {
			t.Errorf("Board display mismatch.\nGot:\n%s\n\nWant:\n%s", got, want)
		}
	})

	t.Run("full board", func(t *testing.T) {
		board := Board{
			{"X", "O", "X"},
			{"O", "X", "O"},
			{"O", "X", "X"},
		}

		want := " X | O | X \n" +
			"-----------\n" +
			" O | X | O \n" +
			"-----------\n" +
			" O | X | X "

		got := displayBoard(board)

		if got != want {
			t.Errorf("Board display mismatch")
		}
	})
}

func TestDisplayInputErrors(t *testing.T) {

	t.Run("displays prompt", func(t *testing.T) {
		board := initBoard()
		input := "5\n"
		var output bytes.Buffer

		getUserMove(&board, strings.NewReader(input), &output)

		want := "Enter your move (1-9): "
		got := output.String()

		if !strings.Contains(got, want) {
			t.Errorf("output missing prompt.\nGot:\n%s\n\nWant to contain:\n%s", got, want)
		}
	})

	t.Run("empty input", func(t *testing.T) {
		board := initBoard()
		input := "\n5\n"
		var output bytes.Buffer

		getUserMove(&board, strings.NewReader(input), &output)

		want := "Invalid input: Input cannot be empty"
		got := output.String()

		if !strings.Contains(got, want) {
			t.Errorf("output missing error message.\nGot:\n%s\n\nWant to contain:\n%s", got, want)
		}
	})

	t.Run("non-number", func(t *testing.T) {
		board := initBoard()
		input := "abc\n5\n"
		var output bytes.Buffer

		getUserMove(&board, strings.NewReader(input), &output)

		want := "Invalid input: Input must be a number"
		got := output.String()

		if !strings.Contains(got, want) {
			t.Errorf("output missing error message.\nGot:\n%s\n\nWant to contain:\n%s", got, want)
		}
	})

	t.Run("out of range", func(t *testing.T) {
		board := initBoard()
		input := "10\n5\n"
		var output bytes.Buffer

		getUserMove(&board, strings.NewReader(input), &output)

		want := "Invalid input: Position must be between 1 and 9"
		got := output.String()

		if !strings.Contains(got, want) {
			t.Errorf("output missing error message.\nGot:\n%s\n\nWant to contain:\n%s", got, want)
		}
	})

	t.Run("occupied position", func(t *testing.T) {
		board := initBoard()
		makeMove(&board, 5, "X")
		input := "5\n7\n"
		var output bytes.Buffer

		getUserMove(&board, strings.NewReader(input), &output)

		want := "Position already taken, try again"
		got := output.String()

		if !strings.Contains(got, want) {
			t.Errorf("output missing error message.\nGot:\n%s\n\nWant to contain:\n%s", got, want)
		}
	})

	t.Run("multiple errors", func(t *testing.T) {
		board := initBoard()
		makeMove(&board, 1, "X")
		input := "\nabc\n10\n1\n5\n"
		var output bytes.Buffer

		getUserMove(&board, strings.NewReader(input), &output)

		got := output.String()

		wantMessages := []string{
			"Invalid input: Input cannot be empty",
			"Invalid input: Input must be a number",
			"Invalid input: Position must be between 1 and 9",
			"Position already taken, try again",
		}

		for _, want := range wantMessages {
			if !strings.Contains(got, want) {
				t.Errorf("output missing error: %q\nGot:\n%s", want, got)
			}
		}
	})

}
