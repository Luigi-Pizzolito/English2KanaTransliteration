package kanatrans

import "fmt"

// RuneString represents a string with Unicode character indexing support
type RuneString struct {
	value string
}

// NewRuneString creates a new RuneString object from the given string
func NewRuneString(s string) *RuneString {
	return &RuneString{value: s}
}

// Len returns the number of Unicode characters in the string
func (s *RuneString) Len() int {
	return len([]rune(s.value))
}

// CharAt returns the Unicode character at the specified index
func (s *RuneString) CharAt(index int) string {
	runes := []rune(s.value)
	if index < 0 || index >= len(runes) {
		panic("index out of range")
		return ""
	}
	return string(runes[index])
}

// Substring returns a substring of the original string from index x to y
func (s *RuneString) Substring(x, y int) string {
	//! Stupid python, lower bound included, upper bound not included
	y--
	if x < 0 || y < 0 || x >= s.Len() || y >= s.Len() || y < x {
		panic(fmt.Sprintf("invalid indices: %d, %d in string %s",x,y,s.value))
		return ""
	}
	runes := []rune(s.value)
	return string(runes[x : y+1])
}

func (s *RuneString) String() string {
	return s.value
}