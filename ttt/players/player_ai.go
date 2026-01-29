package players

import (
	"math"
	"ttt/boards"
)

const (
	WinScore     = 10.0
	DrawScore    = 0.0
	DepthPenalty = 1.0
)

type AIPlayer struct {
	playerSymbol   string
	opponentSymbol string
}

func NewAIPlayer(playerSymbol string, opponentSymbol string) *AIPlayer {
	return &AIPlayer{
		playerSymbol:   playerSymbol,
		opponentSymbol: opponentSymbol,
	}
}

func (ai *AIPlayer) getTerminalScore(board boards.Board, depth int) (float64, bool) {
	winner := board.CheckWinner()

	if winner == ai.playerSymbol {
		return WinScore - float64(depth)*DepthPenalty, true
	}

	if winner == ai.opponentSymbol {
		return float64(depth)*DepthPenalty - WinScore, true
	}

	if board.GetGameStatus() == boards.Draw {
		return DrawScore, true
	}

	return DrawScore, false
}

func (ai *AIPlayer) evaluateMoveForPlayer(board boards.Board, move int, player string, depth int, isMaximizing bool) float64 {
	boardCopy := board
	if err := boardCopy.MakeMove(move, player); err != nil {
		if isMaximizing {
			return math.Inf(1)
		}
		return math.Inf(-1)
	}
	return ai.minimax(boardCopy, depth+1, isMaximizing)
}

func (ai *AIPlayer) minimizeScore(board boards.Board, depth int) float64 {
	minScore := math.Inf(1)

	for _, move := range board.AvailableMoves() {
		score := ai.evaluateMoveForPlayer(board, move, ai.opponentSymbol, depth, true)
		minScore = math.Min(minScore, score)
	}

	return minScore
}

func (ai *AIPlayer) maximizeScore(board boards.Board, depth int) float64 {
	maxScore := math.Inf(-1)

	for _, move := range board.AvailableMoves() {
		score := ai.evaluateMoveForPlayer(board, move, ai.playerSymbol, depth, false)
		maxScore = math.Max(maxScore, score)
	}

	return maxScore
}

func (ai *AIPlayer) minimax(board boards.Board, depth int, isMaximizing bool) float64 {
	if score, isTerminal := ai.getTerminalScore(board, depth); isTerminal {
		return score
	}

	if isMaximizing {
		return ai.maximizeScore(board, depth)
	}
	return ai.minimizeScore(board, depth)
}

func (ai *AIPlayer) evaluateMove(board boards.Board, move int) float64 {
	boardCopy := board
	if err := boardCopy.MakeMove(move, ai.playerSymbol); err != nil {
		return math.Inf(-1)
	}
	return ai.minimax(boardCopy, 0, false)
}

func (ai *AIPlayer) findBestMove(board boards.Board) int {
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

func (ai *AIPlayer) ReadMove(board boards.Board) (int, error) {
	bestMove := ai.findBestMove(board)
	return bestMove, nil
}
