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
				tokens := parseInput(innerUserInput, "=")
				cipherChar, clearChar := rune(tokens[0][0]), rune(tokens[1][0])
				charMap[cipherChar] = clearChar
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
				filteredWords := filterValidWords(words, cipherWord, charMap)
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
					fmt.Println(filterValidWords(words, cipherWord, charMap))
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
				clearChar, _ := waitForInput()
				charMap[cipherChar] = []rune(clearChar)[0]
				break
			}
		}
	}

}

func parseInput(s string, sep string) []string {
	return strings.Split(s, sep)
}

func filterValidWords(words []string, cipherWord []rune, charMap map[rune]rune) []string {
	filteredWords := make([]string, 0)
	for i := 0; i < len(words); i++ {
		word := []rune(words[i])
		isInvalid := false
		for j := 0; j < len(word); j++ {
			mappedChar, exists := charMap[cipherWord[j]]
			wordChar := word[j]

			if exists && mappedChar != '_' && mappedChar != wordChar {
				// if the word has a letter at this position that doesn't match the current guess,
				// it's an invalid suggestion
				isInvalid = true
				break
			}
		}
		if !isInvalid {
			filteredWords = append(filteredWords, string(word))
		}
	}
	return filteredWords
}
