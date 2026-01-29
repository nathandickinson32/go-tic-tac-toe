package players

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
	"ttt/boards"
	tttio "ttt/io"
)

type HumanPlayer struct {
	reader *bufio.Reader
	output io.Writer
}

func NewHumanPlayer(reader *bufio.Reader, output io.Writer) *HumanPlayer {
	return &HumanPlayer{
		reader: reader,
		output: output,
	}
}

func (humanPlayer *HumanPlayer) isPositionAvailable(board boards.Board, position int) bool {
	return board.IsPositionValid(position)
}

var (
	ErrEmptyInput = errors.New("Input cannot be empty")
	ErrNotNumber  = errors.New("Input must be a number")
	ErrOutOfRange = errors.New("Position must be between 1 and 9")
)

func (humanPlayer *HumanPlayer) parseInput(input string) (int, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return 0, ErrEmptyInput
	}

	position, err := strconv.Atoi(input)
	if err != nil {
		return 0, ErrNotNumber
	}

	if position < boards.MinPosition || position > boards.MaxPosition {
		return 0, ErrOutOfRange
	}

	return position, nil
}

func (humanPlayer *HumanPlayer) getValidPosition() (int, error) {
	for {
		tttio.ShowPrompt(humanPlayer.output)
		line, err := humanPlayer.reader.ReadString('\n')
		if err != nil {
			return 0, err
		}

		position, err := humanPlayer.parseInput(line)
		if err != nil {
			tttio.ShowInvalidInput(humanPlayer.output, err)
			continue
		}

		return position, nil
	}
}

func (humanPlayer *HumanPlayer) ReadMove(board boards.Board) (int, error) {
	for {
		position, err := humanPlayer.getValidPosition()
		if err != nil {
			return 0, err
		}

		if humanPlayer.isPositionAvailable(board, position) {
			return position, nil
		}

		tttio.ShowPositionTaken(humanPlayer.output)
	}
}
