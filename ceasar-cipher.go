package main

func offsetCharsBy(s string, offset int32) string {
	offsetString := ""
	for _, char := range s {
		var offsetChar = char
		if isLatin(char) {
			offsetChar = char + offset
			if offsetChar > 'Z' {
				offsetChar -= 26
			}
		}
		offsetString += string(offsetChar)
	}

	return offsetString
}
