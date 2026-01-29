package io

import (
	"bytes"
	"strings"
	"testing"
	"ttt/boards"
)

func TestShowWelcome_WelcomeMessage(t *testing.T) {
	var output bytes.Buffer

	ShowWelcome(&output)

	result := output.String()
	if !strings.Contains(result, "Welcome") {
		t.Error("should display welcoming message")
	}
}

func TestShowWelcome_GameName(t *testing.T) {
	var output bytes.Buffer

	ShowWelcome(&output)

	result := output.String()
	if !strings.Contains(result, "Tic-Tac-Toe") {
		t.Error("should show game name")
	}
}

func TestShowBoard_EmptyBoard(t *testing.T) {
	var output bytes.Buffer
	board := boards.NewBoard()

	ShowBoard(&output, board)

	result := output.String()
	for position := 1; position <= 9; position++ {
		posStr := string(rune('0' + position))
		if !strings.Contains(result, posStr) {
			t.Errorf("should display position %d", position)
		}
	}
}

func TestShowBoard_BoardWithMoves(t *testing.T) {
	var output bytes.Buffer
	board := boards.Board{
		{"X", "2", "O"},
		{"4", "X", "6"},
		{"O", "8", "9"},
	}

	ShowBoard(&output, board)

	result := output.String()
	expectedContent := []string{"X", "O", "2", "4", "6", "8", "9"}
	for _, content := range expectedContent {
		if !strings.Contains(result, content) {
			t.Errorf("should display %s", content)
		}
	}
}

func TestShowBoard_FullBoard(t *testing.T) {
	var output bytes.Buffer
	board := boards.Board{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"O", "X", "X"},
	}

	ShowBoard(&output, board)

	result := output.String()
	xCount := strings.Count(result, "X")
	oCount := strings.Count(result, "O")

	if xCount < 5 {
		t.Errorf("should display all X's, found %d", xCount)
	}
	if oCount < 4 {
		t.Errorf("should display all O's, found %d", oCount)
	}
}

func TestShowBoard_GridStructure(t *testing.T) {
	board := boards.NewBoard()
	var output bytes.Buffer

	ShowBoard(&output, board)

	result := output.String()
	if !strings.Contains(result, "|") {
		t.Error("should include column separators")
	}

	if !strings.Contains(result, "-") {
		t.Error("should include row separators")
	}
}

func TestShowPlayerTurn_XTurn(t *testing.T) {
	var output bytes.Buffer

	ShowPlayerTurn(&output, "X")

	result := output.String()
	if !strings.Contains(result, "X") {
		t.Error("should display X")
	}
	if !strings.Contains(result, "turn") {
		t.Error("should show turn")
	}
}

func TestShowPlayerTurn_OTurn(t *testing.T) {
	var output bytes.Buffer

	ShowPlayerTurn(&output, "O")

	result := output.String()
	if !strings.Contains(result, "O") {
		t.Error("should display O")
	}
	if !strings.Contains(result, "turn") {
		t.Error("should show turn")
	}
}

func TestShowPrompt_AsksForMove(t *testing.T) {
	var output bytes.Buffer

	ShowPrompt(&output)

	result := output.String()
	if !strings.Contains(result, "Enter") || !strings.Contains(result, "move") {
		t.Error("should prompt user to enter a move")
	}
}

func TestShowPrompt_ValidRange(t *testing.T) {
	var output bytes.Buffer

	ShowPrompt(&output)

	result := output.String()
	if !strings.Contains(result, "1") || !strings.Contains(result, "9") {
		t.Error("should show valid position range 1-9")
	}
}

func TestShowPositionTaken_DisplaysMessage(t *testing.T) {
	var output bytes.Buffer

	ShowPositionTaken(&output)

	result := output.String()
	if !strings.Contains(result, "Position") || !strings.Contains(result, "taken") {
		t.Error("should show position is already occupied")
	}
}

func TestShowWinner_XVictory(t *testing.T) {
	var output bytes.Buffer

	ShowWinner(&output, "X")

	result := output.String()
	if !strings.Contains(result, "X") {
		t.Error("should display X")
	}
	if !strings.Contains(result, "win") {
		t.Error("should show win")
	}
}

func TestShowWinner_OVictory(t *testing.T) {
	var output bytes.Buffer

	ShowWinner(&output, "O")

	result := output.String()
	if !strings.Contains(result, "O") {
		t.Error("should display O")
	}
	if !strings.Contains(result, "win") {
		t.Error("should show win")
	}
}

func TestShowDraw_GameOver(t *testing.T) {
	var output bytes.Buffer

	ShowDraw(&output)

	result := output.String()
	if !strings.Contains(result, "Game Over") {
		t.Error("should show game over")
	}
}

func TestShowDraw_BoardFull(t *testing.T) {
	var output bytes.Buffer

	ShowDraw(&output)

	result := output.String()
	if !strings.Contains(result, "full") {
		t.Error("should be full board")
	}
}

func TestShowPlayerTypeSelection_DisplaysOptionsForX(t *testing.T) {
	var output bytes.Buffer

	ShowPlayerTypeSelection(&output, "X")

	result := output.String()
	requiredContent := []string{"X", "Human", "AI", "1", "2"}
	for _, content := range requiredContent {
		if !strings.Contains(result, content) {
			t.Errorf("should include %q", content)
		}
	}
}

func TestShowPlayerTypeSelection_DisplaysOptionsForO(t *testing.T) {
	var output bytes.Buffer

	ShowPlayerTypeSelection(&output, "O")

	result := output.String()
	requiredContent := []string{"O", "Human", "AI", "1", "2"}
	for _, content := range requiredContent {
		if !strings.Contains(result, content) {
			t.Errorf("should include %q", content)
		}
	}
}

func TestShowPlayAgainPrompt_AsksPlayAgain(t *testing.T) {
	var output bytes.Buffer

	ShowPlayAgainPrompt(&output)

	result := output.String()
	if !strings.Contains(result, "Play again") {
		t.Error("should ask if user wants to play again")
	}
}

func TestShowPlayAgainPrompt_ShowsYesNoOptions(t *testing.T) {
	var output bytes.Buffer

	ShowPlayAgainPrompt(&output)

	result := output.String()
	if !strings.Contains(result, "y") && !strings.Contains(result, "n") {
		t.Error("should show y/n options")
	}
}

func TestShowGoodbye_DisplaysGoodBye(t *testing.T) {
	var output bytes.Buffer

	ShowGoodbye(&output)

	result := output.String()
	if !strings.Contains(result, "Thanks") || !strings.Contains(result, "playing") {
		t.Error("should thank user for playing")
	}
}
