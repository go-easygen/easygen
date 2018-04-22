package easygen_test

import (
	"bytes"
	"testing"

	"github.com/go-easygen/easygen"
)

type testItem struct {
	tmplStr, expected string
}

//var tmpl *template.Template

func testStringManipulation(t *testing.T, d []testItem) {
	tmpl = easygen.NewTemplate().Funcs(easygen.FuncDefs())
	for _, tc := range d {
		buf := bytes.NewBufferString("")
		easygen.Process0(tmpl, buf, tc.tmplStr, "test/strings")

		v := buf.String()
		//t.Log(v)
		if v != tc.expected {
			t.Errorf("'%s' expects '%s', got '%s'", tc.tmplStr, tc.expected, v)
		}
	}
}

// for standalone test, change package to `main` and the next func def to,
// func main() {
func TestStringManipulation(t *testing.T) {

	testData := []testItem{
		// == strings functions
		{
			`{{ stringsContains "seafood" "foo" }}`,
			"true",
		},
		{
			`{{ stringsContains "seafood" "bar" }}`,
			"false",
		},
		{
			`{{ stringsContainsAny "team" "i" }}`,
			"false",
		},
		{
			`{{ stringsContainsAny "failure" "u & i" }}`,
			"true",
		},
		{
			`{{ stringsCount "cheese" "e" }}`,
			"3",
		},
		{
			`{{ stringsEqualFold "Go" "go" }}`,
			"true",
		},
		{
			`{{ stringsFields "  foo bar  baz   " }}`,
			"[foo bar baz]",
		},
		{
			`{{ stringsIndex "go gopher" "go" }}`,
			"0",
		},
		{
			`{{ stringsLastIndex "go gopher" "go" }}`,
			"3",
		},
		{
			`ba{{ stringsRepeat "na" 2 }}`,
			"banana",
		},
		{
			`{{ stringsReplace "oink oink oink" "k" "ky" 2 }}`,
			"oinky oinky oink",
		},
		{
			`{{ stringsReplace "oink oink oink" "oink" "moo" -1 }}`,
			"moo moo moo",
		},
		{
			`{{ stringsSplitN "a,b,c" "," 0 }}`,
			`[]`,
		},
		{
			`{{ stringsSplitAfter "a,b,c" "," }}`,
			`[a, b, c]`,
		},
		{
			`{{ stringsSplitAfterN "a,b,c" "," 2 }}`,
			`[a, b,c]`,
		},
		{
			`{{ stringsTitle "her royal highness" }}`,
			"Her Royal Highness",
		},
		{
			`{{ stringsToLower "Gopher" }}`,
			"gopher",
		},
		{
			`{{ stringsToUpper "Gopher" }}`,
			"GOPHER",
		},
		{
			`{{ stringsTrimSpace " \t\n a lone gopher \n\t\r\n" }}`,
			"a lone gopher",
		},
		{
			`{{ stringsTrim " !!! Achtung !!! " "! " }}`,
			"Achtung",
		},
		{
			`{{ stringsTrimPrefix "Goodbye,, world!" "Goodbye," }}`,
			", world!",
		},
		{
			`{{ stringsTrimSuffix "Hello, goodbye, etc!" "goodbye, etc!" }}`,
			"Hello, ",
		},
		// aliases
		{
			`The {{if eq .StrTest "these rights belong to those people"}}eq says Yea{{else}}eq says Nay{{end}} but {{if eqf .StrTest "these rights belong to those people"}}eqf says Yea{{else}}eqf says Nay{{end}}.`,
			"The eq says Nay but eqf says Yea.",
		},

		// == regexp functions
		{
			`{{ regexpFindString "peach punch" "p([a-z]+)ch" }}`,
			"peach",
		},
		{
			`{{ regexpFindAllString "peach punch" "p([a-z]+)ch" -1 }}`,
			"[peach punch]",
		},
		{
			`{{ regexpFindAllString "peach punch" "p([a-z]+)ch" 1 }}`,
			"[peach]",
		},
		{
			`{{ regexpFindStringIndex "peach punch" "p([a-z]+)ch" }}`,
			"[0 5]",
		},
		{
			`{{ regexpFindStringSubmatch "peach punch" "p([a-z]+)ch" }}`,
			"[peach ea]",
		},
		{
			`{{ regexpFindStringSubmatchIndex "peach punch" "p([a-z]+)ch" }}`,
			"[0 5 1 3]",
		},
		//
		{
			`{{ regexpMatchString "HTTPS://site/" "(?i)^http" }}`,
			"true",
		},
		//
		{
			`{{ regexpReplaceAllLiteralString "html HTML Html aa uml bb Uml" "(?i)html|uml" "XML" }}`,
			"XML XML XML aa XML bb XML",
		},
		{
			`{{ regexpReplaceAllString .StrTest "(?i)th[eo]se" "the" }}`,
			"the rights belong to the people",
		},
		{
			`{{ regexpReplaceAllString "This and these are for THOSE people" "(?i)(this|th[eo]se)" "<b>${1}</b>" }}`,
			"<b>This</b> and <b>these</b> are for <b>THOSE</b> people",
		},
		// {
		// 	`{{ regexpReplaceAllStringFunc "a peach" "p([a-z]+)ch" stringsToUpper }}`,
		// 	"a PEACH",
		// },
	}

	testStringManipulation(t, testData)
}

/*
	{
		`{{  }}`,
		"",
	},

*/
