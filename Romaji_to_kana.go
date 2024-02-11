package kanatrans

import (
	"github.com/gojp/kana"
	"strings"
	"unicode"
)

// RomajiToKana struct holds the necessary functions for Romaji to Kana conversion
type RomajiToKana struct {
	strict bool
}

// NewRomajiToKana creates a new instance of RomajiToKana
func NewRomajiToKana(strictClean ...bool) *RomajiToKana {
	// Set cleaner with default non-strict cleaning
	var strictF bool
	if len(strictClean) > 0 {
		strictF = strictClean[0]
	}
	return &RomajiToKana{
		strict: strictF,
	}
}

// Convert converts Romaji to Katakana
func (r2k *RomajiToKana) Convert(s string) string {
	out := kana.RomajiToKatakana(kana.NormalizeRomaji(s))

	if !r2k.strict {
		// non-strict mode
		out = convertToJapanesePunctuation(out)
		return out
	}
	
	// strict mode
	out = r2k.removeNonKatakana(out)
	out = convertToJapanesePunctuationRestricted(out)
	return out
}

func (r2k *RomajiToKana) removeNonKatakana(str string) string {
	var katakanaRunes []rune
	for _, r := range str {
		if unicode.Is(unicode.Katakana, r) || r2k.isInPunctuationMap(r) {
			katakanaRunes = append(katakanaRunes, r)
		}
	}
	return string(katakanaRunes)
}

func (r2k *RomajiToKana) isInPunctuationMap(r rune) bool {
	punctuation := ",.!?;:-~' 　’〜ー：；？！。、"
	return strings.ContainsAny(string(r),punctuation)
}