package internal

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
)

func GameLoop(toGuess string, reader *bufio.Reader) {
	attempts := 0
	fmt.Printf(
		"%s %s %s",
		color.SprintBlue("The word is"),
		color.SprintGreen(strconv.FormatInt(int64(len(toGuess)), 10)),
		color.SprintBlue("letters long\n"),
	)

	for {
		attempts++
		color.PrintBlue("Enter your guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			color.PrintRedln("Error while reading input: " + err.Error())
			return
		}

		input = strings.TrimSuffix(input, "\n")
		if RevealLexicalProximity(toGuess, input) {
			color.PrintGreenln("You win!")
			color.PrintMagentaln(fmt.Sprintf("Solved in %d tries", attempts))
			break
		}
	}
}
