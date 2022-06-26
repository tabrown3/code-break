package main

import (
	"bufio"
	"fmt"
	"os"
)

func isLatin(char rune) bool {
	return char >= 'A' && char <= 'Z'
}

func waitForInput() (string, error) {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	sentence, err := buf.ReadBytes('\n')
	if err != nil {
		return "", err
	} else {
		return string(sentence[:len(sentence)-2]), nil
	}
}
