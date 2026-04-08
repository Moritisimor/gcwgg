package internal

import (
	"fmt"
	"sync"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
)

func RevealLexicalProximity(actual, guess string) bool {
	occurrences := countLetterOccurrences(actual)
	var actualAsSlice []string
	var guessAsSlice []string
	wg := sync.WaitGroup{}

	// I Love parallelism
	wg.Go(func() {
		for _, c := range actual {
			actualAsSlice = append(actualAsSlice, string(c))
		}
	})

	wg.Go(func() {
		for _, c := range guess {
			guessAsSlice = append(guessAsSlice, string(c))
		}
	})

	wg.Wait()
	for i, c := range guessAsSlice {
		if len(actualAsSlice) < i + 1 {
			if occurrences[c] > 0 {
				guessAsSlice[i] = color.SprintYellow(c)
			}

			occurrences[c]--
			continue
		}

		if c == actualAsSlice[i] {
			guessAsSlice[i] = color.SprintGreen(c)
			occurrences[c]--
		}
	}

	for i, c := range guessAsSlice {
		if occurrences[c] > 0 {
			guessAsSlice[i] = color.SprintYellow(c)
		}

		occurrences[c]--
	}

	for _, c := range guessAsSlice {
		fmt.Print(c)
	}

	println()
	if actual == guess {
		return true
	}

	return false
}
