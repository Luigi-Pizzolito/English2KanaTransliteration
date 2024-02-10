package main

import (
	"encoding/json"
	_ "embed"
)

// EngToKana struct holds the necessary functions for English to Katakana conversion
type EngToKana struct {
	db             map[string][]string
	vowelFn        func(string, string) string
	consonantFn    func(string, string) string
	epentheticFn   func(string) string
	moraeFn        func(string) string
	kanaFn         func(string) string
}

//go:embed cmu_ipa.json
var dbFile string

func (e *EngToKana) Init() {
	vowel := NewVowelConverter()
	e.vowelFn = vowel.ConvertVowel
	consonant := NewConsonantConverter()
	e.consonantFn = consonant.ConvertConsonant
	epenthetic := NewEpentheticVowelHandler()
	e.epentheticFn = epenthetic.AddEpentheticVowel
	morae := NewMoraeCreator()
	e.moraeFn = morae.CreateMorae
	kana := NewMoraeKanaConverter()
	e.kanaFn = kana.ConvertMorae
}

// LoadDBFromFile loads the JSON file containing the database.
func (e *EngToKana) LoadDB() error {
	// Unmarshal the JSON data into a map[string][]string
	if err := json.Unmarshal([]byte(dbFile), &e.db); err != nil {
		return err
	}

	return nil
}

// Transcript converts an English word to Katakana
func (e *EngToKana) Transcript(word string) []string {
	phs, ok := e.db[word]
	if !ok {
		return []string{"E_DIC"}
	}

	var result []string
	for _, ph := range phs {
		ph1 := e.vowelFn(word, ph)
		ph2 := e.consonantFn(word, ph1)
		ph3 := e.epentheticFn(ph2)
		morae := e.moraeFn(ph3)
		kana := e.kanaFn(morae)
		result = append(result, kana)
	}
	return result
}

// FromWordList processes a list of words
func (e *EngToKana) FromWordList(words []string) [][]string {
	var result [][]string
	for _, w := range words {
		transcript := e.Transcript(w)
		if len(transcript) > 0 {
			result = append(result, transcript)
		}
	}
	return result
}