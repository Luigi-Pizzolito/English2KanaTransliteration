package main

import (
	"strings"
)

type VowelConverter struct {
	vowels  string
	vowsyms string
}

func NewVowelConverter() *VowelConverter {
	return &VowelConverter{
		vowels:  "aeiou",
		vowsyms: "aɑʌɚæeɛɪijɔoʊu",
	}
}

// ajRule applies the rule for "aj" or "ɑj" && returns the replacement string.
func (vc *VowelConverter) ajRule(word string, wIdx int) string {
	return "ai"
}

// arRule applies the rule for "ɑɹ" && returns the replacement string.
func (vc *VowelConverter) arRule(word string, wIdx int) string {
	return "aa"
}

// awRule applies the rule for "aw" or "ɑw" && returns the replacement string.
func (vc *VowelConverter) awRule(word string, wIdx int) string {
	return "au" // oo?
}

// aShortRule applies the rule for "ɑ" && returns the replacement string.
func (vc *VowelConverter) aShortRule(word string, wIdx int) string {
	if wIdx < len(word) && word[wIdx] == 'o' {
		return "o"
	}
	return "a"
}

// aLongRule applies the rule for "ɚ" && returns the replacement string.
func (vc *VowelConverter) aLongRule(word string, wIdx int) string {
	return "aa"
}

// aeRule applies the rule for "æ" && returns the replacement string.
func (vc *VowelConverter) aeRule(word string, wIdx int) string {
	if wIdx < len(word) && wIdx >= 1 && (word[wIdx-1] == 'c' || word[wIdx-1] == 'g') {
		return "ya"
	}
	return "a"
}

// hatRule applies the rule for "ʌ" && returns the replacement string.
func (vc *VowelConverter) hatRule(word string, wIdx int) string {
	if wIdx+1 < len(word) && word[wIdx] == 'o' && word[wIdx+1] == 'u' {
		return "a"
	} else if wIdx+1 < len(word) && word[wIdx] == 'i' && word[wIdx+1] == 'o' {
		return "o"
	} else if wIdx > 0 && wIdx < len(word) && word[wIdx] == 'e' && word[wIdx-1] == 'l' {
		if wIdx-2 >= 0 && (word[wIdx-2] == 'd' || word[wIdx-2] == 't') {
			return "o"
		} else if wIdx+1 < len(word) && word[wIdx+1] == 't' {
			return "e"
		} else {
			return "u"
		}
	} else if wIdx > 0 && wIdx < len(word) && word[wIdx] == 'u' && word[wIdx-1] == 'j' {
		return "ya"
	} else if wIdx < len(word) && word[wIdx] == 'o' && word != "mother" {
		return "o"
	} else if wIdx < len(word) && word[wIdx] == 'e' {
		return "e"
	} else if wIdx < len(word) && word[wIdx] == 'i' {
		return "i"
	} else {
		return "a"
	}
}

// ejRule applies the rule for "ej" && returns the replacement string.
func (vc *VowelConverter) ejRule(word string, wIdx int) string {
	return "ei"
}

// erRule applies the rule for "ɛɹ" && returns the replacement string.
func (vc *VowelConverter) erRule(word string, wIdx int) string {
	return "eaa"
}

// eRule applies the rule for "ɛ" && returns the replacement string.
func (vc *VowelConverter) eRule(word string, wIdx int) string {
	return "e"
}

// irRule applies the rule for "ɪɹ" && returns the replacement string.
func (vc *VowelConverter) irRule(word string, wIdx int) string {
	return "iaa"
}

// iirLongRule applies the rule for "iɹ" && returns the replacement string.
func (vc *VowelConverter) iirLongRule(word string, wIdx int) string {
	return "iaa"
}

// iLongRule applies the rule for "i" && returns the replacement string.
func (vc *VowelConverter) iLongRule(word string, wIdx int) string {
	return "ii"
}

// iShortRule applies the rule for "ɪ" && returns the replacement string.
func (vc *VowelConverter) iShortRule(word string, wIdx int) string {
	return "i"
}

