package main

import (
	"strings"
)

// EpentheticVowelHandler handles epenthetic vowels.
type EpentheticVowelHandler struct {
	vowels  string
	consym  string
}

// NewEpentheticVowelHandler creates a new instance of EpentheticVowelHandler.
func NewEpentheticVowelHandler() *EpentheticVowelHandler {
	return &EpentheticVowelHandler{
		vowels: "aeiou",
		consym: "bdfghjkmprstz",
	}
}

// dtRule handles d, t.
func (evh *EpentheticVowelHandler) dtRule(ph string, pIdx int) string {
	// d, t
	if pIdx == len(ph)-1 || (pIdx+1 < len(ph) && ph[pIdx] != ph[pIdx+1] && !strings.ContainsAny(string(ph[pIdx+1]), evh.vowels)) {
		return string(ph[pIdx]) + "o"
	}
	return string(ph[pIdx])
}

// bfprzRule handles b, f, p, r, z.
func (evh *EpentheticVowelHandler) bfprzRule(ph string, pIdx int) string {
	// b, f, p, r, z
	if pIdx == len(ph)-1 || (pIdx+1 < len(ph) && ph[pIdx] != ph[pIdx+1] && string(ph[pIdx+1]) != "y" && !strings.ContainsAny(string(ph[pIdx+1]), evh.vowels)) {
		return string(ph[pIdx]) + "u"
	}
	return string(ph[pIdx])
}

// kgmRule handles k, g, m.
func (evh *EpentheticVowelHandler) kgmRule(ph string, pIdx int) string {
	// k, g, m
	if pIdx == len(ph)-1 || (pIdx+1 < len(ph) && ph[pIdx] != ph[pIdx+1] && string(ph[pIdx+1]) != "y" && string(ph[pIdx+1]) != "w" && !strings.ContainsAny(string(ph[pIdx+1]), evh.vowels)) {
		return string(ph[pIdx]) + "u"
	}
	return string(ph[pIdx])
}

// hRule handles cch, ssh.
func (evh *EpentheticVowelHandler) hRule(ph string, pIdx int) string {
	// cch, ssh
	if pIdx >= 1 && ph[pIdx-1] == 'c' && (pIdx == len(ph)-1 || !strings.ContainsAny(string(ph[pIdx+1]), evh.vowels)) {
		return "hi"
	} else if pIdx >= 1 && ph[pIdx-1] == 's' && (pIdx == len(ph)-1 || !strings.ContainsAny(string(ph[pIdx+1]), evh.vowels)) {
		return "hu"
	} else if pIdx+1 < len(ph) && ph[pIdx+1] == 'w' {
		return "ho"
	}
	return "h"
}

// sRule handles s.
func (evh *EpentheticVowelHandler) sRule(ph string, pIdx int) string {
	if pIdx == len(ph)-1 || (pIdx+1 < len(ph) && ph[pIdx] != ph[pIdx+1] && string(ph[pIdx+1]) != "h" && !strings.ContainsAny(string(ph[pIdx+1]), evh.vowels)) {
		return "su"
	}
	return "s"
}

// jRule handles j.
func (evh *EpentheticVowelHandler) jRule(ph string, pIdx int) string {
	if pIdx == len(ph)-1 || (pIdx+1 < len(ph) && !strings.ContainsAny(string(ph[pIdx+1]), evh.vowels)) {
		if string(ph[pIdx+1]) == "j" || string(ph[pIdx+1]) == "y" {
			return "j"
		}
		return "ji"
	}
	return "j"
}

// AddEpentheticVowel adds an epenthetic vowel.
func (evh *EpentheticVowelHandler) AddEpentheticVowel(ph string) string {
	epentheticMap := map[byte]func(string, int) string{
		'b': evh.bfprzRule,
		'd': evh.dtRule,
		'f': evh.bfprzRule,
		'g': evh.kgmRule,
		'h': evh.hRule,
		'j': evh.jRule,
		'k': evh.kgmRule,
		'm': evh.kgmRule,
		'p': evh.bfprzRule,
		'r': evh.bfprzRule,
		's': evh.sRule,
		't': evh.dtRule,
		'z': evh.bfprzRule,
	}
	var result strings.Builder
	pIdx := 0
	for pIdx < len(ph) {
		for pIdx < len(ph) && strings.IndexByte(evh.consym, ph[pIdx]) == -1 {
			// skip if the same vowel continues more than two
			if pIdx >= 2 && ph[pIdx] == ph[pIdx-1] && ph[pIdx] == ph[pIdx-2] {
				// no-op
			} else {
				result.WriteByte(ph[pIdx])
			}
			pIdx++
		}
		if pIdx < len(ph) && strings.IndexByte(evh.consym, ph[pIdx]) != -1 {
			result.WriteString(epentheticMap[ph[pIdx]](ph, pIdx))
			pIdx++
		}
	}
	return result.String()
}