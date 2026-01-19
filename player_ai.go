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

func (ai *AIPlayer) getTerminalScore(board Board, depth int) (float64, bool) {
	winner := ai.rules.CheckWinner(board)

	if winner == ai.playerSymbol {
		return 10.0 - float64(depth), true
	}

	if winner == ai.opponentSymbol {
		return float64(depth) - 10.0, true
	}

	if ai.rules.GetGameStatus(board) == Draw {
		return 0.0, true
	}

	return 0.0, false
}

func (ai *AIPlayer) evaluateMoveForPlayer(board Board, move int, player string, depth int, isMaximizing bool) float64 {
	boardCopy := board
	if err := boardCopy.MakeMove(move, player); err != nil {
		if isMaximizing {
			return math.Inf(1)
		}
		return math.Inf(-1)
	}
	return ai.minimax(boardCopy, depth+1, isMaximizing)
}

func (ai *AIPlayer) minimizeScore(board Board, depth int) float64 {
	minScore := math.Inf(1)

	for _, move := range board.AvailableMoves() {
		score := ai.evaluateMoveForPlayer(board, move, ai.opponentSymbol, depth, true)
		minScore = math.Min(minScore, score)
	}

	return minScore
}

func (ai *AIPlayer) maximizeScore(board Board, depth int) float64 {
	maxScore := math.Inf(-1)

	for _, move := range board.AvailableMoves() {
		score := ai.evaluateMoveForPlayer(board, move, ai.playerSymbol, depth, false)
		maxScore = math.Max(maxScore, score)
	}

	return maxScore
}

func (ai *AIPlayer) minimax(board Board, depth int, isMaximizing bool) float64 {
	if score, isTerminal := ai.getTerminalScore(board, depth); isTerminal {
		return score
	}

	if isMaximizing {
		return ai.maximizeScore(board, depth)
	}
	return ai.minimizeScore(board, depth)
}

func (ai *AIPlayer) evaluateMove(board Board, move int) float64 {
	boardCopy := board
	if err := boardCopy.MakeMove(move, ai.playerSymbol); err != nil {
		return math.Inf(-1)
	}
	return ai.minimax(boardCopy, 0, false)
}

func (ai *AIPlayer) findBestMove(board Board) int {
	bestScore := math.Inf(-1)
	bestMove := 0

	for _, move := range board.AvailableMoves() {
		score := ai.evaluateMove(board, move)
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}

	return bestMove
}

func (ai *AIPlayer) ReadMove(board Board) (int, error) {
	bestMove := ai.findBestMove(board)
	return bestMove, nil
}
