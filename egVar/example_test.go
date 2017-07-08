package egVar_test

import (
	"os"

	"github.com/go-easygen/easygen"
	"github.com/go-easygen/easygen/egVar"
)

//==========================================================================
// listfunc2

// Test the provided listfunc2, template and data
func ExampleGenerate_listfunc2() {
	// Equivalent testing on commandline:
	//   easygen test/listfunc2

	tmpl0 := egVar.NewTemplate().Customize()
	tmpl := tmpl0.Funcs(easygen.FuncDefs()).Funcs(egVar.FuncDefs())
	easygen.Process(tmpl, os.Stdout, "../test/listfunc2")

	// Output:
	// some-init-method 5 5 someInitMethod SomeInitMethod
}
