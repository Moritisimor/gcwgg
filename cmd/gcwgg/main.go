package main

import (
	"math/rand"
	"os"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/gcwgg/internal"
)

func main() {
	if len(os.Args) < 2 {
		color.PrintRedln("Usage: gcwgg <file>")
		return
	}

	words, err := internal.ReadFileToLines(os.Args[1])
	if err != nil {
		color.PrintRedln("Error while reading file: " + err.Error())
		return
	}

	internal.GameLoop(words[rand.Int31n(int32(len(words) - 1))])
}
