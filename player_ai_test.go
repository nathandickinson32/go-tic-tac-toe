package main

import (
	"bytes"
	"slices"
	"testing"
)

type GameSimulator struct {
	rules *GameRules
}

func NewGameSimulator() *GameSimulator {
	return &GameSimulator{
		rules: NewGameRules(),
	}
}

type BoardDepth struct {
	board Board
	depth int
}

func (gameSimulator *GameSimulator) simulateOpponentMoves(boardDepths []BoardDepth, opponentSymbol string) []BoardDepth {
	var result []BoardDepth

	for _, boardDepth := range boardDepths {
		availableMoves := boardDepth.board.AvailableMoves()
		for _, move := range availableMoves {
			boardCopy := boardDepth.board
			if err := boardCopy.MakeMove(move, opponentSymbol); err != nil {
				continue
			}
			result = append(result, BoardDepth{
				board: boardCopy,
				depth: boardDepth.depth + 1,
			})
		}
	}

	return result
}

func (gameSimulator *GameSimulator) simulateAIMoves(boardDepths []BoardDepth, ai *AIPlayer) []BoardDepth {
	var result []BoardDepth

	for _, boardDepth := range boardDepths {
		move, _ := ai.ReadMove(boardDepth.board)
		boardCopy := boardDepth.board
		if err := boardCopy.MakeMove(move, ai.playerSymbol); err != nil {
			continue
		}
		result = append(result, BoardDepth{
			board: boardCopy,
			depth: boardDepth.depth + 1,
		})
	}

	return result
}

func (gameSimulator *GameSimulator) partitionByGameStatus(boardDepths []BoardDepth) (finished []BoardDepth, unfinished []BoardDepth) {
	for _, boardDepth := range boardDepths {
		if gameSimulator.rules.GetGameStatus(boardDepth.board) != InProgress {
			finished = append(finished, boardDepth)
		} else {
			unfinished = append(unfinished, boardDepth)
		}
	}
	return finished, unfinished
}

func (gameSimulator *GameSimulator) simulateAllGames(
	unfinished []BoardDepth,
	finished []BoardDepth,
	ai *AIPlayer,
	opponentSymbol string,
) []BoardDepth {
	if len(unfinished) == 0 {
		return finished
	}

	opponentBoards := gameSimulator.simulateOpponentMoves(unfinished, opponentSymbol)
	opponentFinished, opponentUnfinished := gameSimulator.partitionByGameStatus(opponentBoards)

	aiBoards := gameSimulator.simulateAIMoves(opponentUnfinished, ai)
	aiFinished, aiUnfinished := gameSimulator.partitionByGameStatus(aiBoards)

	allFinished := append(finished, opponentFinished...)
	allFinished = append(allFinished, aiFinished...)

	return gameSimulator.simulateAllGames(aiUnfinished, allFinished, ai, opponentSymbol)
}

func (gameSimulator *GameSimulator) aiNeverLoses(finishedGames []BoardDepth, aiSymbol string) bool {
	for _, bd := range finishedGames {
		winner := gameSimulator.rules.CheckWinner(bd.board)
		if winner != "" && winner != aiSymbol {
			return false
		}
	}
	return true
}

func TestAIPlayer_ChooseBestMove(t *testing.T) {
	rules := NewGameRules()
	var output bytes.Buffer
	consoleOutput := NewConsoleOutput(&output)

	t.Run("returns only available move", func(t *testing.T) {
		ai := NewAIPlayer("X", "O", rules, consoleOutput)

		board := Board{
			{"X", "2", "O"},
			{"O", "X", "X"},
			{"X", "O", "O"},
		}

		move, err := ai.ReadMove(board)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if move != 2 {
			t.Errorf("should choose only available move 2, got %d", move)
		}
	})

	t.Run("chooses corner on empty board", func(t *testing.T) {
		ai := NewAIPlayer("X", "O", rules, consoleOutput)

		board := NewBoard()

		move, err := ai.ReadMove(board)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		corners := []int{1, 3, 7, 9}
		center := 5
		validMoves := append(corners, center)

		isValid := slices.Contains(validMoves, move)

		if !isValid {
			t.Errorf("AI should choose center or corner, got %d", move)
		}
	})

	t.Run("chooses winning move over blocking", func(t *testing.T) {
		ai := NewAIPlayer("X", "O", rules, consoleOutput)

		board := Board{
			{"X", "X", "3"},
			{"O", "5", "6"},
			{"O", "8", "9"},
		}

		move, err := ai.ReadMove(board)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if move != 3 {
			t.Errorf("AI should choose winning move 3, got %d", move)
		}
	})

	t.Run("chooses blocking move when no winning move", func(t *testing.T) {
		ai := NewAIPlayer("X", "O", rules, consoleOutput)

		board := Board{
			{"X", "X", "3"},
			{"4", "O", "6"},
			{"7", "8", "9"},
		}

		move, err := ai.ReadMove(board)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if move != 3 {
			t.Errorf("AI should block at position 3, got %d", move)
		}
	})

	t.Run("chooses best move for O", func(t *testing.T) {
		ai := NewAIPlayer("O", "X", rules, consoleOutput)

		board := Board{
			{"X", "O", "X"},
			{"4", "O", "6"},
			{"7", "X", "9"},
		}

		move, err := ai.ReadMove(board)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		validMoves := []int{4, 6}
		isValid := slices.Contains(validMoves, move)

		if !isValid {
			t.Errorf("AI should choose strategic move (4 or 6), got %d", move)
		}
	})

	t.Run("chooses best move when multiple good options", func(t *testing.T) {
		ai := NewAIPlayer("O", "X", rules, consoleOutput)

		board := Board{
			{"X", "O", "X"},
			{"4", "5", "6"},
			{"O", "X", "9"},
		}

		move, err := ai.ReadMove(board)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		validMoves := []int{4, 5, 6, 9}
		isValid := slices.Contains(validMoves, move)

		if !isValid {
			t.Errorf("AI should choose valid strategic move, got %d", move)
		}
	})
}

