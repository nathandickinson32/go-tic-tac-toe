package main

import "fmt"

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

func makeMove(board *Board, position int, player string) bool {
	if position < 1 || position > 9 {
		return false
	}

	row := (position - 1) / 3
	col := (position - 1) % 3

	if board[row][col] == "X" || board[row][col] == "O" {
		return false
	}

	board[row][col] = player
	return true
}

func main() {
	fmt.Println(initBoard())
}
