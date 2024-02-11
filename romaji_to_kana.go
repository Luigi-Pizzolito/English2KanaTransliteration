package kanatrans

import (
	"github.com/gojp/kana"
	"strings"
	"unicode"
)

type RomajiToKana struct {
	strict bool
}

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