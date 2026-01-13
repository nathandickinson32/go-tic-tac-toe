package main

import "testing"

func TestInitBoard(t *testing.T) {
	board := initBoard()

	if board[0][0] != "1" {
		t.Error("Expected first position to be 1")
	}
}
