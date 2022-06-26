package main

import (
	"fmt"
	"strconv"
	"strings"
)

func generatePatternList(s string) [][]string {
	runeList := []rune(s)
	sentencePatternList := make([][]string, 0)
	for i := 0; i < len(runeList); i++ {
		sentencePatternList = append(sentencePatternList, make([]string, 0))
		sentencePatternList[i] = make([]string, 0)
		for j := 1; j <= (len(runeList)-i) && j <= 14; j++ {
			sentencePatternList[i] = append(sentencePatternList[i], getPattern(string(runeList[i:i+j])))
		}
	}

	return sentencePatternList
}

func getPattern(s string) string {
	runeArray := []rune(s)

	patternMap := make(map[rune]string)
	keyList := make([]rune, 0)
	for i, char := range runeArray {
		_, exists := patternMap[char]
		if !exists {
			curChar := runeArray[i]
			indexList := make([]string, 0)
			for j := i; j < len(runeArray); j++ {
				if runeArray[j] == curChar {
					indexList = append(indexList, strconv.Itoa(j))
				}
			}
			if len(indexList) > 1 {
				patternMap[curChar] = strings.Join(indexList, "-")
				keyList = append(keyList, curChar)
			}
		}
	}

	outString := strconv.Itoa(len(runeArray))
	if len(keyList) > 0 {
		patternList := make([]string, 0)
		for _, key := range keyList {
			patternList = append(patternList, patternMap[key])
		}
		outString = fmt.Sprintf("%s;%s", outString, strings.Join(patternList, ","))
	}

	return outString
}
