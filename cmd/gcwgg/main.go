package main

import (
	"bufio"
	"math/rand"
	"os"
	"strings"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/gcwgg/internal"
)

func main() {
	var words []string
	var err error

	if len(os.Args) < 2 {
		words = internal.GetFallbackData()
	} else {
		words, err = internal.ReadFileToLines(os.Args[1])
		if err != nil {
			color.PrintRedln("Error while reading file: " + err.Error())
			color.PrintRedln("Using fallback words instead.")
			words = internal.GetFallbackData()
		}
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		randNum := rand.Int31n(int32(len(words) - 1))
		randword := words[randNum]
		internal.GameLoop(randword, reader)

		color.PrintBlue("Play again? [Y/n]: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			color.PrintRedln("Error while reading input: " + err.Error())
			return
		}

		if strings.ToLower(strings.TrimSuffix(input, "\n")) == "y" {
			continue
		}

		break
	}
}
