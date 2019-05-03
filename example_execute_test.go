package easygen_test

import (
	"os"

	"github.com/go-easygen/easygen"
	"github.com/go-easygen/easygen/egCal"
	"github.com/go-easygen/easygen/egVar"
)

type variable struct {
	Name string
}

// for standalone test, change package to `main` and the next func def to,
// func main() {
func ExampleExecute() {
	tmpl0 := easygen.NewTemplate().Customize()
	tmpl := tmpl0.Funcs(easygen.FuncDefs()).Funcs(egVar.FuncDefs()).Funcs(egCal.FuncDefs())

	// define driving data of any tye
	v0 := variable{"some-init-method"}
	easygen.Opts.Debug = 1

	// https://godoc.org/github.com/go-easygen/easygen#Execute
	// provide full template file name with extension
	easygen.Execute(tmpl, os.Stdout, "test/var0.tmpl", v0)

	// Output:
	// Input: "some-init-method"
	// Output 1: "SomeInitMethod"
	// Output 2: "SOME_INIT_METHOD"
}

// To show the full code in GoDoc
type dummyExecute struct {
}
