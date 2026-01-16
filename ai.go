package main

import "math"

type AIPlayer struct {
	playerSymbol   string
	opponentSymbol string
	rules          *GameRules
	output         *ConsoleOutput
}

func NewAIPlayer(
	playerSymbol string,
	opponentSymbol string,
	rules *GameRules,
	output *ConsoleOutput,
) *AIPlayer {
	return &AIPlayer{
		playerSymbol:   playerSymbol,
		opponentSymbol: opponentSymbol,
		rules:          rules,
		output:         output,
	}
}

func (ai *AIPlayer) ReadMove(board Board) (int, error) {
	ai.output.ShowPrompt()
	bestMove := ai.findBestMove(board)
	return bestMove, nil
}

func (ai *AIPlayer) findBestMove(board Board) int {
	bestScore := math.Inf(-1)
	bestMove := 0

	for _, move := range board.AvailableMoves() {
		boardCopy := board
		if err := boardCopy.MakeMove(move, ai.playerSymbol); err != nil {
			continue
		}

		score := ai.minimax(boardCopy, 0, false)

		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}

	return bestMove
}

func (ai *AIPlayer) minimax(board Board, depth int, isMaximizing bool) float64 {
	status := ai.rules.GetGameStatus(board)

	if winner := ai.rules.CheckWinner(board); winner == ai.playerSymbol {
		return 10.0 - float64(depth)
	} else if winner == ai.opponentSymbol {
		return float64(depth) - 10.0
	} else if status == Draw {
		return 0.0
	}

	if isMaximizing {
		maxScore := math.Inf(-1)
		for _, move := range board.AvailableMoves() {
			boardCopy := board
			if err := boardCopy.MakeMove(move, ai.playerSymbol); err != nil {
				continue
			}
			score := ai.minimax(boardCopy, depth+1, false)
			maxScore = math.Max(maxScore, score)
		}
		return maxScore
	} else {
		minScore := math.Inf(1)
		for _, move := range board.AvailableMoves() {
			boardCopy := board
			if err := boardCopy.MakeMove(move, ai.opponentSymbol); err != nil {
				continue
			}
			score := ai.minimax(boardCopy, depth+1, true)
			minScore = math.Min(minScore, score)
		}
		return minScore
	}
}
