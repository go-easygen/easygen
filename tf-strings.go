package easygen

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	re  *regexp.Regexp
	rec *regexp.Regexp
)

////////////////////////////////////////////////////////////////////////////
// Normal String Function Definitions

// Taken/adapted from hugo tpl/strings/strings.go

// Substr extracts parts of a string, beginning at the character at the specified
// position, and returns the specified number of characters.
//
// It normally takes two parameters: start and length.
// It can also take one parameter: start, i.e. length is omitted, in which case
// the substring starting from start until the end of the string will be returned.
//
// To extract characters from the end of the string, use a negative start number.
//
// In addition, borrowing from the extended behavior described at http://php.net/substr,
// if length is given and is negative, then that many characters will be omitted from
// the end of string.
func Substr(a string, nums ...interface{}) (string, error) {
	var start, length int

	asRunes := []rune(a)

	switch len(nums) {
	case 0:
		return "", errors.New("too less arguments")
	case 1:
		start = nums[0].(int)
		length = len(asRunes)
	case 2:
		start = nums[0].(int)
		length = nums[1].(int)
	default:
		return "", errors.New("too many arguments")
	}

	if start < -len(asRunes) {
		start = 0
	}
	if start > len(asRunes) {
		return "", fmt.Errorf("start position out of bounds for %d-byte string", len(a))
	}

	var s, e int
	if start >= 0 && length >= 0 {
		s = start
		e = start + length
	} else if start < 0 && length >= 0 {
		s = len(asRunes) + start
		e = len(asRunes) + start + length + 1
	} else if start >= 0 && length < 0 {
		s = start
		e = len(asRunes) + length
	} else {
		s = len(asRunes) + start
		e = len(asRunes) + length
	}

	if s > e {
		return "", fmt.Errorf("calculated start position greater than end position: %d > %d", s, e)
	}
	if e > len(asRunes) {
		e = len(asRunes)
	}

	return string(asRunes[s:e]), nil
}

////////////////////////////////////////////////////////////////////////////
// Regexp Function Definitions

func regexpFindAllString(s string, regExp string, n int) []string {
	return regexp.MustCompile(regExp).FindAllString(s, n)
}

func regexpFindAllStringIndex(s string, regExp string, n int) [][]int {
	return regexp.MustCompile(regExp).FindAllStringIndex(s, n)
}

func regexpFindAllStringSubmatch(s string, regExp string, n int) [][]string {
	return regexp.MustCompile(regExp).FindAllStringSubmatch(s, n)
}

func regexpFindAllStringSubmatchIndex(s string, regExp string, n int) [][]int {
	return regexp.MustCompile(regExp).FindAllStringSubmatchIndex(s, n)
}

func regexpFindString(s string, regExp string) string {
	return regexp.MustCompile(regExp).FindString(s)
}

func regexpFindStringIndex(s string, regExp string) (loc []int) {
	return regexp.MustCompile(regExp).FindStringIndex(s)
}

func regexpFindStringSubmatch(s string, regExp string) []string {
	return regexp.MustCompile(regExp).FindStringSubmatch(s)
}

func regexpFindStringSubmatchIndex(s string, regExp string) []int {
	return regexp.MustCompile(regExp).FindStringSubmatchIndex(s)
}

func regexpMatchString(s string, regExp string) bool {
	return regexp.MustCompile(regExp).MatchString(s)
}

func regexpReplaceAllLiteralString(src, regExp string, repl string) string {
	return regexp.MustCompile(regExp).ReplaceAllLiteralString(src, repl)
}

func regexpReplaceAllString(src, regExp string, repl string) string {
	return regexp.MustCompile(regExp).ReplaceAllString(src, repl)
}

func regexpReplaceAllStringFunc(src string, regExp string, repl func(string) string) string {
	return regexp.MustCompile(regExp).ReplaceAllStringFunc(src, repl)
}

func regexpSplit(s string, regExp string, n int) []string {
	return regexp.MustCompile(regExp).Split(s, n)
}

////////////////////////////////////////////////////////////////////////////
// Misc

// coalesce function takes two or more string arguments and returns the first argument that is not empty.
// The result is empty only if all the arguments are empty.
func coalesce(s ...string) string {
	for _, str := range s {
		if len(str) != 0 && str != "<no value>" {
			return str
		}
	}
	return ""
}

// quote4shell -- quote file name for shell.
// So "%bob's file" will be quoted as '%bob'\''s file'
func quote4shell(s string) string {
	return "'" + strings.Join(strings.Split(s, "'"), `'\''`) + "'"
}
