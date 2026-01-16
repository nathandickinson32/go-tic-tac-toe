package main

import "testing"

func TestGameRules_CheckWinner(t *testing.T) {
	rules := NewGameRules()

	t.Run("empty board", func(t *testing.T) {
		board := NewBoard()

		got := rules.CheckWinner(board)

		if got != "" {
			t.Fatalf("expected no winner, got %q", got)
		}
	})

	t.Run("partial board no winner", func(t *testing.T) {
		board := Board{
			{"X", "O", "3"},
			{"4", "X", "6"},
			{"O", "8", "9"},
		}

		got := rules.CheckWinner(board)

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

		got := rules.CheckWinner(board)

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

		got := rules.CheckWinner(board)

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

		got := rules.CheckWinner(board)

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

		got := rules.CheckWinner(board)

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

		got := rules.CheckWinner(board)

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

		got := rules.CheckWinner(board)

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

		got := rules.CheckWinner(board)

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

		got := rules.CheckWinner(board)

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

		got := rules.CheckWinner(board)

		if got != "" {
			t.Fatalf("expected no winner, got %q", got)
		}
	})
}

func TestGameRules_GetGameStatus(t *testing.T) {
	rules := NewGameRules()

	t.Run("in progress empty board", func(t *testing.T) {
		board := NewBoard()

		status := rules.GetGameStatus(board)

		if status != InProgress {
			t.Errorf("expected InProgress, got %v", status)
		}
	})

	t.Run("in progress partial board", func(t *testing.T) {
		board := Board{
			{"X", "O", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		}

		status := rules.GetGameStatus(board)

		if status != InProgress {
			t.Errorf("expected InProgress, got %v", status)
		}
	})

	t.Run("X wins", func(t *testing.T) {
		board := Board{
			{"X", "X", "X"},
			{"O", "O", "6"},
			{"7", "8", "9"},
		}

		status := rules.GetGameStatus(board)

		if status != XWins {
			t.Errorf("expected XWins, got %v", status)
		}
	})

	t.Run("O wins", func(t *testing.T) {
		board := Board{
			{"O", "O", "O"},
			{"X", "X", "6"},
			{"7", "8", "9"},
		}

		status := rules.GetGameStatus(board)

		if status != OWins {
			t.Errorf("expected OWins, got %v", status)
		}
	})

	t.Run("draw", func(t *testing.T) {
		board := Board{
			{"X", "O", "X"},
			{"X", "O", "O"},
			{"O", "X", "X"},
		}

		status := rules.GetGameStatus(board)

		if status != Draw {
			t.Errorf("expected Draw, got %v", status)
		}
	})
}
