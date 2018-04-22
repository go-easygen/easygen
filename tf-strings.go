package easygen

import (
	"regexp"
	"strings"
)

var (
	re  *regexp.Regexp
	rec *regexp.Regexp
)

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
// CLI Function Definitions

// TFStringsInit does initialization for strings related template functions
func TFStringsInit() {
	re = regexp.MustCompile(`(?i)` + Opts.StrFrom)
	// case sensitive string replace
	rec = regexp.MustCompile(Opts.StrFrom)

}

//==========================================================================
// command-line replacing option template function

func replace(replStr string) string {
	return re.ReplaceAllString(replStr, Opts.StrTo)
}

// replacec does a case sensitive string replace
func replacec(replStr string) string {
	return rec.ReplaceAllString(replStr, Opts.StrTo)
}

// quote4shell -- quote file name for shell.
// So "%bob's file" will be quoted as '%bob'\''s file'
func quote4shell(s string) string {
	return "'" + strings.Join(strings.Split(s, "'"), `'\''`) + "'"
}
