package game

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
	tttio "ttt/io"
	"ttt/players"
)

func runGameSimulation(input string) string {
	reader := strings.NewReader(input)
	var output bytes.Buffer

	bufReader := bufio.NewReader(reader)
	humanPlayerX := players.NewHumanPlayer(bufReader, &output)
	humanPlayerO := players.NewHumanPlayer(bufReader, &output)

	game := NewGame(humanPlayerX, humanPlayerO, &output)
	game.PlayGame()

	return output.String()
}

func TestGame_DisplaysWelcomeMessage(t *testing.T) {
	input := "1\n2\n3\n4\n5\n6\n7\n8\n9\n"

	result := runGameSimulation(input)

	if !strings.Contains(result, "Welcome") {
		t.Error("should display welcome message to user")
	}
}

func TestGame_DisplaysInitialBoardWithAllPositions(t *testing.T) {
	input := "1\n2\n3\n4\n5\n6\n7\n8\n9\n"

	result := runGameSimulation(input)

	for position := 1; position <= 9; position++ {
		posStr := string(rune('0' + position))
		if !strings.Contains(result, posStr) {
			t.Errorf("initial board should display position %d", position)
		}
	}
}

func TestGame_StartsWithPlayerX(t *testing.T) {
	input := "5\n1\n9\n2\n3\n4\n6\n7\n8\n"

	result := runGameSimulation(input)

	xTurnPos := strings.Index(result, "Player X's turn")
	oTurnPos := strings.Index(result, "Player O's turn")

	if xTurnPos == -1 {
		t.Fatal("should show Player X's turn")
	}

	if oTurnPos == -1 {
		t.Fatal("should show Player O's turn")
	}

	if xTurnPos >= oTurnPos {
		t.Error("Player X should have first turn before O")
	}
}

func TestGame_AlternatesBetweenPlayers(t *testing.T) {
	input := "1\n4\n2\n5\n3\n"

	result := runGameSimulation(input)

	xTurnCount := strings.Count(result, "Player X's turn")
	oTurnCount := strings.Count(result, "Player O's turn")

	if xTurnCount != 3 {
		t.Errorf("should have 3 X turns, got %d", xTurnCount)
	}

	if oTurnCount != 2 {
		t.Errorf("should have 2 O turns, got %d", oTurnCount)
	}
}

func TestGame_UpdatesBoardToShowXMoves(t *testing.T) {
	input := "5\n1\n9\n2\n3\n4\n6\n7\n8\n"

	result := runGameSimulation(input)

	if !strings.Contains(result, " X ") {
		t.Error("board should show X after X's moves")
	}
}

func TestGame_UpdatesBoardToShowOMoves(t *testing.T) {
	input := "5\n1\n9\n2\n3\n4\n6\n7\n8\n"

	result := runGameSimulation(input)

	if !strings.Contains(result, " O ") {
		t.Error("board should show O after O's moves")
	}
}

func TestGame_DisplaysBoardMultipleTimes(t *testing.T) {
	input := "5\n1\n9\n2\n3\n4\n6\n7\n8\n"

	result := runGameSimulation(input)

	boardCount := strings.Count(result, " | ")
	if boardCount < 3 {
		t.Errorf("should display board multiple times, got %d board displays", boardCount/3)
	}
}

func TestGame_DetectsXWinTopRow(t *testing.T) {
	input := "1\n4\n2\n5\n3\n"

	result := runGameSimulation(input)

	if !strings.Contains(result, "Player X wins") {
		t.Error("should announce X as winner for top row")
	}

	moveCount := strings.Count(result, "Enter your move")
	if moveCount != 5 {
		t.Errorf("should end after 5 moves, got %d", moveCount)
	}
}

func TestGame_DetectsOWinMiddleRow(t *testing.T) {
	input := "1\n4\n2\n5\n9\n6\n"

	result := runGameSimulation(input)

	if !strings.Contains(result, "Player O wins") {
		t.Error("should announce O as winner for middle row")
	}

	moveCount := strings.Count(result, "Enter your move")
	if moveCount != 6 {
		t.Errorf("should end after 6 moves, got %d", moveCount)
	}
}

