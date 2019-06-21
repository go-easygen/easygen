package egVar_test

import (
	"os"

	"github.com/go-easygen/easygen"
	"github.com/go-easygen/easygen/egVar"
)

// for standalone test, change package to `main` and the next func def to,
// func main() {
func Example() {
	// Equivalent testing on commandline:
	//   easygen test/listfunc2

	tmpl0 := egVar.NewTemplate().Customize()
	tmpl := tmpl0.Funcs(easygen.FuncDefs()).Funcs(egVar.FuncDefs())
	err := easygen.Process(tmpl, os.Stdout, "../test/listfunc2")

	// Output:
	// some-init-method 5 5 someInitMethod SomeInitMethod

	if err != nil {
		panic(err)
	}
}

// To show the full code in GoDoc
type dummy struct {
}
