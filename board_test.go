package main

import (
	"slices"
	"testing"
)

func assertBoardEquals(t *testing.T, got, want Board) {
	t.Helper()
	for row := range 3 {
		for col := range 3 {
			if got[row][col] != want[row][col] {
				t.Errorf("board[%d][%d] = %s, want %s",
					row, col, got[row][col], want[row][col])
			}
		}
	}
}

func TestInitBoard(t *testing.T) {
	t.Run("gets all 9 board positions", func(t *testing.T) {
		board := initBoard()

		expected := [3][3]string{
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		}

		assertBoardEquals(t, board, expected)
	})
}

func TestGetAvailableMoves(t *testing.T) {
	t.Run("gets one available move", func(t *testing.T) {
		testBoard := Board{
			{"O", "X", "O"},
			{"X", "X", "O"},
			{"X", "O", "9"},
		}

		moves := getAvailableMoves(testBoard)

		if len(moves) != 1 {
			t.Fatalf("Expected 1 available move, got %d", len(moves))
		}

		if moves[0] != 9 {
			t.Errorf("Expected move 9, got %d", moves[0])
		}
	})

	t.Run("gets all available moves", func(t *testing.T) {
		testBoard := Board{
			{"O", "X", "O"},
			{"4", "X", "6"},
			{"X", "O", "9"},
		}

		moves := getAvailableMoves(testBoard)

		if len(moves) != 3 {
			t.Fatalf("Expected 3 available moves, got %d", len(moves))
		}

		expectedMoves := []int{4, 6, 9}
		for _, expected := range expectedMoves {
			found := slices.Contains(moves, expected)
			if !found {
				t.Errorf("Expected move %d to be available", expected)
			}
		}
	})

	t.Run("does not contain moves that are taken", func(t *testing.T) {
		testBoard := Board{
			{"O", "X", "O"},
			{"4", "X", "6"},
			{"X", "O", "9"},
		}

		moves := getAvailableMoves(testBoard)

		takenPositions := []int{1, 2, 3, 5, 7, 8}
		for _, taken := range takenPositions {
			for _, move := range moves {
				if move == taken {
					t.Errorf("Position %d should not be available (it's taken)", taken)
				}
			}
		}
	})
}

func TestMakeMove(t *testing.T) {
	t.Run("invalid position", func(t *testing.T) {
		board := initBoard()

		if makeMove(&board, 0, "X") {
			t.Error("position 0")
		}

		if makeMove(&board, 10, "X") {
			t.Error("position 10")
		}

		if makeMove(&board, -1, "X") {
			t.Error("negative position")
		}
	})

	t.Run("occupied position", func(t *testing.T) {
		board := initBoard()
		makeMove(&board, 5, "X")

		if makeMove(&board, 5, "O") {
			t.Error("occupied position")
		}
	})

	t.Run("marks with X", func(t *testing.T) {
		board := initBoard()

		makeMove(&board, 5, "X")

		expectedBoard := Board{
			{"1", "2", "3"},
			{"4", "X", "6"},
			{"7", "8", "9"},
		}

		assertBoardEquals(t, board, expectedBoard)
	})

	t.Run("marks with O", func(t *testing.T) {
		board := initBoard()

		makeMove(&board, 1, "O")

		expectedBoard := Board{
			{"O", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		}

		assertBoardEquals(t, board, expectedBoard)
	})

	t.Run("marks correct positions", func(t *testing.T) {
		tests := []struct {
			position int
			player   string
			row      int
			col      int
		}{
			{1, "X", 0, 0},
			{5, "O", 1, 1},
			{9, "X", 2, 2},
		}

		for _, tt := range tests {
			board := initBoard()
			makeMove(&board, tt.position, tt.player)

			if board[tt.row][tt.col] != tt.player {
				t.Errorf("Position %d: expected %s at [%d][%d], got %s",
					tt.position, tt.player, tt.row, tt.col, board[tt.row][tt.col])
			}
		}
	})
}

func TestCheckWinner(t *testing.T) {

	t.Run("empty board", func(t *testing.T) {
		board := initBoard()

		got := checkWinner(board)

		if got != "" {
			t.Fatalf("expected no winner, got %q", got)
		}
	})

	t.Run("row win for X", func(t *testing.T) {
		board := Board{
			{"X", "X", "X"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		}

		got := checkWinner(board)

		if got != "X" {
			t.Fatalf("got %q, want X", got)
		}
	})

	t.Run("row win for O", func(t *testing.T) {
		board := Board{
			{"1", "2", "3"},
			{"O", "O", "O"},
			{"7", "8", "9"},
		}

		got := checkWinner(board)

		if got != "O" {
			t.Fatalf("got %q, want O", got)
		}
	})

	t.Run("third row win", func(t *testing.T) {
		board := Board{
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"O", "O", "O"},
		}

		got := checkWinner(board)

		if got != "O" {
			t.Fatalf("got %q, want O", got)
		}
	})

	t.Run("column win for X", func(t *testing.T) {
		board := Board{
			{"1", "X", "3"},
			{"4", "X", "6"},
			{"7", "X", "9"},
		}

		got := checkWinner(board)

		if got != "X" {
			t.Fatalf("got %q, want X", got)
		}
	})

	t.Run("column win for O", func(t *testing.T) {
		board := Board{
			{"O", "2", "3"},
			{"O", "5", "6"},
			{"O", "8", "9"},
		}

		got := checkWinner(board)

		if got != "O" {
			t.Fatalf("got %q, want O", got)
		}
	})

	t.Run("third column win", func(t *testing.T) {
		board := Board{
			{"1", "2", "X"},
			{"4", "5", "X"},
			{"7", "8", "X"},
		}

		got := checkWinner(board)

		if got != "X" {
			t.Fatalf("got %q, want X", got)
		}
	})

	t.Run("diagonal win down right", func(t *testing.T) {
		board := Board{
			{"X", "2", "3"},
			{"4", "X", "6"},
			{"7", "8", "X"},
		}

		got := checkWinner(board)

		if got != "X" {
			t.Fatalf("got %q, want X", got)
		}
	})

	t.Run("diagonal win down left", func(t *testing.T) {
		board := Board{
			{"1", "2", "O"},
			{"4", "O", "6"},
			{"O", "8", "9"},
		}

		got := checkWinner(board)

		if got != "O" {
			t.Fatalf("got %q, want O", got)
		}
	})

	t.Run("full board no winner", func(t *testing.T) {
		board := Board{
			{"X", "O", "X"},
			{"X", "O", "O"},
			{"O", "X", "X"},
		}

		got := checkWinner(board)

		if got != "" {
			t.Fatalf("expected no winner, got %q", got)
		}
	})
}
