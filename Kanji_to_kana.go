package kanatrans

import (
	"strings"
	"encoding/json"
	_ "embed"
	"fmt"
)

//go:embed dict/kanjidic2_pronounce.json
var kanjiMapFile string

// KanjiToKana struct holds the necessary functions for Kanji to Kana conversion
type KanjiToKana struct {
	kanaMap map[string]string
}

// NewKanjiToKana creates a new instance of KanjiToKana
func NewKanjiToKana() *KanjiToKana {
	k2k := KanjiToKana{}
	if err := json.Unmarshal([]byte(kanjiMapFile), &k2k.kanaMap); err != nil {
		fmt.Println("ERROR: Loading kanjidic2_pronounce.json")
		return nil
	}
	return &k2k
}

// Convert converts Kanji into Katakana
func (k2k *KanjiToKana) Convert(kanji string) string {
	var kana strings.Builder
	for _, char := range kanji {
		// Check if the character exists in the map
		if val, ok := k2k.kanaMap[string(char)]; ok {
			kana.WriteString(val) // If yes, append the corresponding kana to the result
		} else {
			// If no mapping found, skip the character
		}
	}
	return kana.String()
}