// jaRule applies the rule for "jʌ" && returns the replacement string.
func (vc *VowelConverter) jaRule(word string, wIdx int) string {
	return "yaa"
}

// jaShortRule applies the rule for "jɑ" or "jæ" && returns the replacement string.
func (vc *VowelConverter) jaShortRule(word string, wIdx int) string {
	return "ya"
}

// jawRule applies the rule for "jaw" && returns the replacement string.
func (vc *VowelConverter) jawRule(word string, wIdx int) string {
	return "yoo"
}

// juRule applies the rule for "ju" && returns the replacement string.
func (vc *VowelConverter) juRule(word string, wIdx int) string {
	return "yuu"
}

// juShortRule applies the rule for "jɚ" or "jʊ" && returns the replacement string.
func (vc *VowelConverter) juShortRule(word string, wIdx int) string {
	return "yu"
}

// jiRule applies the rule for "ji" && returns the replacement string.
func (vc *VowelConverter) jiRule(word string, wIdx int) string {
	return "ii"
}

// jeRule applies the rule for "jɛ" or "jɪ" && returns the replacement string.
func (vc *VowelConverter) jeRule(word string, wIdx int) string {
	return "ie"
}

// jejRule applies the rule for "jej" && returns the replacement string.
func (vc *VowelConverter) jejRule(word string, wIdx int) string {
	return "yei"
}

// jowRule applies the rule for "jow" && returns the replacement string.
func (vc *VowelConverter) jowRule(word string, wIdx int) string {
	return "yoo"
}

// joRule applies the rule for "jɔ" && returns the replacement string.
func (vc *VowelConverter) joRule(word string, wIdx int) string {
	return "yo"
}

// ojRule applies the rule for "ɔj" or "ʌj" && returns the replacement string.
func (vc *VowelConverter) ojRule(word string, wIdx int) string {
	return "oi"
}

// orRule applies the rule for "ɔɹ" && returns the replacement string.
func (vc *VowelConverter) orRule(word string, wIdx int) string {
	return "oo"
}

// owRule applies the rule for "ow" && returns the replacement string.
func (vc *VowelConverter) owRule(word string, wIdx int) string {
	return "oo"
}

// oRule applies the rule for "ɔ" && returns the replacement string.
func (vc *VowelConverter) oRule(word string, wIdx int) string {
	if wIdx+1 < len(word) && word[wIdx] == 'a' && word[wIdx+1] == 'u' {
		return "oo"
	}
	return "o"
}

// jurRule applies the rule for "jʊɹ" && returns the replacement string.
func (vc *VowelConverter) jurRule(word string, wIdx int) string {
	return "yuaa"
}

// jsiRule applies the rule for "jsi" && returns the replacement string.
func (vc *VowelConverter) jsiRule(word string, wIdx int) string {
	return "ji"
}

// urRule applies the rule for "ʊɹ" && returns the replacement string.
func (vc *VowelConverter) urRule(word string, wIdx int) string {
	return "uaa"
}

// uShortRule applies the rule for "ʊ" && returns the replacement string.
func (vc *VowelConverter) uShortRule(word string, wIdx int) string {
	return "u"
}

// uLongRule applies the rule for "u" && returns the replacement string.
func (vc *VowelConverter) uLongRule(word string, wIdx int) string {
	return "uu"
}

// jRule applies the rule for the last "j" && returns the replacement string.
func (vc *VowelConverter) jRule(word string, wIdx int) string {
	return "ji"
}

