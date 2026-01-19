package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestConsoleOutput_FormatBoard(t *testing.T) {
	var output bytes.Buffer
	co := NewConsoleOutput(&output)

	t.Run("empty board", func(t *testing.T) {
		board := NewBoard()

		want := " 1 | 2 | 3 \n" +
			"-----------\n" +
			" 4 | 5 | 6 \n" +
			"-----------\n" +
			" 7 | 8 | 9 "

		got := co.formatBoard(board)

		if got != want {
			t.Errorf("Board mismatch.\nGot:\n%s\n\nWant:\n%s", got, want)
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

		got := co.formatBoard(board)

		if got != want {
			t.Errorf("Board mismatch.\nGot:\n%s\n\nWant:\n%s", got, want)
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

		got := co.formatBoard(board)

		if got != want {
			t.Errorf("Board mismatch")
		}
	})
}

func TestConsoleOutput_ShowWelcome(t *testing.T) {
	var output bytes.Buffer
	co := NewConsoleOutput(&output)

	co.ShowWelcome()

	want := "Welcome to Tic-Tac-Toe!"
	got := output.String()

	if !strings.Contains(got, want) {
		t.Errorf("output missing.\nGot:\n%s\n\nWant to contain:\n%s", got, want)
	}
}

func TestConsoleOutput_ShowBoard(t *testing.T) {
	var output bytes.Buffer
	co := NewConsoleOutput(&output)

	board := Board{
		{"X", "2", "3"},
		{"4", "O", "6"},
		{"7", "8", "9"},
	}

	co.ShowBoard(board)

	got := output.String()

	if !strings.Contains(got, " X | 2 | 3 ") {
		t.Error("first row not displayed")
	}

	if !strings.Contains(got, " 4 | O | 6 ") {
		t.Error("second row not displayed")
	}

	if !strings.Contains(got, "-----------") {
		t.Error("separator not displayed")
	}
}

func TestConsoleOutput_ShowPlayerTurn(t *testing.T) {
	var output bytes.Buffer
	co := NewConsoleOutput(&output)

	co.ShowPlayerTurn("X")

	want := "Player X's turn"
	got := output.String()

	if !strings.Contains(got, want) {
		t.Errorf("missing player turn.\nGot:\n%s\n\nWant to contain:\n%s", got, want)
	}
}

func TestConsoleOutput_ShowPrompt(t *testing.T) {
	var output bytes.Buffer
	co := NewConsoleOutput(&output)

	co.ShowPrompt()

	want := "Enter your move (1-9): "
	got := output.String()

	if !strings.Contains(got, want) {
		t.Errorf("missing prompt.\nGot:\n%s\n\nWant to contain:\n%s", got, want)
	}
}

type testError struct {
	msg string
}

func (e testError) Error() string {
	return e.msg
}

func TestInputValidation(t *testing.T) {
	t.Run("rejects invalid input and continues", func(t *testing.T) {
		input := "abc\n10\n-1\n1\n1\n2\n3\n5\n4\n6\n8\n7\n9\n"
		reader := strings.NewReader(input)
		var output bytes.Buffer

		rules := NewGameRules()
		consoleOutput := NewConsoleOutput(&output)
		consoleInput := NewConsoleInput(bufio.NewReader(reader), consoleOutput)
		game := NewGame(rules, consoleInput, consoleInput, consoleOutput, "X")

		consoleOutput.ShowWelcome()
		consoleOutput.ShowBoard(NewBoard())
		game.playTurns()

		result := output.String()

		if !strings.Contains(result, "Invalid input: Input must be a number") {
			t.Error("should show error for non-number")
		}

		if !strings.Contains(result, "Invalid input: Position must be between 1 and 9") {
			t.Error("should show error for out of range")
		}

		if !strings.Contains(result, "Game Over") {
			t.Error("game should complete after invalid input")
		}

		if !strings.Contains(result, "Position already taken") {
			t.Error("should show error for occupied position")
		}
	})
}

func TestConsoleOutput_ShowWinner(t *testing.T) {
	t.Run("X wins", func(t *testing.T) {
		var output bytes.Buffer
		co := NewConsoleOutput(&output)

		co.ShowWinner("X")

		want := "Player X wins!"
		got := output.String()

		if !strings.Contains(got, want) {
			t.Errorf("missing winner message.\nGot:\n%s\n\nWant to contain:\n%s", got, want)
		}
	})

	t.Run("O wins", func(t *testing.T) {
		var output bytes.Buffer
		co := NewConsoleOutput(&output)

		co.ShowWinner("O")

		want := "Player O wins!"
		got := output.String()

		if !strings.Contains(got, want) {
			t.Errorf("missing winner message.\nGot:\n%s\n\nWant to contain:\n%s", got, want)
		}
	})
}

func TestConsoleOutput_ShowDraw(t *testing.T) {
	var output bytes.Buffer
	co := NewConsoleOutput(&output)

	co.ShowDraw()

	want := "Game Over! Board is full."
	got := output.String()

	if !strings.Contains(got, want) {
		t.Errorf("missing draw message.\nGot:\n%s\n\nWant to contain:\n%s", got, want)
	}
}

func TestConsoleOutput_BuildGame(t *testing.T) {
	var output bytes.Buffer
	co := NewConsoleOutput(&output)

	t.Run("shows player type selection", func(t *testing.T) {
		co.ShowPlayerTypeSelection("X")

		result := output.String()

		if !strings.Contains(result, "Select Player X type:") {
			t.Error("should show player selection prompt")
		}

		if !strings.Contains(result, "1. Human") {
			t.Error("should show Human option")
		}

		if !strings.Contains(result, "2. AI") {
			t.Error("should show AI option")
		}
	})

	t.Run("shows first player selection", func(t *testing.T) {
		output.Reset()
		co.ShowFirstPlayerSelection()

		result := output.String()

		if !strings.Contains(result, "Who goes first?") {
			t.Error("should show who goes first prompt")
		}

		if !strings.Contains(result, "1. Player X") {
			t.Error("should show Player X option")
		}

		if !strings.Contains(result, "2. Player O") {
			t.Error("should show Player O option")
		}
	})

	t.Run("shows play again prompt", func(t *testing.T) {
		output.Reset()
		co.ShowPlayAgainPrompt()

		result := output.String()

		if !strings.Contains(result, "Play again? (y/n):") {
			t.Error("should show play again prompt")
		}
	})

	t.Run("shows goodbye message", func(t *testing.T) {
		output.Reset()
		co.ShowGoodbye()

		result := output.String()

		if !strings.Contains(result, "Thanks for playing!") {
			t.Error("should show goodbye message")
		}
	})
}
