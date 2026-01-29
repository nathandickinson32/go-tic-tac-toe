package players

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
	"ttt/boards"
)

func TestHumanPlayer_ValidPosition1(t *testing.T) {
	board := boards.NewBoard()
	input := "1\n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	got, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("should accept valid input: %v", err)
	}

	if got != 1 {
		t.Errorf("got position %d, want 1", got)
	}
}

func TestHumanPlayer_ValidPosition5(t *testing.T) {
	board := boards.NewBoard()
	input := "5\n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	got, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("should accept valid input: %v", err)
	}

	if got != 5 {
		t.Errorf("got position %d, want 5", got)
	}
}

func TestHumanPlayer_ValidPosition9(t *testing.T) {
	board := boards.NewBoard()
	input := "9\n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	got, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("should accept valid input: %v", err)
	}

	if got != 9 {
		t.Errorf("got position %d, want 9", got)
	}
}

func TestHumanPlayer_InputWithLeadingWhitespace(t *testing.T) {
	board := boards.NewBoard()
	input := "  7\n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	got, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("should accept input with leading whitespace: %v", err)
	}

	if got != 7 {
		t.Errorf("got position %d, want 7", got)
	}
}

func TestHumanPlayer_InputWithTrailingWhitespace(t *testing.T) {
	board := boards.NewBoard()
	input := "3  \n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	got, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("should accept input with trailing whitespace: %v", err)
	}

	if got != 3 {
		t.Errorf("got position %d, want 3", got)
	}
}

func TestHumanPlayer_InputWithSurroundingWhitespace(t *testing.T) {
	board := boards.NewBoard()
	input := "  4  \n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	got, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("should accept input with surrounding whitespace: %v", err)
	}

	if got != 4 {
		t.Errorf("got position %d, want 4", got)
	}
}

func TestHumanPlayer_RetriesAfterEmptyInput(t *testing.T) {
	board := boards.NewBoard()
	input := "\n\n5\n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	got, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("should eventually accept valid input: %v", err)
	}

	if got != 5 {
		t.Errorf("got position %d, want 5", got)
	}

	result := output.String()
	if !strings.Contains(result, "Invalid input") {
		t.Error("should show error message for empty input")
	}
}

func TestHumanPlayer_RetriesAfterNonNumericInput(t *testing.T) {
	board := boards.NewBoard()
	input := "abc\n5\n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	got, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("should eventually accept valid input: %v", err)
	}

	if got != 5 {
		t.Errorf("got position %d, want 5", got)
	}

	result := output.String()
	if !strings.Contains(result, "Invalid input") {
		t.Error("should show error message for non-numeric input")
	}
}

func TestHumanPlayer_RetriesAfterPositionZero(t *testing.T) {
	board := boards.NewBoard()
	input := "0\n5\n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	got, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("should eventually accept valid input: %v", err)
	}

	if got != 5 {
		t.Errorf("got position %d, want 5", got)
	}

	result := output.String()
	if !strings.Contains(result, "Invalid input") {
		t.Error("should reject zero as out of range")
	}
}

func TestHumanPlayer_RetriesAfterPositionTooHigh(t *testing.T) {
	board := boards.NewBoard()
	input := "10\n5\n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	got, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("should eventually accept valid input: %v", err)
	}

	if got != 5 {
		t.Errorf("got position %d, want 5", got)
	}

	result := output.String()
	if !strings.Contains(result, "Invalid input") {
		t.Error("should reject positions above 9")
	}
}

func TestHumanPlayer_RetriesAfterNegativePosition(t *testing.T) {
	board := boards.NewBoard()
	input := "-1\n5\n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	got, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("should eventually accept valid input: %v", err)
	}

	if got != 5 {
		t.Errorf("got position %d, want 5", got)
	}

	result := output.String()
	if !strings.Contains(result, "Invalid input") {
		t.Error("should reject negative positions")
	}
}

func TestHumanPlayer_RetriesAfterDecimalNumber(t *testing.T) {
	board := boards.NewBoard()
	input := "5.5\n5\n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	got, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("should eventually accept valid input: %v", err)
	}

	if got != 5 {
		t.Errorf("got position %d, want 5", got)
	}

	result := output.String()
	if !strings.Contains(result, "Invalid input") {
		t.Error("should reject decimal numbers")
	}
}

func TestHumanPlayer_HandlesMultipleDifferentInvalidInputs(t *testing.T) {
	board := boards.NewBoard()
	input := "abc\n0\n10\n5\n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	got, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("should eventually accept valid input: %v", err)
	}

	if got != 5 {
		t.Errorf("got position %d, want 5", got)
	}

	result := output.String()
	if !strings.Contains(result, "Invalid input") {
		t.Error("should show error messages for multiple invalid inputs")
	}
}

