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

	// define driving data of any type
	v0 := "some-init-method"
	// https://godoc.org/github.com/go-easygen/easygen#Execute0
	// provide template string, not file name
	easygen.Execute0(tmpl, os.Stdout, "{{stringsToUpper .}}\n", v0)

	// Demo of using driving data of pure slice/array
	v1 := []string{"red", "blue", "white"}
	easygen.Execute0(tmpl, os.Stdout, "The colors are: {{range .}}{{.}}, {{end}}.\n", v1)

	// Demo output to string
	var b strings.Builder
	easygen.Execute0(tmpl, &b, "The colors are: {{range $i, $color := .}}{{$color}}{{if lt $i ($ | len | minus1)}}, {{end}}{{end}}.\n", v1)
	fmt.Print(b.String())

	// Output:
	// SOME-INIT-METHOD
	// The colors are: red, blue, white, .
	// The colors are: red, blue, white.
}

// To show the full code in GoDoc
type dummyExecute0 struct {
}