func TestGame_DetectsXWinBottomRow(t *testing.T) {
	input := "7\n1\n8\n2\n9\n"

	result := runGameSimulation(input)

	if !strings.Contains(result, "Player X wins") {
		t.Error("should announce X as winner for bottom row")
	}

	moveCount := strings.Count(result, "Enter your move")
	if moveCount != 5 {
		t.Errorf("should end after 5 moves, got %d", moveCount)
	}
}

func TestGame_DetectsXWinLeftColumn(t *testing.T) {
	input := "1\n2\n4\n3\n7\n"

	result := runGameSimulation(input)

	if !strings.Contains(result, "Player X wins") {
		t.Error("should announce X as winner for left column")
	}

	moveCount := strings.Count(result, "Enter your move")
	if moveCount != 5 {
		t.Errorf("should end after 5 moves, got %d", moveCount)
	}
}

func TestGame_DetectsOWinMiddleColumn(t *testing.T) {
	input := "1\n2\n3\n5\n9\n8\n"

	result := runGameSimulation(input)

	if !strings.Contains(result, "Player O wins") {
		t.Error("should announce O as winner for middle column")
	}

	moveCount := strings.Count(result, "Enter your move")
	if moveCount != 6 {
		t.Errorf("should end after 6 moves, got %d", moveCount)
	}
}

func TestGame_DetectsXWinRightColumn(t *testing.T) {
	input := "3\n1\n6\n2\n9\n"

	result := runGameSimulation(input)

	if !strings.Contains(result, "Player X wins") {
		t.Error("should announce X as winner for right column")
	}

	moveCount := strings.Count(result, "Enter your move")
	if moveCount != 5 {
		t.Errorf("should end after 5 moves, got %d", moveCount)
	}
}

func TestGame_DetectsXWinMainDiagonal(t *testing.T) {
	input := "1\n2\n5\n3\n9\n"

	result := runGameSimulation(input)

	if !strings.Contains(result, "Player X wins") {
		t.Error("should announce X as winner for main diagonal")
	}

	moveCount := strings.Count(result, "Enter your move")
	if moveCount != 5 {
		t.Errorf("should end after 5 moves, got %d", moveCount)
	}
}

func TestGame_DetectsOWinAntiDiagonal(t *testing.T) {
	input := "1\n3\n2\n5\n8\n7\n"

	result := runGameSimulation(input)

	if !strings.Contains(result, "Player O wins") {
		t.Error("should announce O as winner for anti-diagonal")
	}

	moveCount := strings.Count(result, "Enter your move")
	if moveCount != 6 {
		t.Errorf("should end after 6 moves, got %d", moveCount)
	}
}

func TestGame_DetectsDrawWhenBoardFull(t *testing.T) {
	input := "1\n2\n3\n5\n4\n6\n8\n7\n9\n"

	result := runGameSimulation(input)

	if !strings.Contains(result, "Game Over") {
		t.Error("should announce game over for draw")
	}

	if !strings.Contains(result, "full") {
		t.Error("should indicate board is full")
	}

	if strings.Contains(result, "wins") {
		t.Error("draw game should not announce a winner")
	}
}

func TestGame_StopsAfterXWinsIn5Moves(t *testing.T) {
	input := "1\n4\n2\n5\n3\n"

	result := runGameSimulation(input)

	promptCount := strings.Count(result, "Enter your move")
	if promptCount != 5 {
		t.Errorf("should prompt exactly 5 times, got %d", promptCount)
	}

	winCount := strings.Count(result, "wins")
	if winCount != 1 {
		t.Errorf("should announce winner exactly once, got %d", winCount)
	}
}

