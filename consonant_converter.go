package kanatrans

import (
	"strings"
)

// ConsonantConverter handles consonant conversions.
type ConsonantConverter struct {
	vowels      string
	consonants  string
}

// NewConsonantConverter creates a new instance of ConsonantConverter.
func NewConsonantConverter() *ConsonantConverter {
	return &ConsonantConverter{
		vowels:     "aeiou",
		consonants: "dgʤʒklmnptvʧŋɹʃðθ",
	}
}

// dRule handles d, dz -- dd, z.
func (cc *ConsonantConverter) dRule(word, ph string, pIdx int) string {
	if pIdx+1 < len(ph) && ph[pIdx+1] == 'z' {
		return "z"
	} else if pIdx >= 1 && strings.ContainsAny(string(ph[pIdx-1]), cc.vowels) && (len(ph) <= 2 || !strings.ContainsAny(string(ph[pIdx-2]), cc.vowels)) {
		return "dd"
	}
	return "d"
}

// gkptRule handles k, g, p, t -- kk, gg, pp, tt.
func (cc *ConsonantConverter) gkptRule(word, ph string, pIdx int) string {
	if len(ph) == 2 && pIdx == 1 {
		return strings.Repeat(string(ph[pIdx]), 2)
	} else if pIdx-1 >= 0 && pIdx+2 < len(ph) && ph[pIdx-1] == 'a' && ph[pIdx+1] == 'u' && ph[pIdx+2] == 'l' {
		return strings.Repeat(string(ph[pIdx]), 2)
	} else if pIdx >= 3 && strings.ContainsAny(string(ph[pIdx-1]), cc.vowels) && strings.ContainsAny(string(ph[pIdx-2]), cc.vowels) && strings.ContainsAny(string(ph[pIdx-3]), cc.vowels) {
		return strings.Repeat(string(ph[pIdx]), 2)
	} else if pIdx >= 2 && strings.ContainsAny(string(ph[pIdx-1]), cc.vowels) && !strings.ContainsAny(string(ph[pIdx-2]), cc.vowels) && (pIdx == len(ph)-1 || !strings.ContainsAny(string(ph[pIdx+1]), cc.vowels)) {
		return strings.Repeat(string(ph[pIdx]), 2)
	}
	return string(ph[pIdx])
}

// dgRule handles ʤ -- j, jj.
func (cc *ConsonantConverter) dgRule(word, ph string, pIdx int) string {
	if pIdx >= 1 && strings.ContainsAny(string(ph[pIdx-1]), cc.vowels) {
		if len(ph) <= 2 || (pIdx+1 < len(ph) && strings.ContainsAny(string(ph[pIdx+1]), cc.vowels)) {
			return "j"
		} else if pIdx-2 >= 0 && !strings.ContainsAny(string(ph[pIdx-2]), cc.vowels) {
			return "jj"
		}
	}
	return "j"
}

// gShortRule handles ʒ.
func (cc *ConsonantConverter) gShortRule(word, ph string, pIdx int) string {
	return "j"
}

// lRule handles l -- r.
func (cc *ConsonantConverter) lRule(word, ph string, pIdx int) string {
	return "r"
}

// mnRule handles m, n not followed by vowel -- N and y.
func (cc *ConsonantConverter) mnRule(word, ph string, pIdx int) string {
	if pIdx > 0 && pIdx+1 < len(ph) && !strings.ContainsAny(string(ph[pIdx+1]), cc.vowels) && string(ph[pIdx+1]) != "y" {
		return "N"
	} else if pIdx == len(ph)-1 && string(ph[pIdx]) == "n" {
		return "N"
	}
	return string(ph[pIdx])
}

// vRule handles v -- b.
func (cc *ConsonantConverter) vRule(word, ph string, pIdx int) string {
	return "b"
}

// tshRule handles ʧ -- ch or cch.
func (cc *ConsonantConverter) tshRule(word, ph string, pIdx int) string {
	if pIdx >= 1 && strings.ContainsAny(string(ph[pIdx-1]), cc.vowels) && (len(ph) <= 2 || !strings.ContainsAny(string(ph[pIdx-2]), cc.vowels)) {
		return "cch"
	}
	return "ch"
}

// ngRule handles ŋ -- N or Ng.
// TODO: darling --> daariN
func (cc *ConsonantConverter) ngRule(word, ph string, pIdx int) string {
	if pIdx+1 < len(ph) && ph[pIdx+1] == 'g' {
		return "N"
	} else if strings.Contains(word, "ng") {
		return "Ng"
	}
	return "N"
}

// rRule handles ɹ.
func (cc *ConsonantConverter) rRule(word, ph string, pIdx int) string {
	return "r"
}

// shRule handles ʃ.
func (cc *ConsonantConverter) shRule(word, ph string, pIdx int) string {
	if pIdx >= 1 && strings.ContainsAny(string(ph[pIdx-1]), cc.vowels) && (len(ph) <= 2 || !strings.ContainsAny(string(ph[pIdx-2]), cc.vowels)) {
		return "ssh"
	}
	return "sh"
}

// thHakuonRule handles ð -- z.
func (cc *ConsonantConverter) thHakuonRule(word, ph string, pIdx int) string {
	return "z"
}

// thClearRule handles θ.
func (cc *ConsonantConverter) thClearRule(word, ph string, iIdx int) string {
	return "s"
}

// ConvertConsonant converts consonants.
func (cc *ConsonantConverter) ConvertConsonant(word, ph string) string {
	consonantMap := map[string]func(string, string, int) string{
		"d":  cc.dRule,
		"g":  cc.gkptRule,
		"ʤ": cc.dgRule,
		"ʒ": cc.gShortRule,
		"k":  cc.gkptRule,
		"l":  cc.lRule,
		"m":  cc.mnRule,
		"n":  cc.mnRule,
		"p":  cc.gkptRule,
		"t":  cc.gkptRule,
		"v":  cc.vRule,
		"ʧ": cc.tshRule,
		"ŋ":  cc.ngRule,
		"ɹ":  cc.rRule,
		"ʃ":  cc.shRule,
		"ð":  cc.thHakuonRule,
		"θ":  cc.thClearRule,
	}

	var result strings.Builder
	pIdx := 0

	phS, wordS := NewRuneString(ph), NewRuneString(word)

	for pIdx < phS.Len() {
		// adds a vowel char as is
		for pIdx < phS.Len() && !strings.Contains(cc.consonants,phS.CharAt(pIdx)) {
			result.WriteString(phS.CharAt(pIdx))
			pIdx++
		}

		// converts a consonant
		if pIdx < phS.Len() && strings.Contains(cc.consonants,phS.CharAt(pIdx)) {
			result.WriteString(consonantMap[phS.CharAt(pIdx)](wordS.String(),phS.String(),pIdx))
			pIdx++
		}
	}
	
	return result.String()
}