// ConvertVowel converts vowels in a word according to phonetic rules && returns the converted string.
func (vc *VowelConverter) ConvertVowel(word, ph string) string {
	vowelMap := map[string]func(word string, wIdx int) string{
		"aj":  vc.ajRule,
		"ɑj": vc.ajRule,
		"ɑɹ": vc.arRule,
		"aw":  vc.awRule,
		"ɑw": vc.arRule,
		"ɑ":  vc.aShortRule,
		"ɚ":  vc.aLongRule,
		"æ":  vc.aeRule,
		"ʌ":  vc.hatRule,
		"ej":  vc.ejRule,
		"ɛɹ": vc.erRule,
		"ɛ":  vc.eRule,
		"ɪɹ": vc.irRule,
		"i":  vc.iLongRule,
		"iɹ": vc.irRule,
		"ɪ":  vc.iShortRule,
		"jʌ": vc.jaRule,
		"jɑ": vc.jaShortRule,
		"jæ": vc.jaShortRule,
		"jaw": vc.jawRule,
		"ju":  vc.juRule,
		"jɚ":  vc.juShortRule,
		"jʊ":  vc.juShortRule,
		"ji":  vc.jiRule,
		"jɪ":  vc.jeRule,
		"jɛ":  vc.jeRule,
		"jej": vc.jejRule,
		"jow": vc.jowRule,
		"jɔ":  vc.joRule,
		"jsi": vc.jsiRule,
		"ɔj":  vc.ojRule,
		"ʌj":  vc.ojRule,
		"ɔɹ":  vc.orRule,
		"ow":  vc.owRule,
		"ɔ":   vc.oRule,
		"jʊɹ": vc.jurRule,
		"ʊɹ":  vc.urRule,
		"ʊ":   vc.uShortRule,
		"u":   vc.uLongRule,
		"j":   vc.jRule,
	}

	var result strings.Builder
	wIdx, pIdx := 0, 0
	
	phS, wordS := NewRuneString(ph), NewRuneString(word)

	for pIdx < phS.Len() {
		// skips consnant in a word
		// in some cases, y is mapped to vowel sound
		for wIdx < wordS.Len() && wordS.CharAt(wIdx) != "y" && wordS.CharAt(wIdx) != "'" && !strings.Contains(vc.vowels, wordS.CharAt(wIdx)) {
			wIdx++
		}
		// adds consonant phonetics to the result, but does nothing for now
		for pIdx < phS.Len() && !strings.Contains(vc.vowsyms, phS.CharAt(pIdx)) {
			result.WriteString(phS.CharAt(pIdx))
			pIdx++
		}
		// convert vowel phonetics
		for pIdx < phS.Len() && strings.Contains(vc.vowsyms, phS.CharAt(pIdx)) {
			if wIdx+1 < wordS.Len() && wordS.CharAt(wIdx) == "u" && strings.Contains(vc.vowels, wordS.CharAt(wIdx+1)) {
				wIdx++
			}
			if pIdx+3 <= phS.Len() && stringInMapKey(phS.Substring(pIdx,pIdx+3),vowelMap) {
				result.WriteString(vowelMap[phS.Substring(pIdx,pIdx+3)](wordS.String(),wIdx))
				pIdx += 3
				wIdx++
			} else if pIdx+2 <= phS.Len() && stringInMapKey(phS.Substring(pIdx,pIdx+2),vowelMap) && (pIdx+2 == phS.Len() || phS.CharAt(pIdx+1) != "ɹ" || !strings.Contains(vc.vowsyms, phS.CharAt(pIdx+2)) ) {
				result.WriteString(vowelMap[phS.Substring(pIdx,pIdx+2)](wordS.String(),wIdx))
				pIdx += 2
				wIdx++
			} else if pIdx < phS.Len() && stringInMapKey(phS.CharAt(pIdx),vowelMap) {
				result.WriteString(vowelMap[phS.CharAt(pIdx)](wordS.String(),wIdx))
				pIdx++
				wIdx++
			} else if pIdx == phS.Len()-1 && phS.CharAt(pIdx) == "j"{
				result.WriteString(vowelMap[phS.CharAt(pIdx)](wordS.String(),wIdx))
				pIdx++
				wIdx++
			}
		}

		// skips rest of vowel chars
		for wIdx < wordS.Len() && strings.Contains(vc.vowels,wordS.CharAt(wIdx)) {
			wIdx++
		}
	}

	// adds rest of consonant symbols
	for pIdx < phS.Len() {
		result.WriteString(phS.CharAt(pIdx))
		pIdx++
	}

	return result.String()
}

func stringInMapKey(str string, m map[string]func(string, int)string) bool {
	_, ok := m[str]
	return ok
}