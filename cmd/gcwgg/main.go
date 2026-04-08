package main

import (
	"math/rand"
	"os"

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

	internal.GameLoop(words[rand.Int31n(int32(len(words) - 1))])
}
