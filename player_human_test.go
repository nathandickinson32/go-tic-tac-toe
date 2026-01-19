package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestConsoleInput_ParseInput(t *testing.T) {
	var output bytes.Buffer
	consoleOutput := NewConsoleOutput(&output)
	ci := NewConsoleInput(bufio.NewReader(strings.NewReader("")), consoleOutput)

	t.Run("non-number", func(t *testing.T) {
		input := "abc"

		_, err := ci.parseInput(input)

		if err == nil {
			t.Error("non-number got nil")
		}

		want := "Input must be a number"
		if err.Error() != want {
			t.Errorf("got error %q, want %q", err.Error(), want)
		}
	})

	t.Run("empty input", func(t *testing.T) {
		input := ""

		_, err := ci.parseInput(input)

		if err == nil {
			t.Error("empty input got nil")
		}

		want := "Input cannot be empty"
		if err.Error() != want {
			t.Errorf("got error %q, want %q", err.Error(), want)
		}
	})

	t.Run("out of range numbers", func(t *testing.T) {
		tests := []string{"0", "10", "100", "-1"}

		for _, input := range tests {
			_, err := ci.parseInput(input)

			if err == nil {
				t.Errorf("expected error for input %q, got nil", input)
			}

			want := "Position must be between 1 and 9"
			if err.Error() != want {
				t.Errorf("input %q: got error %q, want %q", input, err.Error(), want)
			}
		}
	})

	t.Run("multi-digit", func(t *testing.T) {
		input := "55"

		_, err := ci.parseInput(input)

		if err == nil {
			t.Error("multi-digit input got nil")
		}

		want := "Position must be between 1 and 9"
		if err.Error() != want {
			t.Errorf("got error %q, want %q", err.Error(), want)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		input := "5"

		got, err := ci.parseInput(input)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got != 5 {
			t.Errorf("got %d, want 5", got)
		}
	})

	t.Run("valid input with whitespace", func(t *testing.T) {
		input := "  7  "

		got, err := ci.parseInput(input)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got != 7 {
			t.Errorf("got %d, want 7", got)
		}
	})
}

func TestConsoleInput_ReadMove(t *testing.T) {
	t.Run("empty input retries", func(t *testing.T) {
		board := NewBoard()
		input := "\n\n5\n"
		var output bytes.Buffer
		consoleOutput := NewConsoleOutput(&output)
		ci := NewConsoleInput(bufio.NewReader(strings.NewReader(input)), consoleOutput)

		got, err := ci.ReadMove(board)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got != 5 {
			t.Errorf("got %d, want 5", got)
		}
	})

	t.Run("non-number retries", func(t *testing.T) {
		board := NewBoard()
		input := "abc\n5\n"
		var output bytes.Buffer
		consoleOutput := NewConsoleOutput(&output)
		ci := NewConsoleInput(bufio.NewReader(strings.NewReader(input)), consoleOutput)

		got, err := ci.ReadMove(board)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got != 5 {
			t.Errorf("got %d, want 5", got)
		}
	})

	t.Run("out of range input retries", func(t *testing.T) {
		board := NewBoard()
		input := "0\n10\n-1\n5\n"
		var output bytes.Buffer
		consoleOutput := NewConsoleOutput(&output)
		ci := NewConsoleInput(bufio.NewReader(strings.NewReader(input)), consoleOutput)

		got, err := ci.ReadMove(board)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got != 5 {
			t.Errorf("got %d, want 5", got)
		}
	})

	t.Run("occupied position retries", func(t *testing.T) {
		board := NewBoard()
		board.MakeMove(5, "X")

		input := "5\n7\n"
		var output bytes.Buffer
		consoleOutput := NewConsoleOutput(&output)
		ci := NewConsoleInput(bufio.NewReader(strings.NewReader(input)), consoleOutput)

		got, err := ci.ReadMove(board)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got != 7 {
			t.Errorf("got %d, want 7", got)
		}
	})

	t.Run("whitespace around valid input", func(t *testing.T) {
		board := NewBoard()
		input := "  5  \n"
		var output bytes.Buffer
		consoleOutput := NewConsoleOutput(&output)
		ci := NewConsoleInput(bufio.NewReader(strings.NewReader(input)), consoleOutput)

		got, err := ci.ReadMove(board)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got != 5 {
			t.Errorf("got %d, want 5", got)
		}
	})

	t.Run("valid move", func(t *testing.T) {
		board := NewBoard()
		input := "5\n"
		var output bytes.Buffer
		consoleOutput := NewConsoleOutput(&output)
		ci := NewConsoleInput(bufio.NewReader(strings.NewReader(input)), consoleOutput)

		got, err := ci.ReadMove(board)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got != 5 {
			t.Errorf("got %d, want 5", got)
		}
	})

	t.Run("displays prompt", func(t *testing.T) {
		board := NewBoard()
		input := "5\n"
		var output bytes.Buffer

		consoleOutput := NewConsoleOutput(&output)
		ci := NewConsoleInput(bufio.NewReader(strings.NewReader(input)), consoleOutput)

		_, err := ci.ReadMove(board)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		want := "Enter your move (1-9): "
		got := output.String()

		if !strings.Contains(got, want) {
			t.Errorf("output missing prompt.\nGot:\n%s\n\nWant to contain:\n%s", got, want)
		}
	})

	t.Run("shows error for empty input", func(t *testing.T) {
		board := NewBoard()
		input := "\n5\n"
		var output bytes.Buffer

		consoleOutput := NewConsoleOutput(&output)
		ci := NewConsoleInput(bufio.NewReader(strings.NewReader(input)), consoleOutput)

		_, err := ci.ReadMove(board)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		want := "Invalid input: Input cannot be empty"
		got := output.String()

		if !strings.Contains(got, want) {
			t.Errorf("output missing error message.\nGot:\n%s\n\nWant to contain:\n%s", got, want)
		}
	})

	t.Run("shows error for non-number", func(t *testing.T) {
		board := NewBoard()
		input := "abc\n5\n"
		var output bytes.Buffer

		consoleOutput := NewConsoleOutput(&output)
		ci := NewConsoleInput(bufio.NewReader(strings.NewReader(input)), consoleOutput)

		_, err := ci.ReadMove(board)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		want := "Invalid input: Input must be a number"
		got := output.String()

		if !strings.Contains(got, want) {
			t.Errorf("output missing error message.\nGot:\n%s\n\nWant to contain:\n%s", got, want)
		}
	})

	t.Run("shows error for out of range", func(t *testing.T) {
		board := NewBoard()
		input := "10\n5\n"
		var output bytes.Buffer

		consoleOutput := NewConsoleOutput(&output)
		ci := NewConsoleInput(bufio.NewReader(strings.NewReader(input)), consoleOutput)

		_, err := ci.ReadMove(board)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		want := "Invalid input: Position must be between 1 and 9"
		got := output.String()

		if !strings.Contains(got, want) {
			t.Errorf("output missing error message.\nGot:\n%s\n\nWant to contain:\n%s", got, want)
		}
	})

	t.Run("shows error for occupied position", func(t *testing.T) {
		board := NewBoard()
		board.MakeMove(5, "X")
		input := "5\n7\n"
		var output bytes.Buffer

		consoleOutput := NewConsoleOutput(&output)
		ci := NewConsoleInput(bufio.NewReader(strings.NewReader(input)), consoleOutput)

		_, err := ci.ReadMove(board)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		want := "Position already taken, try again"
		got := output.String()

		if !strings.Contains(got, want) {
			t.Errorf("output missing error message.\nGot:\n%s\n\nWant to contain:\n%s", got, want)
		}
	})

	t.Run("multiple errors", func(t *testing.T) {
		board := NewBoard()
		board.MakeMove(1, "X")
		input := "\nabc\n10\n1\n5\n"
		var output bytes.Buffer

		consoleOutput := NewConsoleOutput(&output)
		ci := NewConsoleInput(bufio.NewReader(strings.NewReader(input)), consoleOutput)

		_, err := ci.ReadMove(board)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

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
