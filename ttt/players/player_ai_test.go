package players

import (
	"testing"
	"ttt/boards"
)

func TestAIPlayer_MakesValidMoveOnEmptyBoard(t *testing.T) {
	ai := NewAIPlayer("X", "O")
	board := boards.NewBoard()

	move, err := ai.ReadMove(board)

	if err != nil {
		t.Fatalf("AI should not return error: %v", err)
	}

	boardCopy := board
	if err := boardCopy.MakeMove(move, "X"); err != nil {
		t.Errorf("AI chose invalid move %d: %v", move, err)
	}
}

func TestAIPlayer_MakesValidMoveOnPartialBoard(t *testing.T) {
	ai := NewAIPlayer("X", "O")
	board := boards.Board{
		{"X", "O", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	move, err := ai.ReadMove(board)

	if err != nil {
		t.Fatalf("AI should not return error: %v", err)
	}

	boardCopy := board
	if err := boardCopy.MakeMove(move, "X"); err != nil {
		t.Errorf("AI chose invalid move %d: %v", move, err)
	}
}

func TestAIPlayer_MakesValidMoveOnNearlyFullBoard(t *testing.T) {
	ai := NewAIPlayer("X", "O")
	board := boards.Board{
		{"X", "O", "X"},
		{"4", "5", "O"},
		{"7", "8", "9"},
	}

	move, err := ai.ReadMove(board)

	if err != nil {
		t.Fatalf("AI should not return error: %v", err)
	}

	boardCopy := board
	if err := boardCopy.MakeMove(move, "X"); err != nil {
		t.Errorf("AI chose invalid move %d: %v", move, err)
	}
}

func TestAIPlayer_MakesOnlyMoveAvailable(t *testing.T) {
	ai := NewAIPlayer("X", "O")
	board := boards.Board{
		{"X", "O", "X"},
		{"O", "X", "X"},
		{"X", "O", "9"},
	}

	move, err := ai.ReadMove(board)

	if err != nil {
		t.Fatalf("AI should not return error: %v", err)
	}

	if move != 9 {
		t.Errorf("AI should choose only available move 9, got %d", move)
	}
}

func TestAIPlayer_TakesHorizontalWinTopRow(t *testing.T) {
	ai := NewAIPlayer("X", "O")
	board := boards.Board{
		{"X", "X", "3"},
		{"O", "O", "6"},
		{"7", "8", "9"},
	}

	move, err := ai.ReadMove(board)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if move != 3 {
		t.Errorf("AI should take winning move 3, got %d", move)
	}

	boardCopy := board
	boardCopy.MakeMove(move, "X")
	if boardCopy.CheckWinner() != "X" {
		t.Error("Move 3 should result in AI winning")
	}
}

func TestAIPlayer_TakesVerticalWinLeftColumn(t *testing.T) {
	ai := NewAIPlayer("X", "O")
	board := boards.Board{
		{"X", "2", "3"},
		{"X", "O", "6"},
		{"7", "O", "9"},
	}

	move, err := ai.ReadMove(board)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if move != 7 {
		t.Errorf("AI should take winning move 7, got %d", move)
	}

	boardCopy := board
	boardCopy.MakeMove(move, "X")
	if boardCopy.CheckWinner() != "X" {
		t.Error("Move 7 should result in AI winning")
	}
}

func TestAIPlayer_TakesDiagonalWin(t *testing.T) {
	ai := NewAIPlayer("X", "O")
	board := boards.Board{
		{"X", "2", "O"},
		{"4", "X", "O"},
		{"7", "8", "9"},
	}

	move, err := ai.ReadMove(board)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if move != 9 {
		t.Errorf("AI should take winning move 9, got %d", move)
	}

	boardCopy := board
	boardCopy.MakeMove(move, "X")
	if boardCopy.CheckWinner() != "X" {
		t.Error("Move 9 should result in AI winning")
	}
}

func TestAIPlayer_OPlayerTakesWinningMove(t *testing.T) {
	ai := NewAIPlayer("O", "X")
	board := boards.Board{
		{"O", "O", "3"},
		{"X", "X", "6"},
		{"7", "8", "9"},
	}

	move, err := ai.ReadMove(board)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if move != 3 {
		t.Errorf("AI should take winning move 3, got %d", move)
	}

	boardCopy := board
	boardCopy.MakeMove(move, "O")
	if boardCopy.CheckWinner() != "O" {
		t.Error("Move 3 should result in O winning")
	}
}

func TestAIPlayer_BlocksHorizontalThreat(t *testing.T) {
	ai := NewAIPlayer("X", "O")
	board := boards.Board{
		{"O", "O", "3"},
		{"X", "5", "6"},
		{"X", "8", "9"},
	}

	move, err := ai.ReadMove(board)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if move != 3 {
		t.Errorf("AI should block at 3, got %d", move)
	}
}

func TestAIPlayer_BlocksVerticalThreat(t *testing.T) {
	ai := NewAIPlayer("X", "O")
	board := boards.Board{
		{"O", "X", "3"},
		{"O", "5", "6"},
		{"7", "8", "9"},
	}

	move, err := ai.ReadMove(board)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if move != 7 {
		t.Errorf("AI should block at 7, got %d", move)
	}
}

func TestAIPlayer_BlocksDiagonalThreat(t *testing.T) {
	ai := NewAIPlayer("X", "O")
	board := boards.Board{
		{"O", "2", "3"},
		{"4", "O", "6"},
		{"7", "8", "9"},
	}

	move, err := ai.ReadMove(board)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if move != 9 {
		t.Errorf("AI should block at 9, got %d", move)
	}
}

func TestAIPlayer_PrioritizesWinOverBlock(t *testing.T) {
	ai := NewAIPlayer("X", "O")
	board := boards.Board{
		{"X", "X", "3"},
		{"O", "O", "6"},
		{"7", "8", "9"},
	}

	move, err := ai.ReadMove(board)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if move != 3 {
		t.Errorf("AI should prioritize winning move (3) over blocking (6), got %d", move)
	}
}

func TestAIPlayer_ChoosesStrategicPositionOnEmptyBoard(t *testing.T) {
	ai := NewAIPlayer("X", "O")
	board := boards.NewBoard()

	move, err := ai.ReadMove(board)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	strategicMoves := []int{1, 3, 5, 7, 9}
	isStrategic := false
	for _, strategic := range strategicMoves {
		if move == strategic {
			isStrategic = true
			break
		}
	}

	if !isStrategic {
		t.Errorf("AI should choose strategic position, got %d", move)
	}
}

func TestAIPlayer_RespondsToCenter(t *testing.T) {
	ai := NewAIPlayer("X", "O")
	board := boards.Board{
		{"1", "2", "3"},
		{"4", "O", "6"},
		{"7", "8", "9"},
	}

	move, err := ai.ReadMove(board)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	corners := []int{1, 3, 7, 9}
	isCorner := false
	for _, corner := range corners {
		if move == corner {
			isCorner = true
			break
		}
	}

	if !isCorner {
		t.Errorf("AI should respond to center play with corner, got %d", move)
	}
}

func TestAIPlayer_HandlesLastMoveAvailable(t *testing.T) {
	board := boards.Board{
		{"X", "2", "O"},
		{"O", "X", "X"},
		{"X", "O", "O"},
	}
	ai := NewAIPlayer("X", "O")

	move, err := ai.ReadMove(board)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if move != 2 {
		t.Errorf("only move available is 2, but AI chose %d", move)
	}
}

func TestAIPlayer_HandlesForcedBlockScenario(t *testing.T) {
	ai := NewAIPlayer("O", "X")
	board := boards.Board{
		{"O", "X", "O"},
		{"X", "X", "6"},
		{"O", "8", "9"},
	}

	move, err := ai.ReadMove(board)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	validMoves := []int{6, 8, 9}
	isValid := false
	for _, valid := range validMoves {
		if move == valid {
			isValid = true
			break
		}
	}

	if !isValid {
		t.Errorf("AI should make legal move (6, 8, or 9), got %d", move)
	}
}

type GameSimulator struct{}

func NewGameSimulator() *GameSimulator {
	return &GameSimulator{}
}

type BoardDepth struct {
	board boards.Board
	depth int
}

func (gameSimulator *GameSimulator) simulateOpponentMoves(boardDepths []BoardDepth, opponentSymbol string) []BoardDepth {
	var result []BoardDepth

	for _, boardDepth := range boardDepths {
		for _, move := range boardDepth.board.AvailableMoves() {
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
		move, err := ai.ReadMove(boardDepth.board)
		if err != nil {
			continue
		}

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

func (gameSimulator *GameSimulator) partitionByGameStatus(boardDepths []BoardDepth) (finished, unfinished []BoardDepth) {
	for _, boardDepth := range boardDepths {
		if boardDepth.board.GetGameStatus() != boards.InProgress {
			finished = append(finished, boardDepth)
		} else {
			unfinished = append(unfinished, boardDepth)
		}
	}
	return
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
	for _, boardDepth := range finishedGames {
		winner := boardDepth.board.CheckWinner()
		if winner != "" && winner != aiSymbol {
			return false
		}
	}
	return true
}

func TestAIPlayer_NeverLosesAsSecondPlayer(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping exhaustive simulation in short mode")
	}

	simulator := NewGameSimulator()
	ai := NewAIPlayer("O", "X")
	emptyBoard := boards.NewBoard()

	opponentFirstMoves := simulator.simulateOpponentMoves(
		[]BoardDepth{{board: emptyBoard, depth: 0}},
		"X",
	)

	aiResponses := simulator.simulateAIMoves(opponentFirstMoves, ai)

	finishedGames := simulator.simulateAllGames(aiResponses, []BoardDepth{}, ai, "X")

	if !simulator.aiNeverLoses(finishedGames, "O") {
		t.Errorf("AI playing as O should never lose across all %d possible games", len(finishedGames))
	}

	t.Logf("Verified AI never loses across %d complete games playing as second player", len(finishedGames))
}

func TestAIPlayer_NeverLosesAsFirstPlayer(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping exhaustive simulation in short mode")
	}

	simulator := NewGameSimulator()
	ai := NewAIPlayer("X", "O")

	firstMove, _ := ai.ReadMove(boards.NewBoard())
	firstBoard := boards.NewBoard()
	firstBoard.MakeMove(firstMove, "X")

	startingPositions := []BoardDepth{{board: firstBoard, depth: 1}}

	finishedGames := simulator.simulateAllGames(startingPositions, []BoardDepth{}, ai, "O")

	if !simulator.aiNeverLoses(finishedGames, "X") {
		t.Errorf("AI playing as X should never lose across all %d possible games", len(finishedGames))
	}

	t.Logf("Verified AI never loses across %d complete games playing as first player", len(finishedGames))
}

func TestAIPlayer_VsAIAlwaysDraws(t *testing.T) {
	aiX := NewAIPlayer("X", "O")
	aiO := NewAIPlayer("O", "X")

	board := boards.NewBoard()
	currentPlayer := "X"

	for board.GetGameStatus() == boards.InProgress {
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
			t.Fatalf("AI made invalid move %d: %v", move, err)
		}

		if currentPlayer == "X" {
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}
	}

	status := board.GetGameStatus()
	if status != boards.Draw {
		t.Errorf("Two perfect AIs should draw, got status: %v, winner: %s",
			status, board.CheckWinner())
	}
}

func TestAIPlayer_RespondsQuickly(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping performance test in short mode")
	}

	ai := NewAIPlayer("X", "O")
	board := boards.NewBoard()

	for i := 0; i < 100; i++ {
		_, err := ai.ReadMove(board)
		if err != nil {
			t.Fatalf("AI failed on iteration %d: %v", i, err)
		}
	}
}
