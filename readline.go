// v0.1.0
// Author: Eric DIEHL
// (C) Oct 2024

package toolbox

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

var (
	// ErrParsingFailed is an error occurring when parsing failed  because the expected field is not present or the
	// other fields have not the proper format
	ErrParsingFailed = errors.New("parsing failed ")
)

// ReadTheLine extracts from the reader a line ending with '\n'.
// If strip is true, then it strips away the carriage return characters.
func ReadTheLine(rd io.Reader, strip bool) (string, error) {
	rd1 := bufio.NewReader(rd)
	s, err := rd1.ReadString('\n')
	if strip {
		s = strings.Trim(s, "\r\n") // It is mandatory to strip also \r added by Windows
	}
	return s, err
}

// ReadFieldsInLine reads a line ending with '\n' from the buffer
// and returns the different fields of the scanned line.
func ReadFieldsInLine(rd io.Reader) ([]string, error) {
	s, err := ReadTheLine(rd, true)
	if err != nil {
		return nil, err
	}
	s = strings.Trim(s, "\r\n") // It is mandatory to strip also \r added by Windows
	sParsed := strings.Fields(s)
	return sParsed, nil
}

// ParseFieldsInLine reads a line from r, checks whether the first field is equal to expected.  If it is the case, then
// it returns the other words as a slice of strings or nil if there was no other fields than the expected one.  It
// should be noted the test assumes that expected is a word, i.e., no spaces.  It removes all leading and trailing spaces
//
//	The error may be ErrParsingFailed if there was a first field but not the one expected.  It returns typical scanner
//	errors such as EOF.
func ParseFieldsInLine(rd io.Reader, expected string) ([]string, error) {
	expected = strings.Trim(expected, " ") // removes any trailing and leading spaces
	sParsed, err := ReadFieldsInLine(rd)
	if err != nil {
		return nil, err
	}
	if len(sParsed) == 0 {
		return nil, ErrParsingFailed
	}
	if sParsed[0] != expected {
		return nil, ErrParsingFailed
	}
	return sParsed[1:], nil
}
