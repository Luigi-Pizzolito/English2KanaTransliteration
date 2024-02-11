package kanatrans

import (
	"regexp"
	"strings"
)

// cleanEnglish struct
type cleanEnglish struct {
	strictMode			bool
	apostropheMapping	map[string]string
}

// stingSegment represents a text segment with a type
type stringSegment struct {
	Text string
	Type string
}

// Clean Function to perform cleaning, and then call the corresponding EngToKana string processor
func (ce *cleanEnglish) Clean(line string, processCallback func(string)string) string {
	if ce.strictMode {
		return ce.strictPunctClean(line, processCallback)
	}
	return ce.simpleClean(line, processCallback)
}

// simpleClean Function to perform simple cleaning
func (ce *cleanEnglish) simpleClean(line string, processCallback func(string)string) string {
	// Initial input clean
	inputString := ce.removeNonAlphaKeepSomePuncMore(line)
	// Perform fragment splitting
	// Split the input string by '.' and ',' and other characters
	pattern := regexp.MustCompile(`([?!;:\-~,.])`)
	fields := pattern.Split(inputString, -1)

	segments := []stringSegment{}
	ind := 0
	for _, segment := range fields {
		ind += len(segment)+1
		segment = strings.TrimSpace(segment)
		if ind < len(inputString)-1 {
			indc := ind
			for inputString[indc] == ' ' {
				indc--
			}
			segment += string(inputString[indc])
		}
		
		if segment != "" {
			if strings.HasSuffix(segment, "?") || strings.HasSuffix(segment, "!") || strings.HasSuffix(segment, ";") || strings.HasSuffix(segment, ":") || strings.HasSuffix(segment, "-") || strings.HasSuffix(segment, "~") || strings.HasSuffix(segment, ",") {
				segments = append(segments, stringSegment{Text: segment[:len(segment)-1], Type: "text"})
				segments = append(segments, stringSegment{Type: ConvertToJapanesePunctuation(string(ce.getLastRune(segment)))})
			} else if strings.HasSuffix(segment, ".") {
				segments = append(segments, stringSegment{Text: segment[:len(segment)-1], Type: "text"})
				segments = append(segments, stringSegment{Type: "。"})
			} else {
				segments = append(segments, stringSegment{Text: segment, Type: "text"})
			}
		}
	}
	// Process segments
	return ce.processSegments(segments, processCallback)
}

func (ce *cleanEnglish) getLastRune(str string) rune {
	runes := []rune(str)
	if len(runes) == 0 {
		// Handle empty string case
		return 0
	}
	return runes[len(runes)-1]
}

// Function to perform strict punctuation cleaning; commas and stops only
func (ce *cleanEnglish) strictPunctClean(line string, processCallback func(string)string) string {
	// Initial input clean
	inputString := ce.removeNonAlphaKeepSomePuncMore(line)
	// Perform fragment splitting
	// Split the input string by '.' and ',' and other characters
	pattern := regexp.MustCompile(`([?!;:\-~,.])`)
	fields := pattern.Split(inputString, -1)

	segments := []stringSegment{}
	ind := 0
	for _, segment := range fields {
		ind += len(segment)+1
		segment = strings.TrimSpace(segment)
		if ind < len(inputString)-1 {
			indc := ind
			for inputString[indc] == ' ' {
				indc--
			}
			segment += string(inputString[indc])
		}
		
		if segment != "" {
			if strings.HasSuffix(segment, "?") || strings.HasSuffix(segment, "!") || strings.HasSuffix(segment, ";") || strings.HasSuffix(segment, ":") || strings.HasSuffix(segment, "-") || strings.HasSuffix(segment, "~") || strings.HasSuffix(segment, ",") {
				segments = append(segments, stringSegment{Text: segment[:len(segment)-1], Type: "text"})
				segments = append(segments, stringSegment{Type: "、"})
			} else if strings.HasSuffix(segment, ".") {
				segments = append(segments, stringSegment{Text: segment[:len(segment)-1], Type: "text"})
				segments = append(segments, stringSegment{Type: "。"})
			} else {
				segments = append(segments, stringSegment{Text: segment, Type: "text"})
			}
		}
	}
	// Process segments
	return ce.processSegments(segments, processCallback)
}

