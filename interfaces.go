package main

type InputReader interface {
	ReadMove(board Board) (int, error)
}
