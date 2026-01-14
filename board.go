package main

import (
	"fmt"
)

type Board [3][3]string

func initBoard() Board {
	var board Board
	cell := 1

	for row := range 3 {
		for col := range 3 {
			board[row][col] = fmt.Sprintf("%d", cell)
			cell++
		}
	}
	return board
}

func getAvailableMoves(board Board) []int {
	var moves []int
	cell := 1
	for row := range 3 {
		for col := range 3 {
			if board[row][col] != "X" && board[row][col] != "O" {
				moves = append(moves, cell)
			}
			cell++
		}
	}
	return moves
}

func isValidMove(board *Board, position int, row int, col int) bool {
	if position < 1 || position > 9 || board[row][col] == "X" || board[row][col] == "O" {
		return false
	}
	return true
}

func makeMove(board *Board, position int, player string) bool {
	row := (position - 1) / 3
	col := (position - 1) % 3

	valid := isValidMove(board, position, row, col)

	if valid {
		board[row][col] = player
	}
	return valid
}

func checkWinner(board Board) string {
	lines := [8][3][2]int{
		{{0, 0}, {0, 1}, {0, 2}},
		{{1, 0}, {1, 1}, {1, 2}},
		{{2, 0}, {2, 1}, {2, 2}},
		{{0, 0}, {1, 0}, {2, 0}},
		{{0, 1}, {1, 1}, {2, 1}},
		{{0, 2}, {1, 2}, {2, 2}},
		{{0, 0}, {1, 1}, {2, 2}},
		{{0, 2}, {1, 1}, {2, 0}},
	}

	for _, line := range lines {
		a := board[line[0][0]][line[0][1]]
		b := board[line[1][0]][line[1][1]]
		c := board[line[2][0]][line[2][1]]

		if a == b && b == c {
			return a
		}
	}

	return ""
}
