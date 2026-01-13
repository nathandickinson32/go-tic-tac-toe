package main

import "fmt"

type Board [3][3]string

func initBoard() Board {
	var board Board
	board[0][0] = "1"
	return board
}

func main() {
	fmt.Println("Welcome to Tic-Tac-Toe!")
}
