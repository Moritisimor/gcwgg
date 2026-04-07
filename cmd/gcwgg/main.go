package main

import (
	"math/rand"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/gcwgg/internal"
)

func main() {
	words, err := internal.ReadFileToLines("words.txt")
	if err != nil {
		color.PrintRedln("Error while reading file: " + err.Error())
		return
	}

	internal.GameLoop(words[rand.Int31n(int32(len(words) - 1))])
}
