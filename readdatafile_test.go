package easygen_test

import (
	"os"

	"github.com/go-easygen/easygen"
	"github.com/go-easygen/easygen/egVar"
)

// for standalone test, change package to `main` and the next func def to,
// func main() {
func ExampleReadDataFile() {
	tmplFileName := "test/var0"
	tmpl0 := easygen.NewTemplate().Customize()
	tmpl := tmpl0.Funcs(easygen.FuncDefs()).Funcs(egVar.FuncDefs())

	// To use Execute(), TemplateFileName has to be exact
	tmplFileNameFull := tmplFileName + ".tmpl"

	m := easygen.ReadDataFile(tmplFileName)
	easygen.Execute(tmpl, os.Stdout, tmplFileNameFull, m)

	easygen.Opts.Debug = 0
	m = easygen.ReadDataFile(tmplFileName + ".yaml")
	easygen.Execute(tmpl, os.Stdout, tmplFileNameFull, m)

	tmplFileName = "test/list0j"
	tmplFileNameFull = tmplFileName + ".tmpl"

	m = easygen.ReadDataFile(tmplFileName)
	easygen.Execute(tmpl, os.Stdout, tmplFileNameFull, m)

	m = easygen.ReadDataFile(tmplFileName + ".json")
	easygen.Execute(tmpl, os.Stdout, tmplFileNameFull, m)

	// Output:
	// Input: "some-init-method"
	// Output 1: "SomeInitMethod"
	// Output 2: "SOME_INIT_METHOD"
	// Input: "some-init-method"
	// Output 1: "SomeInitMethod"
	// Output 2: "SOME_INIT_METHOD"
	// The colors are: red, blue, white, .
	// The colors are: red, blue, white, .

}
