package io

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestReadPlayerType_SelectsHumanWith1(t *testing.T) {
	var output bytes.Buffer
	reader := bufio.NewReader(strings.NewReader("1\n"))

	playerType, err := ReadPlayerType(reader, &output)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if playerType != Human {
		t.Errorf("input '1' should select Human, got %v", playerType)
	}
}

func TestReadPlayerType_SelectsAIWith2(t *testing.T) {
	var output bytes.Buffer
	reader := bufio.NewReader(strings.NewReader("2\n"))

	playerType, err := ReadPlayerType(reader, &output)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if playerType != AI {
		t.Errorf("input '2' should select AI, got %v", playerType)
	}
}

func TestReadPlayerType_RetriesAfterInvalidNumber(t *testing.T) {
	var output bytes.Buffer
	reader := bufio.NewReader(strings.NewReader("3\n1\n"))

	playerType, err := ReadPlayerType(reader, &output)

	if err != nil {
		t.Fatalf("should eventually succeed: %v", err)
	}

	if playerType != Human {
		t.Errorf("should accept '1' after retry, got %v", playerType)
	}

	result := output.String()
	if !strings.Contains(result, "Invalid input") {
		t.Error("should show error message for invalid input")
	}
}

func TestReadPlayerType_RetriesAfterText(t *testing.T) {
	var output bytes.Buffer
	reader := bufio.NewReader(strings.NewReader("abc\n2\n"))

	playerType, err := ReadPlayerType(reader, &output)

	if err != nil {
		t.Fatalf("should eventually succeed: %v", err)
	}

	if playerType != AI {
		t.Errorf("should accept '2' after retry, got %v", playerType)
	}

	result := output.String()
	if !strings.Contains(result, "Invalid input") {
		t.Error("should show error message for text input")
	}
}

func TestReadPlayerType_RetriesAfterEmptyInput(t *testing.T) {
	var output bytes.Buffer
	reader := bufio.NewReader(strings.NewReader("\n1\n"))

	playerType, err := ReadPlayerType(reader, &output)

	if err != nil {
		t.Fatalf("should eventually succeed: %v", err)
	}

	if playerType != Human {
		t.Errorf("should accept '1' after retry, got %v", playerType)
	}

	result := output.String()
	if !strings.Contains(result, "Invalid input") {
		t.Error("should show error message for empty input")
	}
}

func TestReadPlayerType_HandlesMultipleInvalidInputs(t *testing.T) {
	var output bytes.Buffer
	reader := bufio.NewReader(strings.NewReader("5\nabc\n\n2\n"))

	playerType, err := ReadPlayerType(reader, &output)

	if err != nil {
		t.Fatalf("should eventually succeed: %v", err)
	}

	if playerType != AI {
		t.Errorf("should accept '2' after multiple retries, got %v", playerType)
	}

	result := output.String()
	if !strings.Contains(result, "Invalid input") {
		t.Error("should show error messages")
	}
}

func TestReadPlayAgain_AcceptsLowercaseY(t *testing.T) {
	var output bytes.Buffer
	reader := bufio.NewReader(strings.NewReader("y\n"))

	playAgain, err := ReadPlayAgain(reader, &output)

	if err != nil {
		t.Fatalf("should accept 'y': %v", err)
	}

	if !playAgain {
		t.Error("'y' should return true for play again")
	}
}

func TestReadPlayAgain_AcceptsUppercaseY(t *testing.T) {
	var output bytes.Buffer
	reader := bufio.NewReader(strings.NewReader("Y\n"))

	playAgain, err := ReadPlayAgain(reader, &output)

	if err != nil {
		t.Fatalf("should accept 'Y': %v", err)
	}

	if !playAgain {
		t.Error("'Y' should return true for play again")
	}
}

func TestReadPlayAgain_AcceptsLowercaseYes(t *testing.T) {
	var output bytes.Buffer
	reader := bufio.NewReader(strings.NewReader("yes\n"))

	playAgain, err := ReadPlayAgain(reader, &output)

	if err != nil {
		t.Fatalf("should accept 'yes': %v", err)
	}

	if !playAgain {
		t.Error("'yes' should return true for play again")
	}
}

func TestReadPlayAgain_AcceptsUppercaseYES(t *testing.T) {
	var output bytes.Buffer
	reader := bufio.NewReader(strings.NewReader("YES\n"))

	playAgain, err := ReadPlayAgain(reader, &output)

	if err != nil {
		t.Fatalf("should accept 'YES': %v", err)
	}

	if !playAgain {
		t.Error("'YES' should return true for play again")
	}
}

