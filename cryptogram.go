package main

import (
	"fmt"
	"os"
	"strconv"
)

func startCryptogram(cipherText string) {
	patternList := generatePatternList(cipherText)
	patternMap := getPatternMap()

	charMap := make(map[rune]rune)
	for {
		// INITIAL PRINTOUT
		fmt.Println(cipherText)

		for _, char := range cipherText {
			if !isLatin(char) {
				fmt.Print(string(char))
			} else {
				mappedChar, ok := charMap[char]
				if !ok {
					fmt.Print("_")
				} else {
					fmt.Printf(string(mappedChar))
				}
			}
		}
		fmt.Println()

		// START USER INPUT
		for {
			userInput, _ := waitForInput()
			if userInput == "quit" {
				os.Exit(0)
			} else if userInput == "reset" {
				charMap = make(map[rune]rune)
				break
			} else if userInput == "assign" {
				innerUserInput, _ := waitForInput()
				cipherChar, clearChar := parseUserInput(innerUserInput)
				charMap[cipherChar] = clearChar
				break
			} else if userInput == "pattern-at" {
				innerUserInput, _ := waitForInput()
				inputIndex, inputWidth := parseUserInput(innerUserInput)
				index, _ := strconv.Atoi(string(inputIndex))
				width, _ := strconv.Atoi(string(inputWidth))
				fmt.Println(patternList[index][width])
			} else if userInput == "patterns-at" {
				innerUserInput, _ := waitForInput()
				index, _ := strconv.Atoi(innerUserInput)
				fmt.Println(patternList[index])
			} else if userInput == "words-at" {
				innerUserInput, _ := waitForInput()
				inputIndex, inputWidth := parseUserInput(innerUserInput)
				index, _ := strconv.Atoi(string(inputIndex))
				width, _ := strconv.Atoi(string(inputWidth))
				fmt.Println(patternMap[patternList[index][width]])
			} else if userInput == "all-words-at" {
				innerUserInput, _ := waitForInput()
				index, _ := strconv.Atoi(innerUserInput)
				for _, pattern := range patternList[index] {
					fmt.Println(patternMap[pattern])
				}
			} else if userInput == "indices-for-word" {
				innerUserInput, _ := waitForInput()
				inputPattern := getPattern(innerUserInput)
				indices := make([]int, 0)
				for i, patterns := range patternList {
					windowLenIndex := len(innerUserInput) - 1
					if len(patterns) > windowLenIndex {
						if patterns[windowLenIndex] == inputPattern {
							indices = append(indices, i)
						}
					}
				}
				fmt.Println(indices)
			} else if userInput == "assign-nth-char" {
				innerUserInput, _ := waitForInput()
				index, _ := strconv.Atoi(innerUserInput)
				cipherChar := []rune(cipherText)[index]
				clearChar, _ := waitForInput()
				charMap[cipherChar] = []rune(clearChar)[0]
				break
			}
		}
	}

}

func parseUserInput(s string) (rune, rune) {
	runes := []rune(s)
	return runes[0], runes[2]
}
