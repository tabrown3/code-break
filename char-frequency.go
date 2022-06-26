package main

import (
	"fmt"
	"sort"
)

type CharFrequency struct {
	Char      rune
	Frequency int
}

type CharFrequencies = []CharFrequency

func printCharFrequency(s string) {
	charCountMap := make(map[rune]int)
	for _, char := range s {
		_, charCountOk := charCountMap[char]

		if !charCountOk {
			charCountMap[char] = 1
		} else {
			charCountMap[char]++
		}
	}

	charFrequencies := make(CharFrequencies, 0)
	for char, charCount := range charCountMap {
		charFrequencies = append(charFrequencies, CharFrequency{Char: char, Frequency: charCount})
	}

	sort.Slice(charFrequencies, func(i, j int) bool {
		return charFrequencies[i].Frequency > charFrequencies[j].Frequency
	})

	for _, charFrequency := range charFrequencies {
		fmt.Printf("%c:%d ", charFrequency.Char, charFrequency.Frequency)
	}
}
