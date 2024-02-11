package kanatrans

import (
	"strings"
	"unicode"
)

// KanjiSplitter is a class to split a string into segments of Roman, Japanese, and Han text
type KanjiSplitter struct{
	kanjiCallback		func(string) string
	kanaCallback		func(string) string
	romanCallback		func(string) string
	punctCallback		func(string) string
}

// NewKanjiSplitter creates a new instance of KanjiSplitter
func NewKanjiSplitter(kanjiCallback, kanaCallback, romanCallback, punctCallback func(string) string) *KanjiSplitter {
	ks := KanjiSplitter{
		kanjiCallback: kanjiCallback,
		kanaCallback: kanaCallback,
		romanCallback: romanCallback,
		punctCallback: punctCallback,
	}
	return &ks
}

// SeparateAndProcess separates the input string into segments of Roman, Japanese, and Han text,
// and processes each segment accordingly
func (ks *KanjiSplitter) SeparateAndProcess(input string) string {
	var result strings.Builder

	// Separate the input string into segments of Roman, Japanese, and Han text
	segments := ks.separateRomanJapaneseAndHan(input)

	// Iterate over the segments
	for _, segment := range segments {
		// Process each segment differently based on its content
		if ks.isJapanese(segment) {
			// Call the ProcessJapanese function for Japanese segments
			result.WriteString(ks.kanaCallback(segment))
		} else if ks.isHan(segment) {
			// Call the ProcessHan function for Han (Chinese) segments
			result.WriteString(ks.kanjiCallback(segment))
		} else if ks.isPunctuation(segment) {
			// Call the function to handle punctuation
			result.WriteString(ks.punctCallback(segment))
		} else {
			// Call the ProcessRoman function for Roman segments
			result.WriteString(ks.romanCallback(segment))
		}
	}

	return result.String()
}

// Function to separate the input string into segments of Roman, Japanese, and Han text
func (ks *KanjiSplitter) separateRomanJapaneseAndHan(input string) []string {
	var segments []string
	var currentSegment strings.Builder
	lastType := ks.segmentType(rune(input[0]))

	for _, char := range input {
		currentType := ks.segmentType(char)

		if currentType != lastType {
			// Start a new segment
			if currentSegment.String() != "" {
				segments = append(segments, currentSegment.String())
			}
			currentSegment.Reset()
			currentSegment.WriteString(string(char))
		} else {
			// Continue the current segment
			currentSegment.WriteString(string(char))
		}

		lastType = currentType
	}

	// Append the last segment
	if currentSegment.String() != "" {
		segments = append(segments, currentSegment.String())
	}

	return segments
}

// Function to determine the type of segment (Roman, Japanese, or Han)
func (ks *KanjiSplitter) segmentType(char rune) int {
	switch {
	case char >= 'あ' && char <= 'ん':
		return 1 // Japanese
	case char >= '一' && char <= '龯':
		return 2 // Han (Chinese)
	case ks.isPunctuation(string(char)):
		return 3 // Punctuation
	default:
		return 0 // Roman or other
	}
}

// Function to check if a string contains Japanese text (hiragana or katakana)
func (ks *KanjiSplitter) isJapanese(s string) bool {
    for _, char := range s {
        if unicode.Is(unicode.Hiragana, char) || unicode.Is(unicode.Katakana, char) {
            return true
        }
    }
    return false
}

// Function to check if a string contains Han (Chinese) text
func (ks *KanjiSplitter) isHan(s string) bool {
    for _, char := range s {
        if unicode.Is(unicode.Han, char) {
            return true
        }
    }
    return false
}

// Function to check if a string contains punctuation characters or their Japanese equivalents
func (ks *KanjiSplitter) isPunctuation(s string) bool {
	punctuationChars := " ?!;:-~,.'？！；：〜ー、。　"

	// Iterate over each character in the string
	for _, char := range s {
		// Check if the character is punctuation or its Japanese equivalent
		if strings.ContainsRune(punctuationChars, char) {
			return true
		}
	}
	return false
}