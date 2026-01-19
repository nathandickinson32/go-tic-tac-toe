package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestConsoleInput_ReadPlayerType(t *testing.T) {
	var output bytes.Buffer
	consoleOutput := NewConsoleOutput(&output)

	t.Run("reads human player type", func(t *testing.T) {
		input := strings.NewReader("1\n")
		ci := NewConsoleInput(bufio.NewReader(input), consoleOutput)

		playerType, err := ci.ReadPlayerType()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if playerType != Human {
			t.Errorf("expected Human, got %v", playerType)
		}
	})

	t.Run("reads AI player type", func(t *testing.T) {
		input := strings.NewReader("2\n")
		ci := NewConsoleInput(bufio.NewReader(input), consoleOutput)

		playerType, err := ci.ReadPlayerType()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if playerType != AI {
			t.Errorf("expected AI, got %v", playerType)
		}
	})

	t.Run("retries on invalid input", func(t *testing.T) {
		input := strings.NewReader("3\nabc\n1\n")
		output.Reset()
		ci := NewConsoleInput(bufio.NewReader(input), consoleOutput)

		playerType, err := ci.ReadPlayerType()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if playerType != Human {
			t.Errorf("expected Human after retries, got %v", playerType)
		}

		result := output.String()
		if !strings.Contains(result, "Invalid input") {
			t.Error("should show invalid input error")
		}
	})
}

func TestConsoleInput_ReadFirstPlayer(t *testing.T) {
	var output bytes.Buffer
	consoleOutput := NewConsoleOutput(&output)

	t.Run("reads Player X first", func(t *testing.T) {
		input := strings.NewReader("1\n")
		ci := NewConsoleInput(bufio.NewReader(input), consoleOutput)

		firstPlayer, err := ci.ReadFirstPlayer()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if firstPlayer != "X" {
			t.Errorf("expected X, got %s", firstPlayer)
		}
	})

	t.Run("reads Player O first", func(t *testing.T) {
		input := strings.NewReader("2\n")
		ci := NewConsoleInput(bufio.NewReader(input), consoleOutput)

		firstPlayer, err := ci.ReadFirstPlayer()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if firstPlayer != "O" {
			t.Errorf("expected O, got %s", firstPlayer)
		}
	})

	t.Run("retries on invalid input", func(t *testing.T) {
		input := strings.NewReader("3\nabc\n1\n")
		output.Reset()
		ci := NewConsoleInput(bufio.NewReader(input), consoleOutput)

		firstPlayer, err := ci.ReadFirstPlayer()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if firstPlayer != "X" {
			t.Errorf("expected X after retries, got %s", firstPlayer)
		}

		result := output.String()
		if !strings.Contains(result, "Invalid input") {
			t.Error("should show invalid input error")
		}
	})
}

