package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create an instance of EngToKana
	engToKana := NewEngToKana()
	// Create an instance of KanjiToKana
	kanjiToKana := NewKanjiToKana()
	// Create an instance of GanaToKana
	hiraganaToKana := NewHiraganaToKatakana()
	// Create an instance of KanjiSplitter
	kanjiSplitter := NewKanjiSplitter(
		kanjiToKana.Convert,			// Kanji callback
		hiraganaToKana.Convert,			// Gana & Kana callback
		engToKana.TranscriptSentence,	// English callback
		convertToJapanesePunctuation,	// Punctuation callback
	)

	// Listen to stdin indefinitely
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break // Exit loop on error
		}

		// Call convertString function with the accumulated line
		result := kanjiSplitter.SeparateAndProcess(line)

		// Output the result
		fmt.Print(result+"\n")
	}
}