package kanatrans

type AllToKana struct {
	e2k		*EngToKana
	k2k		*KanjiToKana
	h2k		*HiraganaToKatakana
	ks		*KanjiSplitter
}

func NewAllToKana(strictPunct ...bool) *AllToKana {
	// Set optional strict bool
	var strict bool
	if len(strictPunct) > 0 {
		strict = strictPunct[0]
	}

	// Instantiate class
	a2k := AllToKana{}

	// Create an instance of EngToKana
	a2k.e2k = NewEngToKana(true)
	// Create an instance of KanjiToKana
	a2k.k2k = NewKanjiToKana()
	// Create an instance of GanaToKana
	a2k.h2k = NewHiraganaToKatakana()
	// Determine punctuation converter to use
	var puncFP func(string) string
	if strict {
		puncFP = convertToJapanesePunctuationRestricted
	} else {
		puncFP = convertToJapanesePunctuation
	}
	// Create an instance of KanjiSplitter
	a2k.ks = NewKanjiSplitter(
		a2k.k2k.Convert,					// Kanji callback
		a2k.h2k.Convert,					// Gana & Kana callback
		a2k.e2k.TranscriptSentence,			// English callback
		puncFP,								// Punctuation callback
	)

	// Return instance
	return &a2k
}

// Function to perform conversion
func (a2k *AllToKana) Convert(s string) string {
	return a2k.ks.SeparateAndProcess(s)
}