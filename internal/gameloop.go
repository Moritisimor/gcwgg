package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
)

func GameLoop(toGuess string) {
	fmt.Println(toGuess)
	reader := bufio.NewReader(os.Stdin)
	attempts := 0
	for {
		attempts++
		color.PrintBlue("Enter your guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			color.PrintRedln("Error while reading input: " + err.Error())
		}

		input = strings.TrimSuffix(input, "\n")
		if RevealLexicalProximity(toGuess, input) {
			color.PrintGreenln("You win!")
			color.PrintMagentaln(fmt.Sprintf("Solved in %d tries", attempts))
			return
		}
	}
}
