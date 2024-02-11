package kanatrans

import (
	"strings"
	"unicode"
)

// HiraganaToKana struct holds the necessary functions for Hiragana to Katakana conversion
type HiraganaToKana struct{}

// NewHiraganaToKana creates a new instance of HiraganaToKana
func NewHiraganaToKana() *HiraganaToKana {
	return &HiraganaToKana{}
}

// Convert converts Hiragana characters to Katakana while leaving Katakana characters unchanged
func (h2k *HiraganaToKana) Convert(input string) string {
	var result strings.Builder
	for _, char := range input {
		if h2k.isHiragana(char) {
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
func (h2k *HiraganaToKana) isHiragana(r rune) bool {
	// return r >= 'ぁ' && r <= 'ゖ'
	return unicode.Is(unicode.Hiragana, r)
}