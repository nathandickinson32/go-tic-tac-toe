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

func TestBoardConditions(t *testing.T) {

	t.Run("getting all positions", func(t *testing.T) {

		t.Run("gets all 9 board positions", func(t *testing.T) {
			board := initBoard()

			expected := [3][3]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			}

			assertBoardEquals(t, board, expected)
		})
	})

	t.Run("getting available moves", func(t *testing.T) {

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
	})

	t.Run("make-move", func(t *testing.T) {

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
	})

	t.Run("display-board", func(t *testing.T) {

		t.Run("empty board", func(t *testing.T) {
			board := initBoard()

			want := " 1 | 2 | 3 \n" +
				"-----------\n" +
				" 4 | 5 | 6 \n" +
				"-----------\n" +
				" 7 | 8 | 9 "

			got := displayBoard(board)

			if got != want {
				t.Errorf("Board display mismatch.\nGot:\n%s\n\nWant:\n%s", got, want)
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

			got := displayBoard(board)

			if got != want {
				t.Errorf("Board display mismatch.\nGot:\n%s\n\nWant:\n%s", got, want)
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

			got := displayBoard(board)

			if got != want {
				t.Errorf("Board display mismatch")
			}
		})
	})
}