func TestAIPlayer_Basics(t *testing.T) {
	rules := NewGameRules()
	var output bytes.Buffer
	consoleOutput := NewConsoleOutput(&output)

	t.Run("always returns valid move", func(t *testing.T) {
		ai := NewAIPlayer("X", "O", rules, consoleOutput)

		testBoards := []Board{
			NewBoard(),
			{
				{"X", "O", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			{
				{"X", "O", "X"},
				{"4", "5", "O"},
				{"7", "8", "9"},
			},
		}

		for i, board := range testBoards {
			move, err := ai.ReadMove(board)

			if err != nil {
				t.Errorf("board %d: unexpected error: %v", i, err)
				continue
			}

			boardCopy := board
			if err := boardCopy.MakeMove(move, "X"); err != nil {
				t.Errorf("board %d: AI chose invalid move %d: %v", i, move, err)
			}
		}
	})

	t.Run("chooses winning move", func(t *testing.T) {
		ai := NewAIPlayer("X", "O", rules, consoleOutput)

		board := Board{
			{"X", "X", "3"},
			{"O", "O", "6"},
			{"7", "8", "9"},
		}

		move, err := ai.ReadMove(board)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if move != 3 {
			t.Errorf("AI should choose winning move 3, got %d", move)
		}
	})

	t.Run("blocks opponent winning move", func(t *testing.T) {
		ai := NewAIPlayer("X", "O", rules, consoleOutput)

		board := Board{
			{"O", "O", "3"},
			{"X", "5", "6"},
			{"X", "8", "9"},
		}

		move, err := ai.ReadMove(board)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if move != 3 {
			t.Errorf("AI should block at position 3, got %d", move)
		}
	})
}

func TestAIPlayer_Minimax(t *testing.T) {
	rules := NewGameRules()
	var output bytes.Buffer
	consoleOutput := NewConsoleOutput(&output)

	t.Run("scores winning position higher than tie", func(t *testing.T) {
		ai := NewAIPlayer("X", "O", rules, consoleOutput)

		winningBoard := Board{
			{"X", "X", "3"},
			{"O", "O", "6"},
			{"7", "8", "9"},
		}

		tieBoard := Board{
			{"X", "O", "X"},
			{"X", "O", "O"},
			{"O", "X", "X"},
		}

		winScore := ai.minimax(winningBoard, 0, true)
		tieScore := ai.minimax(tieBoard, 0, true)

		if winScore <= tieScore {
			t.Errorf("winning score (%f) should be greater than tie score (%f)", winScore, tieScore)
		}
	})

	t.Run("scores tie higher than losing position", func(t *testing.T) {
		ai := NewAIPlayer("X", "O", rules, consoleOutput)

		tieBoard := Board{
			{"X", "O", "X"},
			{"X", "O", "O"},
			{"O", "X", "X"},
		}

		losingBoard := Board{
			{"O", "O", "O"},
			{"X", "X", "6"},
			{"7", "8", "9"},
		}

		tieScore := ai.minimax(tieBoard, 0, true)
		loseScore := ai.minimax(losingBoard, 0, true)

		if tieScore <= loseScore {
			t.Errorf("tie score (%f) should be greater than losing score (%f)", tieScore, loseScore)
		}
	})

	t.Run("gives higher scores to better moves", func(t *testing.T) {
		ai := NewAIPlayer("X", "O", rules, consoleOutput)

		board := Board{
			{"X", "X", "3"},
			{"O", "O", "6"},
			{"7", "8", "9"},
		}

		makeMoveAndScore := func(position int) float64 {
			boardCopy := board
			boardCopy.MakeMove(position, "X")
			return ai.minimax(boardCopy, 1, false)
		}

		bestMoveScore := makeMoveAndScore(3)
		nextBestScore := makeMoveAndScore(6)
		worstMoveScore := makeMoveAndScore(9)

		if bestMoveScore <= nextBestScore {
			t.Errorf("best move score (%f) should be > next best score (%f)", bestMoveScore, nextBestScore)
		}
		if nextBestScore <= worstMoveScore {
			t.Errorf("next best score (%f) should be > worst score (%f)", nextBestScore, worstMoveScore)
		}
	})

	t.Run("maximizes score for maximizing player", func(t *testing.T) {
		ai := NewAIPlayer("X", "O", rules, consoleOutput)

		board := Board{
			{"X", "X", "3"},
			{"O", "O", "6"},
			{"7", "8", "9"},
		}

		score := ai.minimax(board, 0, true)

		if score < 0 {
			t.Errorf("maximizing player should have non-negative score, got %f", score)
		}
	})

	t.Run("minimizes score when opponent is winning", func(t *testing.T) {
		ai := NewAIPlayer("X", "O", rules, consoleOutput)

		board := Board{
			{"X", "X", "3"},
			{"O", "O", "6"},
			{"7", "8", "9"},
		}

		score := ai.minimax(board, 0, false)

		if score >= 0 {
			t.Errorf("opponent winning should have negative score, got %f", score)
		}
	})

	t.Run("returns 0 for tie game", func(t *testing.T) {
		ai := NewAIPlayer("X", "O", rules, consoleOutput)

		board := Board{
			{"X", "O", "X"},
			{"X", "O", "O"},
			{"O", "X", "X"},
		}

		score := ai.minimax(board, 0, true)

		if score != 0 {
			t.Errorf("tie game should score 0, got %f", score)
		}
	})
}

func TestAIPlayer_Simulation(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping simulation in short mode")
	}

	rules := NewGameRules()
	var output bytes.Buffer
	consoleOutput := NewConsoleOutput(&output)
	simulator := NewGameSimulator()

	t.Run("AI never loses as O playing second", func(t *testing.T) {
		ai := NewAIPlayer("O", "X", rules, consoleOutput)

		emptyBoard := NewBoard()
		opponentFirstMoves := simulator.simulateOpponentMoves(
			[]BoardDepth{{board: emptyBoard, depth: 0}},
			"X",
		)

		aiResponses := simulator.simulateAIMoves(opponentFirstMoves, ai)

		finishedGames := simulator.simulateAllGames(aiResponses, []BoardDepth{}, ai, "X")

		if !simulator.aiNeverLoses(finishedGames, "O") {
			t.Errorf("AI as O should never lose across all %d possible games", len(finishedGames))
		}

		t.Logf("Simulated %d complete games: AI never lost", len(finishedGames))
	})

	t.Run("AI never loses as X playing first", func(t *testing.T) {
		ai := NewAIPlayer("X", "O", rules, consoleOutput)

		firstMove, _ := ai.ReadMove(NewBoard())
		firstBoard := NewBoard()
		firstBoard.MakeMove(firstMove, "X")

		startingPositions := []BoardDepth{{board: firstBoard, depth: 1}}
		finishedGames := simulator.simulateAllGames(startingPositions, []BoardDepth{}, ai, "O")

		if !simulator.aiNeverLoses(finishedGames, "X") {
			t.Errorf("AI as X should never lose across all %d possible games", len(finishedGames))
		}

		t.Logf("Simulated %d complete games: AI never lost", len(finishedGames))
	})
}

func TestAIPlayer_TwoAIsAlwaysDraw(t *testing.T) {
	t.Run("two perfect AIs always draw", func(t *testing.T) {
		rules := NewGameRules()
		var output bytes.Buffer
		consoleOutput := NewConsoleOutput(&output)

		aiX := NewAIPlayer("X", "O", rules, consoleOutput)
		aiO := NewAIPlayer("O", "X", rules, consoleOutput)

		board := NewBoard()
		currentPlayer := "X"

		for rules.GetGameStatus(board) == InProgress {
			var move int
			var err error

			if currentPlayer == "X" {
				move, err = aiX.ReadMove(board)
			} else {
				move, err = aiO.ReadMove(board)
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if err := board.MakeMove(move, currentPlayer); err != nil {
				t.Fatalf("AI made invalid move: %d, error: %v", move, err)
			}

			if currentPlayer == "X" {
				currentPlayer = "O"
			} else {
				currentPlayer = "X"
			}
		}

		status := rules.GetGameStatus(board)
		if status != Draw {
			t.Errorf("Two perfect AIs should draw, got status: %v", status)
		}
	})
}
