package main

import (
	"os"
)

func main() {
	output := NewConsoleOutput(os.Stdout)
	output.ShowWelcome()
	startGame(output)
}