func TestHumanPlayer_RejectsOccupiedPositionByX(t *testing.T) {
	board := boards.NewBoard()
	board.MakeMove(5, "X")
	input := "5\n7\n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	got, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("should eventually accept valid move: %v", err)
	}

	if got != 7 {
		t.Errorf("should accept valid position 7, got %d", got)
	}

	result := output.String()
	if !strings.Contains(result, "Position already taken") {
		t.Error("should show rejection message for occupied position")
	}
}

func TestHumanPlayer_RejectsOccupiedPositionByO(t *testing.T) {
	board := boards.NewBoard()
	board.MakeMove(1, "O")
	input := "1\n2\n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	got, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("should eventually accept valid move: %v", err)
	}

	if got != 2 {
		t.Errorf("should accept valid position 2, got %d", got)
	}

	result := output.String()
	if !strings.Contains(result, "Position already taken") {
		t.Error("should show rejection message for occupied position")
	}
}

func TestHumanPlayer_ValidMove(t *testing.T) {
	board := boards.NewBoard()
	board.MakeMove(5, "X")
	input := "6\n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	got, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("should accept adjacent position: %v", err)
	}

	if got != 6 {
		t.Errorf("should accept position 6 adjacent to occupied 5, got %d", got)
	}
}

func TestHumanPlayer_DisplaysPromptToUser(t *testing.T) {
	board := boards.NewBoard()
	input := "5\n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	_, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	result := output.String()
	if !strings.Contains(result, "Enter your move") {
		t.Error("should prompt user to enter move")
	}
}

func TestHumanPlayer_DisplaysPromptOnEachRetry(t *testing.T) {
	board := boards.NewBoard()
	input := "abc\n10\n5\n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	_, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	result := output.String()
	promptCount := strings.Count(result, "Enter your move")

	if promptCount != 3 {
		t.Errorf("should show prompt for each attempt (3), got %d prompts", promptCount)
	}
}

func TestHumanPlayer_HandlesMultipleBadInput(t *testing.T) {
	board := boards.NewBoard()
	board.MakeMove(1, "X")
	input := "\nabc\n10\n1\n5\n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	got, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("should eventually accept valid move: %v", err)
	}

	if got != 5 {
		t.Errorf("got position %d, want 5", got)
	}

	result := output.String()
	if !strings.Contains(result, "Invalid input") {
		t.Error("should show invalid input errors")
	}
	if !strings.Contains(result, "Position already taken") {
		t.Error("should show occupied position error")
	}
}

func TestHumanPlayer_HandlesMultipleOccupiedAttempts(t *testing.T) {
	board := boards.NewBoard()
	board.MakeMove(1, "X")
	board.MakeMove(2, "O")
	board.MakeMove(3, "X")
	input := "1\n2\n3\n5\n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	got, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("should eventually accept valid move: %v", err)
	}

	if got != 5 {
		t.Errorf("got position %d, want 5", got)
	}

	result := output.String()
	rejectionCount := strings.Count(result, "Position already taken")
	if rejectionCount != 3 {
		t.Errorf("should show 3 rejection messages, got %d", rejectionCount)
	}
}

func TestHumanPlayer_HandlesAlternatingInvalidAndOccupied(t *testing.T) {
	board := boards.NewBoard()
	board.MakeMove(5, "X")
	input := "abc\n5\n0\n7\n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	got, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("should eventually accept valid move: %v", err)
	}

	if got != 7 {
		t.Errorf("got position %d, want 7", got)
	}

	result := output.String()
	if !strings.Contains(result, "Invalid input") {
		t.Error("should show invalid input errors")
	}
	if !strings.Contains(result, "Position already taken") {
		t.Error("should show occupied position error")
	}
}

func TestHumanPlayer_AcceptsLastAvailablePosition(t *testing.T) {
	board := boards.Board{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"X", "O", "9"},
	}
	input := "9\n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	got, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("should accept only available move: %v", err)
	}

	if got != 9 {
		t.Errorf("should accept position 9, got %d", got)
	}
}

func TestHumanPlayer_AllOccupiedPositionsBeforeAcceptingValid(t *testing.T) {
	board := boards.Board{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"X", "O", "9"},
	}
	input := "1\n2\n3\n4\n5\n6\n7\n8\n9\n"
	var output bytes.Buffer
	human := NewHumanPlayer(bufio.NewReader(strings.NewReader(input)), &output)

	got, err := human.ReadMove(board)

	if err != nil {
		t.Fatalf("should eventually accept valid position: %v", err)
	}

	if got != 9 {
		t.Errorf("should accept position 9, got %d", got)
	}

	result := output.String()
	rejectionCount := strings.Count(result, "Position already taken")
	if rejectionCount != 8 {
		t.Errorf("should reject 8 occupied positions, got %d rejections", rejectionCount)
	}
}