func TestParseFirstPlayer(t *testing.T) {
	var output bytes.Buffer
	consoleOutput := NewConsoleOutput(&output)
	ci := NewConsoleInput(bufio.NewReader(nil), consoleOutput)

	tests := []struct {
		input    string
		expected string
		wantErr  bool
	}{
		{"1", "X", false},
		{"2", "O", false},
		{"", "X", true},
		{"3", "X", true},
		{"abc", "X", true},
		{"0", "X", true},
	}

	for _, tt := range tests {
		result, err := ci.parseFirstPlayer(tt.input)

		if tt.wantErr {
			if err == nil {
				t.Errorf("input %q: expected error, got nil", tt.input)
			}
		} else {
			if err != nil {
				t.Errorf("input %q: unexpected error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("input %q: expected %v, got %v", tt.input, tt.expected, result)
			}
		}
	}
}

func TestGameWithFirstPlayer(t *testing.T) {
	rules := NewGameRules()
	var output bytes.Buffer
	consoleOutput := NewConsoleOutput(&output)

	t.Run("game starts with Player O when O goes first", func(t *testing.T) {
		input := strings.NewReader("5\n1\n9\n2\n3\n4\n6\n7\n8\n")
		consoleInput := NewConsoleInput(bufio.NewReader(input), consoleOutput)

		game := NewGame(rules, consoleInput, consoleInput, consoleOutput, "O")
		game.PlayGame()

		result := output.String()

		firstOPos := strings.Index(result, "Player O's turn")
		firstXPos := strings.Index(result, "Player X's turn")

		if firstOPos == -1 {
			t.Error("Player O's turn not found")
		}

		if firstXPos == -1 {
			t.Error("Player X's turn not found")
		}

		if firstOPos >= firstXPos {
			t.Error("Player O should go before Player X when O goes first")
		}
	})

	t.Run("game starts with Player X when X goes first", func(t *testing.T) {
		output.Reset()
		input := strings.NewReader("5\n1\n9\n2\n3\n4\n6\n7\n8\n")
		consoleInput := NewConsoleInput(bufio.NewReader(input), consoleOutput)

		game := NewGame(rules, consoleInput, consoleInput, consoleOutput, "X")
		game.PlayGame()

		result := output.String()

		firstXPos := strings.Index(result, "Player X's turn")
		firstOPos := strings.Index(result, "Player O's turn")

		if firstXPos == -1 {
			t.Error("Player X's turn not found")
		}

		if firstOPos == -1 {
			t.Error("Player O's turn not found")
		}

		if firstXPos >= firstOPos {
			t.Error("Player X should go before Player O when X goes first")
		}
	})
}

func TestCreatePlayerInput(t *testing.T) {
	rules := NewGameRules()
	var output bytes.Buffer
	consoleOutput := NewConsoleOutput(&output)
	reader := bufio.NewReader(strings.NewReader(""))

	t.Run("creates Human player input", func(t *testing.T) {
		playerInput := createPlayerInput(Human, "X", "O", reader, consoleOutput, rules)

		_, isConsoleInput := playerInput.(*ConsoleInput)
		if !isConsoleInput {
			t.Error("expected ConsoleInput for Human player")
		}
	})

	t.Run("creates AI player input", func(t *testing.T) {
		playerInput := createPlayerInput(AI, "X", "O", reader, consoleOutput, rules)

		aiPlayer, isAI := playerInput.(*AIPlayer)
		if !isAI {
			t.Error("expected AIPlayer for AI player type")
		}

		if aiPlayer != nil && aiPlayer.playerSymbol != "X" {
			t.Errorf("expected AI player symbol X, got %s", aiPlayer.playerSymbol)
		}
	})

	t.Run("creates AI with opponent symbol", func(t *testing.T) {
		playerInput := createPlayerInput(AI, "O", "X", reader, consoleOutput, rules)

		aiPlayer, isAI := playerInput.(*AIPlayer)
		if !isAI {
			t.Fatal("expected AIPlayer for AI player type")
		}

		if aiPlayer.playerSymbol != "O" {
			t.Errorf("expected AI player symbol O, got %s", aiPlayer.playerSymbol)
		}

		if aiPlayer.opponentSymbol != "X" {
			t.Errorf("expected AI opponent symbol X, got %s", aiPlayer.opponentSymbol)
		}
	})
}

func TestParsePlayerType(t *testing.T) {
	var output bytes.Buffer
	consoleOutput := NewConsoleOutput(&output)
	ci := NewConsoleInput(bufio.NewReader(strings.NewReader("")), consoleOutput)

	tests := []struct {
		input    string
		expected PlayerType
		wantErr  bool
	}{
		{"1", Human, false},
		{"2", AI, false},
		{"", Human, true},
		{"3", Human, true},
		{"abc", Human, true},
		{"0", Human, true},
	}

	for _, test := range tests {
		result, err := ci.parsePlayerType(test.input)

		if test.wantErr {
			if err == nil {
				t.Errorf("input %q: expected error, got nil", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("input %q: unexpected error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("input %q: expected %v, got %v", test.input, test.expected, result)
			}
		}
	}
}

func TestParsePlayAgain(t *testing.T) {
	var output bytes.Buffer
	consoleOutput := NewConsoleOutput(&output)
	ci := NewConsoleInput(bufio.NewReader(strings.NewReader("")), consoleOutput)

	tests := []struct {
		input    string
		expected bool
		wantErr  bool
	}{
		{"y", true, false},
		{"Y", true, false},
		{"yes", true, false},
		{"YES", true, false},
		{"Yes", true, false},
		{"n", false, false},
		{"N", false, false},
		{"no", false, false},
		{"NO", false, false},
		{"No", false, false},
		{"", false, true},
		{"maybe", false, true},
		{"invalid", false, true},
	}

	for _, test := range tests {
		result, err := ci.parsePlayAgain(test.input)

		if test.wantErr {
			if err == nil {
				t.Errorf("input %q: expected error, got nil", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("input %q: unexpected error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("input %q: expected %v, got %v", test.input, test.expected, result)
			}
		}
	}
}
func TestConsoleInput_ReadPlayAgain(t *testing.T) {
	var output bytes.Buffer
	consoleOutput := NewConsoleOutput(&output)

	t.Run("reads yes", func(t *testing.T) {
		tests := []string{"y\n", "Y\n", "yes\n", "YES\n", "Yes\n"}

		for _, test := range tests {
			input := strings.NewReader(test)
			ci := NewConsoleInput(bufio.NewReader(input), consoleOutput)

			playAgain, err := ci.ReadPlayAgain()

			if err != nil {
				t.Errorf("input %q: unexpected error: %v", test, err)
			}

			if !playAgain {
				t.Errorf("input %q: expected true, got false", test)
			}
		}
	})

	t.Run("reads no", func(t *testing.T) {
		tests := []string{"n\n", "N\n", "no\n", "NO\n", "No\n"}

		for _, test := range tests {
			input := strings.NewReader(test)
			ci := NewConsoleInput(bufio.NewReader(input), consoleOutput)

			playAgain, err := ci.ReadPlayAgain()

			if err != nil {
				t.Errorf("input %q: unexpected error: %v", test, err)
			}

			if playAgain {
				t.Errorf("input %q: expected false, got true", test)
			}
		}
	})

	t.Run("retries on invalid input", func(t *testing.T) {
		input := strings.NewReader("maybe\ninvalid\ny\n")
		output.Reset()
		ci := NewConsoleInput(bufio.NewReader(input), consoleOutput)

		playAgain, err := ci.ReadPlayAgain()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !playAgain {
			t.Error("expected true after retries")
		}

		result := output.String()
		if !strings.Contains(result, "Invalid input") {
			t.Error("should show invalid input error")
		}
	})
}
