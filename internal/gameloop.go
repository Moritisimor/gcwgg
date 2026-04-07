package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GameLoop(toGuess string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading input: " + err.Error())
		}

		input = strings.TrimSuffix(input, "\n")
		if input == toGuess {
			fmt.Println("You win!")
			return
		} else {
			fmt.Println("You lose!")
		}
	}
}
