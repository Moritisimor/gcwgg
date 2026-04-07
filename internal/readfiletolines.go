package internal

import (
	"os"
	"strings"
)

func ReadFileToLines(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(content), "\n"), nil
}
