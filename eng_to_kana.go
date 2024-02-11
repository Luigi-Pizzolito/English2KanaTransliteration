package kanatrans

import (
	"strings"
	"encoding/json"
	_ "embed"
)

// EngToKana struct holds the necessary functions for English to Katakana conversion
type EngToKana struct {
	db             map[string][]string
	cleanFn		   func(string,func(string)string) string
	vowelFn        func(string, string) string
	consonantFn    func(string, string) string
	epentheticFn   func(string) string
	moraeFn        func(string) string
	kanaFn         func(string) string
	recoverFn	   func(string) string
}

//go:embed dict/cmu_ipa.json
var dbFile string

func NewEngToKana(strictClean ...bool) *EngToKana {
	// Instantiate class
	e := EngToKana{}

	// Set english cleaner with default non-strict cleaning
	var strictF bool
	if len(strictClean) > 0 {
		strictF = strictClean[0]
	}
	clean := NewEnglishCleaner(strictF)

	// Set other classes and link function pointers
	e.cleanFn = clean.Clean
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
	r2k := NewRomajiToKana(strictF)
	e.recoverFn = r2k.Convert

	// Load cmu_ipa english phoneme pronounce dictionary
	e.loadDB()

	// Return instance
	return &e
}

// LoadDBFromFile loads the JSON file containing the database.
func (e *EngToKana) loadDB() error {
	// Unmarshal the JSON data into a map[string][]string
	if err := json.Unmarshal([]byte(dbFile), &e.db); err != nil {
		return err
	}

	return nil
}

// Transcript converts an English word to Katakana
func (e *EngToKana) TranscriptWord(word string) string {
	phs, ok := e.db[word]
	if !ok {
		// return "E_DIC"
		// If no match found, try to recover by using Romaji2Kana
		return e.recoverFn(word)
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
	return result[0]
}

func (e *EngToKana) TranscriptSentence(line string) string {
	// Clean string with call back to process clean sentence fragment
	return e.cleanFn(line, e.transcriptCleanSentenceFragment)
}

// processes a line of text containing English words into Kana
func (e *EngToKana) transcriptCleanSentenceFragment(line string) string {
	var result strings.Builder
	words := strings.Fields(line)

	// Iterate over the words and convert each to Katakana
	for _, word := range words {
		katakanaWords := e.TranscriptWord(word)
		// Placeholder logic to join Katakana words
		result.WriteString(katakanaWords)

		// TODO: Recover func call in case of E_DIC
		// TODO: Add and try romaji_to_kana.go
	}
	return result.String()
}