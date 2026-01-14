package main

import (
	"bytes"
	"strings"
	"testing"
)

func runGame(input string) string {
	reader := strings.NewReader(input)
	var output bytes.Buffer

	startGame(reader, &output)

	return output.String()
}

func TestPlayGame(t *testing.T) {
	t.Run("displays welcome message", func(t *testing.T) {
		input := "1\n2\n3\n4\n5\n6\n7\n8\n9\n"
		got := runGame(input)
		want := "Welcome to Tic-Tac-Toe!"

		if !strings.Contains(got, want) {
			t.Errorf("output missing welcome message.\nGot:\n%s", got)
		}
	})

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

		if !strings.Contains(got, " 1 | 2 | 3 \n-----------\n 4 | X | 6 ") {
			t.Error("board not displayed after X's first move")
		}

		if !strings.Contains(got, " O | 2 | 3 \n-----------\n 4 | X | 6 ") {
			t.Error("board not displayed after O's first move")
		}
	})

	t.Run("player X starts first", func(t *testing.T) {
		input := "1\n2\n3\n4\n5\n6\n7\n8\n9\n"
		got := runGame(input)

		xPos := strings.Index(got, "Player X's turn")
		oPos := strings.Index(got, "Player O's turn")

		if xPos == -1 {
			t.Error("Player X's turn not found")
		}

		if oPos == -1 {
			t.Error("Player O's turn not found")
		}

		if xPos >= oPos {
			t.Error("Player X should go before Player O")
		}
	})
}

func TestEndGame(t *testing.T) {

	t.Run("X wins", func(t *testing.T) {
		input := "1\n4\n2\n5\n3\n"
		output := runGame(input)

		if !strings.Contains(output, "Player X wins!") {
			t.Fatalf("winner not announced")
		}

		if strings.Count(output, "Enter your move") != 5 {
			t.Fatalf("game should stop after 5 moves")
		}
	})

	t.Run("O wins", func(t *testing.T) {
		input := "1\n4\n2\n5\n9\n6\n"
		output := runGame(input)

		if !strings.Contains(output, "Player O wins!") {
			t.Fatal("O should win")
		}
	})

	t.Run("draw game", func(t *testing.T) {
		input := "1\n2\n3\n5\n4\n6\n8\n7\n9\n"
		output := runGame(input)

		if !strings.Contains(output, "Game Over") {
			t.Fatal("expected draw game over")
		}
	})
}
