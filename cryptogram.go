package main

import (
	"fmt"
	"os"
	"strconv"
)

func startCryptogram(cipherTexts []string) {
	patternLists := make([][][]string, 0)
	for i := 0; i < len(cipherTexts); i++ {
		patternLists = append(patternLists, generatePatternList(cipherTexts[i]))
	}

	patternMap := getPatternMap()
	charMap := make(map[rune]rune)
	for {
		// INITIAL PRINTOUT
		for i := 0; i < len(cipherTexts); i++ {
			fmt.Printf("%d: %s\n", i, cipherTexts[i])
		}
		for i := 0; i < len(cipherTexts); i++ {
			fmt.Printf("%d: ", i)
			for _, char := range cipherTexts[i] {
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
		}

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
				cipherChar, clearChar := parseUserInput2(innerUserInput)
				charMap[cipherChar] = clearChar
				break
			} else if userInput == "pattern-at" {
				innerUserInput, _ := waitForInput()
				inputSentenceIndex, inputCharIndex, inputWidth := parseUserInput3(innerUserInput)
				indexSentence, _ := strconv.Atoi(string(inputSentenceIndex))
				indexChar, _ := strconv.Atoi(string(inputCharIndex))
				width, _ := strconv.Atoi(string(inputWidth))
				fmt.Println(patternLists[indexSentence][indexChar][width])
			} else if userInput == "patterns-at" {
				innerUserInput, _ := waitForInput()
				inputSentenceIndex, inputCharIndex := parseUserInput2(innerUserInput)
				sentenceIndex, _ := strconv.Atoi(string(inputSentenceIndex))
				charIndex, _ := strconv.Atoi(string(inputCharIndex))
				fmt.Println(patternLists[sentenceIndex][charIndex])
			} else if userInput == "words-at" {
				innerUserInput, _ := waitForInput()
				inputSentenceIndex, inputCharIndex, inputWidth := parseUserInput3(innerUserInput)
				indexSentence, _ := strconv.Atoi(string(inputSentenceIndex))
				indexChar, _ := strconv.Atoi(string(inputCharIndex))
				width, _ := strconv.Atoi(string(inputWidth))
				fmt.Println(patternMap[patternLists[indexSentence][indexChar][width]])
			} else if userInput == "all-words-at" {
				innerUserInput, _ := waitForInput()
				indexSentenceInput, indexCharInput := parseUserInput2(innerUserInput)
				indexSentence, _ := strconv.Atoi(string(indexSentenceInput))
				indexChar, _ := strconv.Atoi(string(indexCharInput))
				for _, pattern := range patternLists[indexSentence][indexChar] {
					fmt.Println(patternMap[pattern])
				}
			} else if userInput == "indices-for-word" {
				innerUserInput, _ := waitForInput()
				inputPattern := getPattern(innerUserInput)
				sentenceIndices := make([][]int, 0)
				for i := 0; i < len(patternLists); i++ {
					sentenceIndices = append(sentenceIndices, make([]int, 0))
					for j, patterns := range patternLists[i] {
						windowLenIndex := len(innerUserInput) - 1
						if len(patterns) > windowLenIndex {
							if patterns[windowLenIndex] == inputPattern {
								sentenceIndices[i] = append(sentenceIndices[i], j)
							}
						}
					}
				}
				fmt.Println(sentenceIndices)
			} else if userInput == "assign-nth-char" {
				innerUserInput, _ := waitForInput()
				indexSentenceInput, indexCharInput := parseUserInput2(innerUserInput)
				indexSentence, _ := strconv.Atoi(string(indexSentenceInput))
				indexChar, _ := strconv.Atoi(string(indexCharInput))
				cipherChar := []rune(cipherTexts[indexSentence])[indexChar]
				clearChar, _ := waitForInput()
				charMap[cipherChar] = []rune(clearChar)[0]
				break
			}
		}
	}

}

func parseUserInput2(s string) (rune, rune) {
	runes := []rune(s)
	return runes[0], runes[2]
}

func parseUserInput3(s string) (rune, rune, rune) {
	runes := []rune(s)
	return runes[0], runes[2], runes[4]
}
