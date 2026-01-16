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

func TestBoard_Initialize(t *testing.T) {
	t.Run("gets all 9 board positions", func(t *testing.T) {
		board := NewBoard()

		expected := [3][3]string{
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		}

		assertBoardEquals(t, board, expected)
	})
}

func TestBoard_AvailableMoves(t *testing.T) {
	t.Run("gets one available move", func(t *testing.T) {
		testBoard := Board{
			{"O", "X", "O"},
			{"X", "X", "O"},
			{"X", "O", "9"},
		}

		moves := testBoard.AvailableMoves()

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

		moves := testBoard.AvailableMoves()

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

		moves := testBoard.AvailableMoves()

		takenPositions := []int{1, 2, 3, 5, 7, 8}
		for _, taken := range takenPositions {
			for _, move := range moves {
				if move == taken {
					t.Errorf("Position %d should not be available (it's taken)", taken)
				}
			}
		}
	})

	t.Run("empty board returns all moves", func(t *testing.T) {
		board := NewBoard()
		moves := board.AvailableMoves()

		if len(moves) != 9 {
			t.Errorf("Expected 9 available moves on empty board, got %d", len(moves))
		}
	})
}

func TestBoard_MakeMove(t *testing.T) {
	t.Run("invalid position returns error", func(t *testing.T) {
		board := NewBoard()

		if err := board.MakeMove(0, "X"); err == nil {
			t.Error("position 0 should return error")
		}

		if err := board.MakeMove(10, "X"); err == nil {
			t.Error("position 10 should return error")
		}

		if err := board.MakeMove(-1, "X"); err == nil {
			t.Error("negative position should return error")
		}
	})

	t.Run("occupied position returns error", func(t *testing.T) {
		board := NewBoard()
		board.MakeMove(5, "X")

		if err := board.MakeMove(5, "O"); err == nil {
			t.Error("occupied position should return error")
		}
	})

	t.Run("marks with X", func(t *testing.T) {
		board := NewBoard()

		err := board.MakeMove(5, "X")

		if err != nil {
			t.Errorf("valid move should not return error: %v", err)
		}

		expectedBoard := Board{
			{"1", "2", "3"},
			{"4", "X", "6"},
			{"7", "8", "9"},
		}

		assertBoardEquals(t, board, expectedBoard)
	})

	t.Run("marks with O", func(t *testing.T) {
		board := NewBoard()

		err := board.MakeMove(1, "O")

		if err != nil {
			t.Errorf("valid move should not return error: %v", err)
		}

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
			{2, "O", 0, 1},
			{7, "X", 2, 0},
		}

		for _, tt := range tests {
			board := NewBoard()
			err := board.MakeMove(tt.position, tt.player)
			if err != nil {
				t.Errorf("Position %d: unexpected error: %v", tt.position, err)
			}

			if board[tt.row][tt.col] != tt.player {
				t.Errorf("Position %d: expected %s at [%d][%d], got %s",
					tt.position, tt.player, tt.row, tt.col, board[tt.row][tt.col])
			}
		}
	})

	t.Run("multiple moves", func(t *testing.T) {
		board := NewBoard()

		moves := []struct {
			position int
			player   string
		}{
			{1, "X"},
			{2, "O"},
			{5, "X"},
			{9, "O"},
		}

		for _, move := range moves {
			if err := board.MakeMove(move.position, move.player); err != nil {
				t.Errorf("move %d for %s should be valid: %v", move.position, move.player, err)
			}
		}

		expected := Board{
			{"X", "O", "3"},
			{"4", "X", "6"},
			{"7", "8", "O"},
		}

		assertBoardEquals(t, board, expected)
	})
}