func (ce *cleanEnglish) processSegments(segments []stringSegment, processCallback func(string)string) string {
	// Process the struct slices
	var outputString strings.Builder
	for _, structSlice := range segments {
		if structSlice.Type == "text" {
			// if segment type is text, perform the conversion
			text := structSlice.Text
			text = ce.replaceApostrophes(text)
			text = strings.ToLower(text)
			text = processCallback(text)
			outputString.WriteString(text)
		} else {
			// if segment type is punctuation, just print the punct
			outputString.WriteString(structSlice.Type)
		}
	}

	return outputString.String()
}

// RemoveNonAlpha removes non-alphabetic characters from a string except for apostrophes
func (ce *cleanEnglish) removeNonAlphaKeepSomePunc(line string) string {
	reg := regexp.MustCompile(`[^a-zA-Z' ]`)
	return reg.ReplaceAllString(line, "")
}

// RemoveNonAlpha removes non-alphabetic characters from a string except for apostrophes and some punctuation
func (ce *cleanEnglish) removeNonAlphaKeepSomePuncMore(line string) string {
	reg := regexp.MustCompile(`[^a-zA-Z?!;:\-~,.' ]`)
	return reg.ReplaceAllString(line, "")
}

// Function to replace words with apostrophes
func (ce *cleanEnglish) replaceApostrophes(line string) string {
	// Iterate over the apostrophe mapping and replace words
	for word, correctedWord := range ce.apostropheMapping {
		if strings.Contains(line, word) {
			line = strings.ReplaceAll(line, word, correctedWord)
		}
	}
	return line
}

// Initialiser
func newEnglishCleaner(strictOpt ...bool) *cleanEnglish {
	var optBool bool
	if len(strictOpt) > 0 {
		optBool = strictOpt[0]
	}
	return &cleanEnglish{
		// strict cleaning mode
		strictMode: optBool,
		// Dictionary to map words with apostrophes to their proper forms
		apostropheMapping: map[string]string{
			"aint": "ain't",
			"arent": "aren't",
			"cant": "can't",
			"couldnt": "couldn't",
			"couldve": "could've",
			"didnt": "didn't",
			"doesnt": "doesn't",
			"dont": "don't",
			"hadnt": "hadn't",
			"hasnt": "hasn't",
			"havent": "haven't",
			"he'd": "he'd",
			"he's": "he's",
			"hed": "he'd",
			"heres": "here's",
			"hes": "he's",
			"I'd": "I'd",
			"id": "I'd",
			"ill": "I'll",
			"im": "I'm",
			"isnt": "isn't",
			"it'd": "it'd",
			"it's": "it's",
			"its": "it's",
			"ive": "I've",
			"lets": "let's",
			"mightnt": "mightn't",
			"mustnt": "mustn't",
			"neednt": "needn't",
			"shant": "shan't",
			"she'd": "she'd",
			"she's": "she's",
			"shed": "she'd",
			"shes": "she's",
			"shouldnt": "shouldn't",
			"shouldve": "should've",
			"thats": "that's",
			"theres": "there's",
			"they'd": "they'd",
			"theyd": "they'd",
			"theyll": "they'll",
			"theyre": "they're",
			"theyve": "they've",
			"wasnt": "wasn't",
			"we'd": "we'd",
			"wed": "we'd",
			"were": "we're",
			"werent": "weren't",
			"weve": "we've",
			"whatll": "what'll",
			"whatre": "what're",
			"whats": "what's",
			"whatve": "what've",
			"wheres": "where's",
			"whod": "who'd",
			"whodve": "who'd've",
			"wholl": "who'll",
			"whore": "who're",
			"whos": "who's",
			"whove": "who've",
			"wont": "won't",
			"wouldnt": "wouldn't",
			"wouldve": "would've",
			"you'd": "you'd",
			"youd": "you'd",
			"youll": "you'll",
			"youre": "you're",
			"youve": "you've",
		},
	}
}