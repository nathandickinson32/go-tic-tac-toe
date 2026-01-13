package main

import "fmt"

type Board [3][3]string

func initBoard() Board {
	var board Board
	num := 1

	for i := range 3 {
		for j := range 3 {
			board[i][j] = fmt.Sprintf("%d", num)
			num++
		}
	}
	return board
}

func main() {
	fmt.Println(initBoard())
}