func TestGame_StopsAfterOWinsIn6Moves(t *testing.T) {
	input := "1\n4\n2\n5\n9\n6\n"

	result := runGameSimulation(input)

	promptCount := strings.Count(result, "Enter your move")
	if promptCount != 6 {
		t.Errorf("should prompt exactly 6 times, got %d", promptCount)
	}

	winCount := strings.Count(result, "wins")
	if winCount != 1 {
		t.Errorf("should announce winner exactly once, got %d", winCount)
	}
}

func TestGame_StopsAfterDrawIn9Moves(t *testing.T) {
	input := "1\n2\n3\n5\n4\n6\n8\n7\n9\n"

	result := runGameSimulation(input)

	promptCount := strings.Count(result, "Enter your move")
	if promptCount != 9 {
		t.Errorf("draw game should prompt exactly 9 times, got %d", promptCount)
	}

	if strings.Count(result, "Game Over") != 1 {
		t.Error("should announce game over exactly once")
	}
}

func TestGame_PromptsForEachMove(t *testing.T) {
	input := "1\n4\n2\n5\n3\n"

	result := runGameSimulation(input)

	promptCount := strings.Count(result, "Enter your move")
	if promptCount != 5 {
		t.Errorf("should prompt 5 times for 5 moves, got %d", promptCount)
	}
}

func TestGame_AnnouncesXTurns(t *testing.T) {
	input := "1\n2\n3\n5\n4\n6\n8\n7\n9\n"

	result := runGameSimulation(input)

	xAnnouncements := strings.Count(result, "Player X's turn")
	if xAnnouncements != 5 {
		t.Errorf("should announce X's turn 5 times, got %d", xAnnouncements)
	}
}

func TestGame_AnnouncesOTurns(t *testing.T) {
	input := "1\n2\n3\n5\n4\n6\n8\n7\n9\n"

	result := runGameSimulation(input)

	oAnnouncements := strings.Count(result, "Player O's turn")
	if oAnnouncements != 4 {
		t.Errorf("should announce O's turn 4 times, got %d", oAnnouncements)
	}
}

func TestStartGame_HumanVsHumanDeclineReplay(t *testing.T) {
	input := "1\n1\n1\n4\n2\n5\n3\nn\n"
	reader := strings.NewReader(input)
	var output bytes.Buffer

	bufReader := bufio.NewReader(reader)

	tttio.ShowNewline(&output)
	testGame := BuildGame(bufReader, &output)
	testGame.PlayGame()

	tttio.ShowPlayAgainPrompt(&output)
	playAgain, _ := tttio.ReadPlayAgain(bufReader, &output)

	if playAgain {
		t.Error("should not want to play again")
	}

	result := output.String()

	if !strings.Contains(result, "Select Player X type") {
		t.Error("should show player X type selection")
	}

	if !strings.Contains(result, "Select Player O type") {
		t.Error("should show player O type selection")
	}

	if !strings.Contains(result, "Welcome") {
		t.Error("should show welcome message")
	}

	if !strings.Contains(result, "Player X wins") {
		t.Error("should show X wins")
	}

	if !strings.Contains(result, "Play again") {
		t.Error("should show play again prompt")
	}
}

func TestStartGame_HumanVsAI(t *testing.T) {
	input := "1\n2\n1\n2\n3\n4\n5\nn\n"
	reader := strings.NewReader(input)
	var output bytes.Buffer

	bufReader := bufio.NewReader(reader)

	tttio.ShowNewline(&output)
	testGame := BuildGame(bufReader, &output)
	testGame.PlayGame()

	result := output.String()

	if !strings.Contains(result, "Select Player X type") {
		t.Error("should show player X type selection")
	}

	if !strings.Contains(result, "Select Player O type") {
		t.Error("should show player O type selection")
	}

	if !strings.Contains(result, "Welcome") {
		t.Error("should show welcome message")
	}

	hasResult := strings.Contains(result, "wins") || strings.Contains(result, "Game Over")
	if !hasResult {
		t.Error("game should have a result")
	}
}

