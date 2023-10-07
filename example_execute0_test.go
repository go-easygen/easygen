package easygen_test

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-easygen/easygen"
)

// for standalone test, change package to `main` and the next func def to,
// func main() {
func ExampleExecute0() {
	easygen.Opts.Debug = 1

	tmpl := easygen.NewTemplate().Funcs(easygen.FuncDefs())

	// define driving data
	v0 := easygen.EgData{"v": "some-init-method"}
	// https://godoc.org/github.com/go-easygen/easygen#Execute0
	// provide template string, not file name
	easygen.Execute0(tmpl, os.Stdout, "{{stringsToUpper .v}}\n", v0)

	// Demo of using driving data of slice/array
	v1 := easygen.EgData{"v": []string{"red", "blue", "white"}}
	easygen.Execute0(tmpl, os.Stdout, "The colors are: {{range .v}}{{.}}, {{end}}.\n", v1)

	// Demo output to string
	var b strings.Builder
	easygen.Execute0(tmpl, &b, "The colors are: {{range $i, $color := .v}}{{$color}}{{if lt $i ($.v | len | minus1)}}, {{end}}{{end}}.\n", v1)
	fmt.Print(b.String())

	// Output:
	// SOME-INIT-METHOD
	// The colors are: red, blue, white, .
	// The colors are: red, blue, white.
}

// To show the full code in GoDoc
type dummyExecute0 struct {
}
