package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func startCryptogram(cipherTexts []string) {
	patternLists := make([][][]string, 0)
	for i := 0; i < len(cipherTexts); i++ {
		patternLists = append(patternLists, generatePatternList(cipherTexts[i]))
	}

	patternMap := getPatternMap()
	cipherToClearCharMap := make(map[rune]rune)
	clearToCipherCharMap := make(map[rune]rune)
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
					mappedChar, ok := cipherToClearCharMap[char]
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
				cipherToClearCharMap = make(map[rune]rune)
				clearToCipherCharMap = make(map[rune]rune)
				break
			} else if userInput == "assign" {
				innerUserInput, _ := waitForInput()
				tokens := parseInput(innerUserInput, "=")
				cipherChar, clearChar := rune(tokens[0][0]), rune(tokens[1][0])
				if cipherChar == '_' {
					oldClearChar := cipherToClearCharMap[cipherChar]
					delete(cipherToClearCharMap, cipherChar)
					delete(clearToCipherCharMap, oldClearChar)
				} else {
					cipherToClearCharMap[cipherChar] = clearChar
					clearToCipherCharMap[clearChar] = cipherChar
				}
				break
			} else if userInput == "pattern-at" {
				innerUserInput, _ := waitForInput()
				tokens := parseInput(innerUserInput, " ")
				inputSentenceIndex, inputCharIndex, inputWidth := tokens[0], tokens[1], tokens[2]
				indexSentence, _ := strconv.Atoi(inputSentenceIndex)
				indexChar, _ := strconv.Atoi(inputCharIndex)
				width, _ := strconv.Atoi(inputWidth)
				fmt.Println(patternLists[indexSentence][indexChar][width])
			} else if userInput == "patterns-at" {
				innerUserInput, _ := waitForInput()
				tokens := parseInput(innerUserInput, " ")
				inputSentenceIndex, inputCharIndex := tokens[0], tokens[1]
				sentenceIndex, _ := strconv.Atoi(inputSentenceIndex)
				charIndex, _ := strconv.Atoi(inputCharIndex)
				fmt.Println(patternLists[sentenceIndex][charIndex])
			} else if userInput == "words-at" {
				innerUserInput, _ := waitForInput()
				tokens := parseInput(innerUserInput, " ")
				inputSentenceIndex, inputCharIndex, inputWidth := tokens[0], tokens[1], tokens[2]
				indexSentence, _ := strconv.Atoi(inputSentenceIndex)
				indexChar, _ := strconv.Atoi(inputCharIndex)
				width, _ := strconv.Atoi(inputWidth)
				words := patternMap[patternLists[indexSentence][indexChar][width]]
				cipherWord := []rune(cipherTexts[indexSentence])[indexChar : indexChar+width+1]
				filteredWords := filterValidWords(words, cipherWord, cipherToClearCharMap, clearToCipherCharMap)
				fmt.Print(filteredWords)
				fmt.Println()
			} else if userInput == "all-words-at" {
				innerUserInput, _ := waitForInput()
				tokens := parseInput(innerUserInput, " ")
				indexSentenceInput, indexCharInput := tokens[0], tokens[1]
				indexSentence, _ := strconv.Atoi(indexSentenceInput)
				indexChar, _ := strconv.Atoi(indexCharInput)
				for width, pattern := range patternLists[indexSentence][indexChar] {
					words := patternMap[pattern]
					cipherWord := []rune(cipherTexts[indexSentence])[indexChar : indexChar+width+1]
					fmt.Println(filterValidWords(words, cipherWord, cipherToClearCharMap, clearToCipherCharMap))
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
				tokens := parseInput(innerUserInput, " ")
				indexSentenceInput, indexCharInput := tokens[0], tokens[1]
				indexSentence, _ := strconv.Atoi(indexSentenceInput)
				indexChar, _ := strconv.Atoi(indexCharInput)
				cipherChar := []rune(cipherTexts[indexSentence])[indexChar]
				clearString, _ := waitForInput()
				clearChar := []rune(clearString)[0]
				if cipherChar == '_' {
					oldClearChar := cipherToClearCharMap[cipherChar]
					delete(cipherToClearCharMap, cipherChar)
					delete(clearToCipherCharMap, oldClearChar)
				} else {
					cipherToClearCharMap[cipherChar] = clearChar
					clearToCipherCharMap[clearChar] = cipherChar
				}
				break
			}
		}
	}

}

func parseInput(s string, sep string) []string {
	return strings.Split(s, sep)
}

func filterValidWords(clearWords []string, cipherWord []rune, cipherToClearCharMap map[rune]rune, clearToCipherCharMap map[rune]rune) []string {
	filteredWords := make([]string, 0)
	for i := 0; i < len(clearWords); i++ {
		clearWord := []rune(clearWords[i])
		isInvalid := false
		for j := 0; j < len(clearWord); j++ {
			cipherChar := cipherWord[j]
			wordChar := clearWord[j]

			if cipherChar == wordChar {
				// in a cryptogram, a letter cannot stand for itself (e.g. assign G=G is invalid)
				isInvalid = true
				break
			}

			guessChar, positionHasGuessAssigned := cipherToClearCharMap[cipherWord[j]]

			if positionHasGuessAssigned && guessChar != wordChar {
				// if the word has a letter at this position that doesn't match the current guess,
				// it's an invalid suggestion
				isInvalid = true
				break
			}

			_, neededCharAlreadyAssigned := clearToCipherCharMap[wordChar]

			if !positionHasGuessAssigned && neededCharAlreadyAssigned {
				// if the word needs a letter at this position that has already been assigned as
				// a guessed clear text substitution, the suggestion is invalid
				isInvalid = true
				break
			}
		}
		if !isInvalid {
			filteredWords = append(filteredWords, string(clearWord))
		}
	}
	return filteredWords
}
