package main

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"
)

func StartGame(reader io.Reader, writer io.Writer) {
	rules := NewGameRules()
	output := NewConsoleOutput(writer)
	input := NewConsoleInput(bufio.NewReader(reader), output)
	game := NewGame(rules, input, input, output, "X")
	game.PlayGame()
}

func runGame(input string) string {
	reader := strings.NewReader(input)
	var output bytes.Buffer
	StartGame(reader, &output)
	return output.String()
}

func TestGame_Flow(t *testing.T) {

	t.Run("alternates between players", func(t *testing.T) {
		input := "1\n2\n3\n4\n5\n6\n7\n8\n9\n"
		got := runGame(input)

		if !strings.Contains(got, "Player X's turn") {
			t.Error("missing Player X's turn")
		}

		if !strings.Contains(got, "Player O's turn") {
			t.Error("missing Player O's turn")
		}
	})

	t.Run("displays board after each move", func(t *testing.T) {
		input := "5\n1\n9\n2\n3\n4\n6\n7\n8\n"
		got := runGame(input)

		if !strings.Contains(got, " 1 | 2 | 3 \n-----------\n 4 | X | 6 \n-----------\n 7 | 8 | 9") {
			t.Error("board not displayed after X's first move")
		}

		if !strings.Contains(got, " O | 2 | 3 \n-----------\n 4 | X | 6 \n-----------\n 7 | 8 | 9") {
			t.Error("board not displayed after O's first move")
		}
	})

	t.Run("displays initial board", func(t *testing.T) {
		input := "1\n2\n3\n4\n5\n6\n7\n8\n9\n"
		got := runGame(input)
		if !strings.Contains(got, " 1 | 2 | 3 \n-----------\n 4 | 5 | 6 \n-----------\n 7 | 8 | 9") {
			t.Error("initial board not displayed")
		}
	})
}

func TestGame_EndConditions(t *testing.T) {
	t.Run("X wins horizontally", func(t *testing.T) {
		input := "1\n4\n2\n5\n3\n"
		output := runGame(input)

		if !strings.Contains(output, "Player X wins!") {
			t.Fatalf("winner not announced")
		}

		if strings.Count(output, "Enter your move") != 5 {
			t.Fatalf("game should stop after 5 moves")
		}
	})

	t.Run("O wins horizontally", func(t *testing.T) {
		input := "1\n4\n2\n5\n9\n6\n"
		output := runGame(input)

		if !strings.Contains(output, "Player O wins!") {
			t.Fatal("O should win")
		}

		if strings.Count(output, "Enter your move") != 6 {
			t.Fatal("game should stop after 6 moves")
		}
	})

	t.Run("X wins vertically", func(t *testing.T) {
		input := "1\n2\n4\n3\n7\n"
		output := runGame(input)

		if !strings.Contains(output, "Player X wins!") {
			t.Fatal("X should win vertically")
		}
	})

	t.Run("X wins diagonally", func(t *testing.T) {
		input := "1\n2\n5\n3\n9\n"
		output := runGame(input)

		if !strings.Contains(output, "Player X wins!") {
			t.Fatal("X should win diagonally")
		}
	})

	t.Run("draw game", func(t *testing.T) {
		input := "1\n2\n3\n5\n4\n6\n8\n7\n9\n"
		output := runGame(input)

		if !strings.Contains(output, "Game Over") {
			t.Fatal("expected draw game over")
		}

		if !strings.Contains(output, "Board is full") {
			t.Fatal("expected full board message")
		}
	})

	t.Run("game stops after win", func(t *testing.T) {
		input := "1\n4\n2\n5\n3\n"
		output := runGame(input)

		promptCount := strings.Count(output, "Enter your move")
		if promptCount != 5 {
			t.Errorf("expected 5 prompts, got %d", promptCount)
		}

		if strings.Count(output, "Player X wins!") != 1 {
			t.Error("should announce winner exactly once")
		}
	})
}
