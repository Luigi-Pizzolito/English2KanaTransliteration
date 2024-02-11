package main

import (
	"strings"
)

// Function to convert normal punctuations to their Japanese equivalents
func convertToJapanesePunctuation(str string) string {
	punctuationMap := map[rune]string{
		',': "、",
		'.': "。",
		'!': "！",
		'?': "？",
		';': "；",
		':': "：",
		'-': "ー",
		'~': "〜",
		'\'': "’",
		' ': "　",
	}

	// Iterate over each character in the string
	var result strings.Builder
	for _, char := range str {
		// Check if the character is in the punctuation map
		if repl, ok := punctuationMap[char]; ok {
			// If yes, append its Japanese equivalent to the result
			result.WriteString(repl)
		} else {
			// If not, append the original character to the result
			result.WriteRune(char)
		}
	}

	return result.String()
}

// Function to convert normal punctuations to their Japanese equivalents
func convertToJapanesePunctuationRestricted(str string) string {
	punctuationMap := map[rune]string{
		',': "、",
		'.': "。",
		'!': "、",
		'?': "、",
		';': "、",
		':': "、",
		'-': "、",
		'~': "、",
		'\'': "、",
		' ': "",
		'　': "",
	}

	// Iterate over each character in the string
	var result strings.Builder
	for _, char := range str {
		// Check if the character is in the punctuation map
		if repl, ok := punctuationMap[char]; ok {
			// If yes, append its Japanese equivalent to the result
			result.WriteString(repl)
		} else {
			// If not, append the original character to the result
			result.WriteRune(char)
		}
	}

	return strings.TrimSpace(result.String())
}