package internal

import (
	"bufio"
	"os"
	"strings"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
)

func GameLoop(toGuess string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		color.PrintBlue("Enter your guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			color.PrintRedln("Error while reading input: " + err.Error())
		}

		input = strings.TrimSuffix(input, "\n")
		if RevealProgress(toGuess, input) {
			color.PrintGreenln("You win!")
			return
		}
	}
}
