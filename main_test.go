package main

import "testing"

func TestBoard(t *testing.T) {
	board := initBoard()
	if board[0][0] != "1" {
		t.Error("Expected board[0][0] to be 1, got:", board[0][0])
	}

	if board[0][1] != "2" {
		t.Error("Expected board[0][1] to be 2, got:", board[0][1])
	}

	if board[0][2] != "3" {
		t.Error("Expected board[0][2] to be 3, got:", board[0][2])
	}

	if board[1][0] != "4" {
		t.Error("Expected board[1][0] to be 4, got:", board[1][0])
	}

	if board[2][2] != "9" {
		t.Error("Expected board[2][2] to be 9, got:", board[2][2])
	}
}

func TestInitBoard(t *testing.T) {
	board := initBoard()

	expected := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	position := 0

	for i := range 3 {
		for j := range 3 {
			if board[i][j] != expected[position] {
				t.Errorf("Expected board[%d][%d] to be '%s', got '%s'",
					i, j, expected[position], board[i][j])
			}
			position++
		}
	}
}
