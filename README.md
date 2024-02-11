# English2KanaTransliteration
Convert English phrases into phonetic Japanese kana approximations; also known as Englishru. Does not translate English into Japanese, but translates English words into their approximate pronounciations in Japanese. *No dependencies!*

Based on the English to Katakana transcription code written in Python by [Yoko Harada (@yokolet)](https://github.com/yokolet/transcript) Please see that repo for details on the phonetic conversion.

English to phoneme conversion based on [CMUDict](https://people.umass.edu/nconstan/CMU-IPA/). Kanji to Katakana convertion based on [KANJIDIC2](http://nihongo.monash.edu/kanjidic2/index.html). Thanks to [JMDict](https://pkg.go.dev/github.com/foosoft/jmdict) and [kana](https://github.com/gojp/kana). Please refer to those licenses for non-free implementations.

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
	// Create an instance of AllToKana
	allToKana := NewAllToKana()

	// Listen to stdin indefinitely
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break // Exit loop on error
		}

		// Call convertString function with the accumulated line
		result := allToKana.Convert(line)

		// Output the result
		fmt.Print(result+"\n")
	}
}
```

Sample Output:
```
❯ go run .
Hello there.
ヘロー　ゼアー。
With this program, you can make Japanese text to speech speak in English!
ウィズ　ジス　プローラ、　ユー　キャン　メイク　ジャーンイーズ　テックスト　ツー　スピーチ　スピーク　イン　イングシュ！
Hello! こんにちは~ ヘロー, 松原。
ヘロー！　コンニチハ〜　ヘロー、　ショウゲン。
```

### Using individual modules

#### All2Katakana
```go
```

#### Eng2Katakana
```go
// Create an instance of EngToKana
engToKana := NewEngToKana()
// Usage
kana := engToKana.TranscriptSentence("Hello World!")
```

#### Kanji2Katakana
```go
// Create an instance of KanjiToKana
kanjiToKana := NewKanjiToKana()
// Usage
kana := kanjiToKana.Convert("初音ミク")
```
This needs some work, it just takes the most common pronouciation of each Kanji instead of the correct one for the context. Pull requests are welcome!

#### Hiragana2Katakana
```go
// Create an instance of HiraganaToKana
hiraganaToKana := NewHiraganaToKatakana()
// Usage
kana := hiraganaToKana.Convert("こんにちは")
```

#### ConvertPunctuation
```go
// Usage
japanesePunctuation := convertToJapanesePunctuation("Hello, World!")
```

### Note for using with Japanese-only text-to-speech (TTS)
This module is intended to allow TTS which only support Japanese to speak english (such as AquesTalk, Softalk, etc). These TTS usually have some limitations in what punctuation may be present in the input; with only commas and stops being interpreted as a pause and all other punctuation causing an error.

To use this module for such TTS input, you may enable *strict input cleaning mode* (only Japanese comma and stop on output) by passing a bool in the initialiser for `EngToKana` and `AllToKana` classes:
```go
// Create an instance of AllToKana with strict punctuation output
allToKana := NewAllToKana(true)
```
```go
// Create an instance of EngToKana with strict punctuation output
engToKana := NewEngToKana(true)
```
You may also use the function `convertToJapanesePunctuationRestricted` instead of `convertToJapanesePunctuation`.

### Custom callbacks to proccess Kanji, Kana, English & Punctuation
Internally, the `AllToKana` proccess function uses a `KanjiSplitter` class to call `func(string) string` functions which handle Kanji, Kana, English and Punctuation respectively:
```go
// Create an instance of KanjiSplitter with proccesing callbacks
kanjiSplitter := NewKanjiSplitter(
	kanjiToKana.Convert,					// Kanji callback
	hiraganaToKana.Convert,					// Gana & Kana callback
	engToKana.TranscriptSentence,			// English callback
	convertToJapanesePunctuation,			// Punctuation callback
)
```
If required, you may use a `KanjiSplitter` with custom callback functions to provide different processing.