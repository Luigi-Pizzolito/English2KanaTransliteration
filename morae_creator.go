package kanatrans

import (
	"unicode/utf8"
	"fmt"
)

// moraeCreator creates morae from phonetics.
type moraeCreator struct {
	vowels string
}

// newMoraeCreator creates a new instance of moraeCreator.
func newMoraeCreator() *moraeCreator {
	return &moraeCreator{
		vowels: "aeiou",
	}
}

// CreateMorae creates morae from phonetics.
func (mc *moraeCreator) CreateMorae(ph string) string {
	result := string(ph[0]) // Add the first character to the result
	for pIdx, r := range ph[1:] {
		// Increment pIdx by 1 to match the index of the current rune
		pIdx++

		// Get previous character
		runeAtIndex, size := utf8.DecodeRuneInString(ph[pIdx-1:])
		if runeAtIndex == utf8.RuneError && size == 0 {
			fmt.Println("Invalid UTF-8 encoding or index out of range")
			return ""
		}

		// Previous character is vowel
		if mc.isVowel(string(runeAtIndex)) {
			result += "." + string(r) // Add a dot before the current character
		} else {
			// Previous character is consonant
			if runeAtIndex == r || runeAtIndex == 'N' {
				result += "." + string(r) // Add a dot before the current character
			} else {
				result += string(r) // Add the current character as is
			}
		}
	}
	return result
}

// isVowel checks if a character is a vowel.
func (mc *moraeCreator) isVowel(char string) bool {
	for _, v := range mc.vowels {
		if string(v) == char {
			return true
		}
	}
	return false
}