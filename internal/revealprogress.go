package internal

import (
	"fmt"
	"slices"
	"time"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
)

func RevealProgress(actual, guess string) bool {
	var actualAsSlice []string
	for _, c := range actual {
		actualAsSlice = append(actualAsSlice, string(c))
	}

	var lettersLeft []string
	for _, c := range guess {
		lettersLeft = append(lettersLeft, string(c))
	}

	for i, c := range guess {
		time.Sleep(time.Millisecond * 100)
		letter := string(c)
		if len(actualAsSlice) <= i {
			if slices.Contains(actualAsSlice, letter) && slices.Contains(lettersLeft, letter) {
				color.PrintYellow(letter)
			} else {
				color.PrintRed(letter)
			}

			continue
		}

		letterIndex := slices.Index(lettersLeft, letter)
		if letter == actualAsSlice[i] {
			color.PrintGreen(letter)
			lettersLeft = slices.Delete(lettersLeft, letterIndex, letterIndex + 1)
		} else if slices.Contains(actualAsSlice, letter) && slices.Contains(lettersLeft, letter) {
			lettersLeft = slices.Delete(lettersLeft, letterIndex, letterIndex + 1)
			color.PrintYellow(letter)
		} else {
			color.PrintRed(letter)
		}
	}

	fmt.Println()
	if actual == guess {
		return true
	}

	return false
}