func TestStartGame_AIVsHuman(t *testing.T) {
	input := "2\n1\n1\n2\n3\nn\n"
	reader := strings.NewReader(input)
	var output bytes.Buffer

	bufReader := bufio.NewReader(reader)

	tttio.ShowNewline(&output)
	testGame := BuildGame(bufReader, &output)
	testGame.PlayGame()

	result := output.String()

	if !strings.Contains(result, "Select Player X type") {
		t.Error("should show player X type selection")
	}

	if !strings.Contains(result, "Select Player O type") {
		t.Error("should show player O type selection")
	}

	if !strings.Contains(result, "Welcome") {
		t.Error("should show welcome message")
	}

	hasResult := strings.Contains(result, "wins") || strings.Contains(result, "Game Over")
	if !hasResult {
		t.Error("game should have a result")
	}
}

func TestStartGame_AIVsAI(t *testing.T) {
	input := "2\n2\nn\n"
	reader := strings.NewReader(input)
	var output bytes.Buffer

	bufReader := bufio.NewReader(reader)

	tttio.ShowNewline(&output)
	testGame := BuildGame(bufReader, &output)
	testGame.PlayGame()

	tttio.ShowPlayAgainPrompt(&output)
	playAgain, _ := tttio.ReadPlayAgain(bufReader, &output)

	result := output.String()

	if !strings.Contains(result, "Select Player X type") {
		t.Error("should show player X type selection")
	}

	if !strings.Contains(result, "Select Player O type") {
		t.Error("should show player O type selection")
	}

	if !strings.Contains(result, "Welcome") {
		t.Error("should show welcome message")
	}

	if !strings.Contains(result, "Game Over") {
		t.Error("AI vs AI should result in draw")
	}

	if playAgain {
		t.Error("should not want to play again")
	}
}

func TestStartGame_PlayMultipleGames(t *testing.T) {
	input := "1\n1\n1\n4\n2\n5\n3\ny\n1\n1\n1\n4\n2\n5\n9\n6\nn\n"
	reader := strings.NewReader(input)
	var output bytes.Buffer

	bufReader := bufio.NewReader(reader)

	tttio.ShowNewline(&output)
	game1 := BuildGame(bufReader, &output)
	game1.PlayGame()

	tttio.ShowPlayAgainPrompt(&output)
	playAgain1, _ := tttio.ReadPlayAgain(bufReader, &output)

	result1 := output.String()

	if !playAgain1 {
		t.Error("should want to play again after first game")
	}

	if !strings.Contains(result1, "Player X wins") {
		t.Error("first game should show X wins")
	}

	tttio.ShowNewline(&output)
	game2 := BuildGame(bufReader, &output)
	game2.PlayGame()

	tttio.ShowPlayAgainPrompt(&output)
	playAgain2, _ := tttio.ReadPlayAgain(bufReader, &output)

	result2 := output.String()

	if playAgain2 {
		t.Error("should not want to play again after second game")
	}

	if !strings.Contains(result2, "Player O wins") {
		t.Error("second game should show O wins")
	}

	xSelectionCount := strings.Count(result2, "Select Player X type")
	if xSelectionCount != 2 {
		t.Errorf("should select player X type twice, got %d", xSelectionCount)
	}

	oSelectionCount := strings.Count(result2, "Select Player O type")
	if oSelectionCount != 2 {
		t.Errorf("should select player O type twice, got %d", oSelectionCount)
	}
}

func TestStartGame_GoodbyeMessageWhenDecliningReplay(t *testing.T) {
	input := "1\n1\n1\n4\n2\n5\n3\nn\n"
	reader := strings.NewReader(input)
	var output bytes.Buffer

	bufReader := bufio.NewReader(reader)

	tttio.ShowNewline(&output)
	testGame := BuildGame(bufReader, &output)
	testGame.PlayGame()

	tttio.ShowPlayAgainPrompt(&output)
	playAgain, _ := tttio.ReadPlayAgain(bufReader, &output)

	if !playAgain {
		tttio.ShowGoodbye(&output)
	}

	result := output.String()

	if !strings.Contains(result, "Thanks for playing") {
		t.Error("should show goodbye message when declining replay")
	}
}
