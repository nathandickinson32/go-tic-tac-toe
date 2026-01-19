package main

import (
	"bufio"
)

type PlayerType int

const (
	Human PlayerType = iota
	AI
)

func createPlayerInput(
	playerType PlayerType,
	symbol string,
	opponentSymbol string,
	reader *bufio.Reader,
	output *ConsoleOutput,
	rules *GameRules,
) InputReader {
	if playerType == AI {
		return NewAIPlayer(symbol, opponentSymbol, rules, output)
	}
	return NewConsoleInput(reader, output)
}
