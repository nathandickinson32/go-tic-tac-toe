package main

import "fmt"

type Board [3][3]string

func initBoard() Board {
	var board Board
	num := 1

	for row := range 3 {
		for col := range 3 {
			board[row][col] = fmt.Sprintf("%d", num)
			num++
		}
	}
	return board
}

func getAvailableMoves(board Board) []int {
	var moves []int
	num := 1
	for i := range 3 {
		for j := range 3 {
			if board[i][j] != "X" && board[i][j] != "O" {
				moves = append(moves, num)
			}
			num = num + 1
		}
	}
	return moves
}

func main() {
	fmt.Println(initBoard())
}
