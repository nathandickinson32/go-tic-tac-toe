package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestStartGameWithMode_HumanVsHuman(t *testing.T) {
	t.Run("human vs human mode X wins", func(t *testing.T) {
		input := strings.NewReader("1\n4\n2\n5\n3\n")
		var output bytes.Buffer

		StartGameWithMode(input, &output, HumanVsHuman)

		result := output.String()

		if !strings.Contains(result, "Welcome to Tic-Tac-Toe!") {
			t.Error("should display welcome")
		}

		if !strings.Contains(result, "Player X wins!") {
			t.Error(result)
		}
	})

	t.Run("human vs human mode O wins", func(t *testing.T) {
		input := strings.NewReader("1\n4\n2\n5\n9\n6\n")
		var output bytes.Buffer

		StartGameWithMode(input, &output, HumanVsHuman)

		result := output.String()

		if !strings.Contains(result, "Welcome to Tic-Tac-Toe!") {
			t.Error("should display welcome")
		}

		if !strings.Contains(result, "Player O wins!") {
			t.Error(result)
		}
	})
}

func TestStartGameWithMode_HumanVsAI(t *testing.T) {
	t.Run("human vs AI mode", func(t *testing.T) {
		input := strings.NewReader("1\n3\n4\n6\n8\n9\n7\n2\n5\n")
		var output bytes.Buffer

		StartGameWithMode(input, &output, HumanVsAI)

		result := output.String()

		if !strings.Contains(result, "Welcome to Tic-Tac-Toe!") {
			t.Error("should display welcome")
		}

		hasEnding := strings.Contains(result, "wins!")
		if !hasEnding {
			t.Error("AI should win")
		}
	})
}

func TestStartGameWithMode_AIVsAI(t *testing.T) {
	t.Run("AI vs AI mode always draws", func(t *testing.T) {
		input := strings.NewReader("")
		var output bytes.Buffer

		StartGameWithMode(input, &output, AIVsAI)

		result := output.String()

		if !strings.Contains(result, "Welcome to Tic-Tac-Toe!") {
			t.Error("should display welcome")
		}

		if !strings.Contains(result, "Game Over") {
			t.Error("AIs should draw")
		}
	})
}
