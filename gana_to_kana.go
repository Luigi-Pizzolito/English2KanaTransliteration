package kanatrans

import (
	"strings"
	"unicode"
)

// HiraganaToKatakana is a struct representing the Hiragana to Katakana converter
type HiraganaToKatakana struct{}

// NewHiraganaToKatakana creates a new instance of HiraganaToKatakana
func NewHiraganaToKatakana() *HiraganaToKatakana {
	return &HiraganaToKatakana{}
}

// Convert converts Hiragana characters to Katakana while leaving Katakana characters unchanged
func (hk *HiraganaToKatakana) Convert(input string) string {
	var result strings.Builder
	for _, char := range input {
		if hk.isHiragana(char) {
			// Convert Hiragana to Katakana
			result.WriteRune(char + 'ァ' - 'ぁ')
		} else {
			// Keep Katakana characters unchanged
			result.WriteRune(char)
		}
	}
	return result.String()
}

// isHiragana checks if a rune is a Hiragana character
func (hk *HiraganaToKatakana) isHiragana(r rune) bool {
	// return r >= 'ぁ' && r <= 'ゖ'
	return unicode.Is(unicode.Hiragana, r)
}