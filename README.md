# English2KanaTransliteration
Convert English phrases into phonetic Japanese kana approximations; also known as Englishru. Does not translate English into Japanese, but translates English words into their approximate pronounciations in Japanese.

Based on the English to Katakana transcription code written in Python by [Yoko Harada (@yokolet)](https://github.com/yokolet/transcript) Please see that repo for details on the phonetic conversion.

English to phoneme conversion based on [CMUDict](https://people.umass.edu/nconstan/CMU-IPA/). Kanji to Katakana convertion based on [KANJIDIC2](http://nihongo.monash.edu/kanjidic2/index.html). Thanks to [JMDict](https://pkg.go.dev/github.com/foosoft/jmdict).Please refer to those licenses for non-free implementations.

It is a port in Golang with some additional functions:
- Filtering functions to split, parse, and rejoin sentences which contain punctuation or improper contractions.

Work in progress:
- Also **accepts Japanese input**; converts any Kanji characters into their most common Hiragana pronounciation, converts Hiragana into Katakana, leaves Katakana as is.


## Usage Example
Below is an example go file to test this module. It reads input from `stdin`, converts the English sentences into their Japanese transliteration and prints them to `stdout`.

```go
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
```

Sample Output:
```
❯ go run .
Hello there
ヘローゼアー
With this program, you can make Japanese text to speech speak in English
ウィズジスプローラ、ユーキャンメイクジャーンイーズテックストツースピーチスピークインイングシュ
Hello~ こんにちは、 ヘロー, 你好! World! 山川。木田
ヘロー〜　コンニチハ、　ヘロー、　ジコウ！　ワールド！　サンセン。ボクデン
```


### Note for using with Japanese-only text-to-speech (TTS)
This module is intended to allow TTS which only support Japanese to speak english (such as AquesTalk, Softalk, etc). These TTS usually have some limitations in what punctuation may be present in the input; with only commas and stops being interpreted as a pause and all other punctuation causing an error.

To use this module for such TTS input, you may enable *strict input cleaning mode* (only Japanese comma and stop on output) by passing a bool in the initialiser:
```go
// Create an instance of EngToKana
engToKana := NewEngToKana(true)
```
You may also use the function `convertToJapanesePunctuationRestricted` instead of `convertToJapanesePunctuation`:
```go
kanjiSplitter := NewKanjiSplitter(
	kanjiToKana.Convert,					// Kanji callback
	hiraganaToKana.Convert,					// Gana & Kana callback
	engToKana.TranscriptSentence,			// English callback
	convertToJapanesePunctuationRestricted,	// Punctuation callback
)
```