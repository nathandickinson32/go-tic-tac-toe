package boards

import (
	"testing"
)

func assertBoardEquals(t *testing.T, got, want Board, context string) {
	t.Helper()
	for row := range 3 {
		for col := range 3 {
			if got[row][col] != want[row][col] {
				t.Errorf("%s: board[%d][%d] = %s, want %s",
					context, row, col, got[row][col], want[row][col])
			}
		}
	}
}

func TestBoard_InitialBoard(t *testing.T) {
	board := NewBoard()

	expected := Board{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	assertBoardEquals(t, board, expected, "new board")
}

func TestBoard_MakeMove(t *testing.T) {
	board := NewBoard()

	err := board.MakeMove(1, "X")

	if err != nil {
		t.Fatalf("position 1 should accept move: %v", err)
	}

	if board[0][0] != "X" {
		t.Errorf("expected X at [0][0], got %s", board[0][0])
	}
}

func TestBoard_MoveToCenter(t *testing.T) {
	board := NewBoard()

	err := board.MakeMove(5, "O")

	if err != nil {
		t.Fatalf("position 5 should accept move: %v", err)
	}

	if board[1][1] != "O" {
		t.Errorf("expected O at [1][1], got %s", board[1][1])
	}
}

func TestBoard_MoveToBottomRight(t *testing.T) {
	board := NewBoard()

	err := board.MakeMove(9, "X")

	if err != nil {
		t.Fatalf("position 9 should accept move: %v", err)
	}

	if board[2][2] != "X" {
		t.Errorf("expected X at [2][2], got %s", board[2][2])
	}
}

func TestBoard_CompleteGameSequence(t *testing.T) {
	board := NewBoard()

	board.MakeMove(1, "X")
	board.MakeMove(2, "O")
	board.MakeMove(5, "X")
	board.MakeMove(9, "O")

	expected := Board{
		{"X", "O", "3"},
		{"4", "X", "6"},
		{"7", "8", "O"},
	}

	assertBoardEquals(t, board, expected, "after move sequence")
}

func TestBoard_Position0(t *testing.T) {
	board := NewBoard()

	err := board.MakeMove(0, "X")

	if err == nil {
		t.Error("should return error for position 0")
	}
}

func TestBoard_Position10(t *testing.T) {
	board := NewBoard()

	err := board.MakeMove(10, "X")

	if err == nil {
		t.Error("should return error for position 10")
	}
}

func TestBoard_NegativePosition(t *testing.T) {
	board := NewBoard()

	err := board.MakeMove(-1, "X")

	if err == nil {
		t.Error("should return error for negative position")
	}
}

func TestBoard_OccupiedByX(t *testing.T) {
	board := NewBoard()
	board.MakeMove(5, "X")

	err := board.MakeMove(5, "O")

	if err == nil {
		t.Error("should return error for position occupied by X")
	}
}

func TestBoard_OccupiedByO(t *testing.T) {
	board := NewBoard()
	board.MakeMove(3, "O")

	err := board.MakeMove(3, "X")

	if err == nil {
		t.Error("should return error for position occupied by O")
	}
}

func TestBoard_NoWinnerOnEmptyBoard(t *testing.T) {
	board := NewBoard()

	winner := board.CheckWinner()

	if winner != "" {
		t.Errorf("empty board should have no winner, got %q", winner)
	}
}

func TestBoard_XWinTopRow(t *testing.T) {
	board := Board{
		{"X", "X", "X"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	winner := board.CheckWinner()

	if winner != "X" {
		t.Errorf("should be X winner on top row, got %q", winner)
	}
}

func TestBoard_OWinMiddleRow(t *testing.T) {
	board := Board{
		{"1", "2", "3"},
		{"O", "O", "O"},
		{"7", "8", "9"},
	}

	winner := board.CheckWinner()

	if winner != "O" {
		t.Errorf("should be O winner on middle row, got %q", winner)
	}
}

func TestBoard_OWinBottomRow(t *testing.T) {
	board := Board{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"O", "O", "O"},
	}

	winner := board.CheckWinner()

	if winner != "O" {
		t.Errorf("should be O winner on bottom row, got %q", winner)
	}
}

func TestBoard_OWinLeftColumn(t *testing.T) {
	board := Board{
		{"O", "2", "3"},
		{"O", "5", "6"},
		{"O", "8", "9"},
	}

	winner := board.CheckWinner()

	if winner != "O" {
		t.Errorf("should be O winner on left column, got %q", winner)
	}
}

func TestBoard_XWinMiddleColumn(t *testing.T) {
	board := Board{
		{"1", "X", "3"},
		{"4", "X", "6"},
		{"7", "X", "9"},
	}

	winner := board.CheckWinner()

	if winner != "X" {
		t.Errorf("should be X winner on middle column, got %q", winner)
	}
}

func TestBoard_XWinRightColumn(t *testing.T) {
	board := Board{
		{"1", "2", "X"},
		{"4", "5", "X"},
		{"7", "8", "X"},
	}

	winner := board.CheckWinner()

	if winner != "X" {
		t.Errorf("should be X winner on right column, got %q", winner)
	}
}

func TestBoard_XWinMainDiagonal(t *testing.T) {
	board := Board{
		{"X", "2", "3"},
		{"4", "X", "6"},
		{"7", "8", "X"},
	}

	winner := board.CheckWinner()

	if winner != "X" {
		t.Errorf("should be X winner on down right diagonal, got %q", winner)
	}
}

func TestBoard_OWinAntiDiagonal(t *testing.T) {
	board := Board{
		{"1", "2", "O"},
		{"4", "O", "6"},
		{"O", "8", "9"},
	}

	winner := board.CheckWinner()

	if winner != "O" {
		t.Errorf("should be O winner on down left diagonal, got %q", winner)
	}
}

func TestBoard_NoWinner(t *testing.T) {
	board := Board{
		{"X", "O", "3"},
		{"4", "X", "6"},
		{"O", "8", "9"},
	}

	winner := board.CheckWinner()

	if winner != "" {
		t.Errorf("partial board should have no winner, got %q", winner)
	}
}

func TestBoard_InProgress(t *testing.T) {
	board := NewBoard()

	status := board.GetGameStatus()

	if status != InProgress {
		t.Errorf("new board should be in progress, got %v", status)
	}
}

func TestBoard_XWins(t *testing.T) {
	board := Board{
		{"X", "X", "X"},
		{"O", "O", "6"},
		{"7", "8", "9"},
	}

	status := board.GetGameStatus()

	if status != XWins {
		t.Errorf("should show XWins status, got %v", status)
	}
}

func TestBoard_OWins(t *testing.T) {
	board := Board{
		{"O", "O", "O"},
		{"X", "X", "6"},
		{"7", "8", "9"},
	}

	status := board.GetGameStatus()

	if status != OWins {
		t.Errorf("should show OWins status, got %v", status)
	}
}

func TestBoard_DrawOnFullBoard(t *testing.T) {
	board := Board{
		{"X", "O", "X"},
		{"X", "O", "O"},
		{"O", "X", "X"},
	}

	status := board.GetGameStatus()

	if status != Draw {
		t.Errorf("full board with no winner should be draw, got %v", status)
	}
}

func TestBoard_FullBoardMove(t *testing.T) {
	board := Board{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"X", "O", "9"},
	}

	err := board.MakeMove(9, "X")

	if err != nil {
		t.Errorf("should accept move to last position: %v", err)
	}

	if board[2][2] != "X" {
		t.Errorf("position 9 should have X, got %s", board[2][2])
	}
}

func TestBoard_NoMovesAvailableAfterLastMove(t *testing.T) {
	board := Board{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"X", "O", "9"},
	}
	board.MakeMove(9, "X")

	moves := board.AvailableMoves()

	if len(moves) != 0 {
		t.Errorf("should have no available moves on full board, got %d", len(moves))
	}
}
