package main

import (
	"strings"
	"testing"
)

func TestUserInput(t *testing.T) {
	t.Run("parse-input", func(t *testing.T) {

		t.Run("non-number", func(t *testing.T) {
			input := "abc"

			_, err := parseInput(input)

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

			_, err := parseInput(input)

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
				_, err := parseInput(input)

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

			_, err := parseInput(input)

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

			got, err := parseInput(input)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != 5 {
				t.Errorf("got %d, want 5", got)
			}
		})

		t.Run("valid input with whitespace", func(t *testing.T) {
			input := "  7  "

			got, err := parseInput(input)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != 7 {
				t.Errorf("got %d, want 7", got)
			}
		})
	})

	t.Run("get-user-move", func(t *testing.T) {

		t.Run("empty input", func(t *testing.T) {
			board := initBoard()
			input := "\n\n5\n"

			got := getUserMove(&board, strings.NewReader(input), nil)
			want := 5

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})

		t.Run("non-number", func(t *testing.T) {
			board := initBoard()
			input := "abc\n5\n"

			got := getUserMove(&board, strings.NewReader(input), nil)
			want := 5

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})

		t.Run("out of range input", func(t *testing.T) {
			board := initBoard()
			input := "0\n10\n-1\n5\n"

			got := getUserMove(&board, strings.NewReader(input), nil)
			want := 5

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})

		t.Run("occupied position", func(t *testing.T) {
			board := initBoard()
			makeMove(&board, 5, "X")

			input := "5\n7\n"

			got := getUserMove(&board, strings.NewReader(input), nil)
			want := 7

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})

		t.Run("whitespace around valid input", func(t *testing.T) {
			board := initBoard()
			input := "  5  \n"

			got := getUserMove(&board, strings.NewReader(input), nil)
			want := 5

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})

		t.Run("valid move", func(t *testing.T) {
			board := initBoard()
			input := "5\n"

			got := getUserMove(&board, strings.NewReader(input), nil)
			want := 5

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	})
}