func TestReadPlayAgain_AcceptsLowercaseN(t *testing.T) {
	var output bytes.Buffer
	reader := bufio.NewReader(strings.NewReader("n\n"))

	playAgain, err := ReadPlayAgain(reader, &output)

	if err != nil {
		t.Fatalf("should accept 'n': %v", err)
	}

	if playAgain {
		t.Error("'n' should return false for play again")
	}
}

func TestReadPlayAgain_AcceptsUppercaseN(t *testing.T) {
	var output bytes.Buffer
	reader := bufio.NewReader(strings.NewReader("N\n"))

	playAgain, err := ReadPlayAgain(reader, &output)

	if err != nil {
		t.Fatalf("should accept 'N': %v", err)
	}

	if playAgain {
		t.Error("'N' should return false for play again")
	}
}

func TestReadPlayAgain_AcceptsLowercaseNo(t *testing.T) {
	var output bytes.Buffer
	reader := bufio.NewReader(strings.NewReader("no\n"))

	playAgain, err := ReadPlayAgain(reader, &output)

	if err != nil {
		t.Fatalf("should accept 'no': %v", err)
	}

	if playAgain {
		t.Error("'no' should return false for play again")
	}
}

func TestReadPlayAgain_AcceptsUppercaseNO(t *testing.T) {
	var output bytes.Buffer
	reader := bufio.NewReader(strings.NewReader("NO\n"))

	playAgain, err := ReadPlayAgain(reader, &output)

	if err != nil {
		t.Fatalf("should accept 'NO': %v", err)
	}

	if playAgain {
		t.Error("'NO' should return false for play again")
	}
}

func TestReadPlayAgain_RetriesAfterMaybeAcceptsYes(t *testing.T) {
	var output bytes.Buffer
	reader := bufio.NewReader(strings.NewReader("maybe\ny\n"))

	playAgain, err := ReadPlayAgain(reader, &output)

	if err != nil {
		t.Fatalf("should eventually succeed: %v", err)
	}

	if !playAgain {
		t.Error("should accept 'y' after retry")
	}

	result := output.String()
	if !strings.Contains(result, "Invalid input") {
		t.Error("should show error message")
	}
}

func TestReadPlayAgain_RetriesAfterInvalidAcceptsNo(t *testing.T) {
	var output bytes.Buffer
	reader := bufio.NewReader(strings.NewReader("invalid\nn\n"))

	playAgain, err := ReadPlayAgain(reader, &output)

	if err != nil {
		t.Fatalf("should eventually succeed: %v", err)
	}

	if playAgain {
		t.Error("should accept 'n' after retry")
	}

	result := output.String()
	if !strings.Contains(result, "Invalid input") {
		t.Error("should show error message")
	}
}

func TestReadPlayAgain_RetriesAfterEmptyAcceptsYes(t *testing.T) {
	var output bytes.Buffer
	reader := bufio.NewReader(strings.NewReader("\ny\n"))

	playAgain, err := ReadPlayAgain(reader, &output)

	if err != nil {
		t.Fatalf("should eventually succeed: %v", err)
	}

	if !playAgain {
		t.Error("should accept 'y' after empty retry")
	}

	result := output.String()
	if !strings.Contains(result, "Invalid input") {
		t.Error("should show error message for empty")
	}
}

func TestReadPlayAgain_HandlesMultipleInvalidInputs(t *testing.T) {
	var output bytes.Buffer
	reader := bufio.NewReader(strings.NewReader("maybe\nwhat\n\nn\n"))

	playAgain, err := ReadPlayAgain(reader, &output)

	if err != nil {
		t.Fatalf("should eventually succeed: %v", err)
	}

	if playAgain {
		t.Error("should accept 'n' after multiple retries")
	}

	result := output.String()
	if !strings.Contains(result, "Invalid input") {
		t.Error("should show error messages")
	}
}

func TestInputOutput_PlayerSelectionFlow(t *testing.T) {
	input := "1\n"
	var output bytes.Buffer
	reader := bufio.NewReader(strings.NewReader(input))

	ShowPlayerTypeSelection(&output, "X")
	playerType, _ := ReadPlayerType(reader, &output)

	result := output.String()

	if !strings.Contains(result, "Human") || !strings.Contains(result, "AI") {
		t.Error("should display player type options")
	}

	if playerType != Human {
		t.Error("should accept human player selection")
	}
}

func TestInputOutput_PlayAgainFlow(t *testing.T) {
	input := "y\n"
	var output bytes.Buffer
	reader := bufio.NewReader(strings.NewReader(input))

	ShowPlayAgainPrompt(&output)
	playAgain, _ := ReadPlayAgain(reader, &output)

	result := output.String()

	if !strings.Contains(result, "Play again") {
		t.Error("should display play again prompt")
	}

	if !playAgain {
		t.Error("should accept yes response")
	}
}
