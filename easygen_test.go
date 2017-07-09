package easygen_test

import (
	"os"
	"text/template"

	"github.com/go-easygen/easygen"
)

var tmpl *template.Template

func init() {
	tmpl = easygen.NewTemplate().Funcs(easygen.FuncDefs())
}

////////////////////////////////////////////////////////////////////////////
// Lists

//==========================================================================
// list0 data + string template
// I.e.: EasyGen -ts "{{range .Colors}}{{.}}, {{end}}" test/list0

// Test string template with list0 data
func ExampleProcess0_list0StrTemplate() {
	// Equivalent testing on commandline:
	//   easygen -ts '{{range .Colors}}{{.}}, {{end}}' test/list0
	easygen.Process0(tmpl, os.Stdout,
		"{{range .Colors}}{{.}}, {{end}}", "test/list0")
	// Output:
	// red, blue, white,
}

////////////////////////////////////////////////////////////////////////////
// Strings Test

// Test string comparison in template
func ExampleProcess0_stringsCmp() {
	// Equivalent testing on commandline:
	//   easygen -ts '{{The {{if ... {{end}}.' test/strings0
	easygen.Process0(tmpl, os.Stdout,
		`The {{if eq .StrTest "-AB-axxb- HTML Html html"}}eq says Yea{{else}}eq says Nay{{end}} but {{if eqf .StrTest "-AB-axxb- HTML Html html"}}eqf says Yea{{else}}eqf says Nay{{end}}.`, "test/strings0")
	// Output:
	// The eq says Nay but eqf says Yea.
}

// Test the string split function in template
func ExampleProcess0_split0() {
	// Equivalent testing on commandline:
	//   easygen -ts '{{split .Colorlist}}' test/list0
	easygen.Process0(tmpl, os.Stdout,
		`{{split .Colorlist}}`, "test/list0")
	// Output:
	// [red blue white]
}

// Test the string split function in template again
func ExampleProcess0_split1() {
	// Equivalent testing on commandline:
	//   easygen -ts '{{range ... {{end}}' test/list0
	easygen.Process0(tmpl, os.Stdout,
		`{{range (split .Colorlist)}}{{.}} {{end}}`, "test/list0")
	// Output:
	// red blue white
}
