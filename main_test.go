package main

import (
	"slices"
	"testing"
)

func TestBoardConditions(t *testing.T) {

	t.Run("getting all positions", func(t *testing.T) {

		t.Run("gets all 9 board positions for 3x3 board", func(t *testing.T) {
			board := initBoard()

			expected := [3][3]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			}

			for i := range 3 {
				for j := range 3 {
					if board[i][j] != expected[i][j] {
						t.Errorf("Expected board[%d][%d] to be '%s', got '%s'",
							i, j, expected[i][j], board[i][j])
					}
				}
			}
		})
	})

	t.Run("getting available 3x3 moves", func(t *testing.T) {

		t.Run("gets one available move 3x3", func(t *testing.T) {
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

		t.Run("gets all available moves for 3x3", func(t *testing.T) {
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

		t.Run("does not contain moves that are taken for 3x3", func(t *testing.T) {
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
	})

}
