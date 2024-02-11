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

// NewEngToKana creates a new instance of EngToKana
func NewEngToKana(strictClean ...bool) *EngToKana {
	// Instantiate class
	e2k := EngToKana{}

	// Set english cleaner with default non-strict cleaning
	var strictF bool
	if len(strictClean) > 0 {
		strictF = strictClean[0]
	}
	clean := newEnglishCleaner(strictF)

	// Set other classes and link function pointers
	e2k.cleanFn = clean.Clean
	vowel := newVowelConverter()
	e2k.vowelFn = vowel.ConvertVowel
	consonant := newConsonantConverter()
	e2k.consonantFn = consonant.ConvertConsonant
	epenthetic := newEpentheticVowelHandler()
	e2k.epentheticFn = epenthetic.AddEpentheticVowel
	morae := newMoraeCreator()
	e2k.moraeFn = morae.CreateMorae
	kana := newMoraeKanaConverter()
	e2k.kanaFn = kana.ConvertMorae
	r2k := NewRomajiToKana(strictF)
	e2k.recoverFn = r2k.Convert

	// Load cmu_ipa english phoneme pronounce dictionary
	e2k.loadDB()

	// Return instance
	return &e2k
}

// LoadDBFromFile loads the JSON file containing the database.
func (e2k *EngToKana) loadDB() error {
	// Unmarshal the JSON data into a map[string][]string
	if err := json.Unmarshal([]byte(dbFile), &e2k.db); err != nil {
		return err
	}

	return nil
}

// TranscriptWord converts an English word to Katakana
func (e2k *EngToKana) TranscriptWord(word string) string {
	phs, ok := e2k.db[word]
	if !ok {
		// return "E_DIC"
		// If no match found, try to recover by using Romaji2Kana
		return e2k.recoverFn(word)
	}

	var result []string
	for _, ph := range phs {
		ph1 := e2k.vowelFn(word, ph)
		ph2 := e2k.consonantFn(word, ph1)
		ph3 := e2k.epentheticFn(ph2)
		morae := e2k.moraeFn(ph3)
		kana := e2k.kanaFn(morae)
		result = append(result, kana)
	}
	return result[0]
}

// TranscriptSentence converts an English sentence to Katakana
func (e2k *EngToKana) TranscriptSentence(line string) string {
	// Clean string with call back to process clean sentence fragment
	return e2k.cleanFn(line, e2k.transcriptCleanSentenceFragment)
}

// processes a line of text containing English words into Kana
func (e2k *EngToKana) transcriptCleanSentenceFragment(line string) string {
	var result strings.Builder
	words := strings.Fields(line)

	// Iterate over the words and convert each to Katakana
	for _, word := range words {
		katakanaWords := e2k.TranscriptWord(word)
		// Placeholder logic to join Katakana words
		result.WriteString(katakanaWords)

		// TODO: Recover func call in case of E_DIC
		// TODO: Add and try romaji_to_kana.go
	}
	return result.String()
}