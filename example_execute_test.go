package easygen_test

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-easygen/easygen"
	"github.com/go-easygen/easygen/egCal"
	"github.com/go-easygen/easygen/egFilePath"
	"github.com/go-easygen/easygen/egVar"
)

// for standalone test, change package to `main` and the next func def to,
// func main() {
func ExampleExecute() {
	easygen.Opts.Debug = 1

	tmpl0 := easygen.NewTemplate().Customize()
	tmpl := tmpl0.Funcs(easygen.FuncDefs()).Funcs(egFilePath.FuncDefs()).
		Funcs(egVar.FuncDefs()).Funcs(egCal.FuncDefs())

	// define driving data
	v0 := easygen.EgData{"Name": "some-init-method"}
	// https://godoc.org/github.com/go-easygen/easygen#Execute
	// provide full template file name with extension
	easygen.Execute(tmpl, os.Stdout, "test/var0.tmpl", v0)

	// Demo of using driving data of slice/array
	v1 := easygen.EgData{"v": []string{"red", "blue", "white"}}
	easygen.Execute(tmpl, os.Stdout, "test/list00.tmpl", v1)

	// Demo output to string
	var b strings.Builder
	easygen.Execute(tmpl, &b, "test/list00f.tmpl", v1)
	fmt.Print(b.String())

	// Output:
	// Input: "some-init-method"
	// Output 1: "SomeInitMethod"
	// Output 2: "SOME_INIT_METHOD"
	// The colors are: red, blue, white, .
	// The colors are: red, blue, white.
}
