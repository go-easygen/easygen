package easygen_test

import (
	"fmt"
	"os"

	"github.com/go-easygen/easygen"
	"github.com/go-easygen/easygen/egCal"
)

// API example for Generate
func ExampleGenerate() {
	// Equivalent testing on commandline:
	//   easygen test/listfunc1

	fmt.Println(easygen.Generate(false, "test/listfunc1"))
	// Output:
	// red, blue, white.
}

// API example for Generate2
func ExampleGenerate2() {
	// Equivalent testing on commandline:
	//   easygen -tf test/list0E test/list0

	fmt.Println(os.Getenv("SHELL"))
	fmt.Println(easygen.Generate2(false, "test/list0E", "test/list0"))
	// Output:
	// /bin/bash
	// The colors are: red, blue, white, .
	// The system shell is: /bin/bash
	// Different shells are: /bin/bash-red, /bin/bash-blue, /bin/bash-white,
}

// API example for Generate0
func ExampleGenerate0() {
	// Equivalent testing on commandline:
	//   easygen -ts '{{range .Colors}}{{.}}, {{end}}' test/list0

	fmt.Println(easygen.Generate0(false, "{{range .Colors}}{{.}}, {{end}}", "test/list0"))
	// Output:
	// red, blue, white,
}

// API V2 example
func ExampleAPI2() {
	tmpl0 := easygen.NewTemplate().Customize()
	tmpl := tmpl0.Funcs(easygen.FuncDefs())
	easygen.Process(tmpl, os.Stdout, "test/var0")

	// Output:
	// Input: "some-init-method"
	// Output 1: "SomeInitMethod"
	// Output 2: "SOME_INIT_METHOD"
}

func ExampleAPI2_egCal() {
	tmpl0 := easygen.NewTemplate().Customize()
	tmpl := tmpl0.Funcs(easygen.FuncDefs()).Funcs(egCal.FuncDefs())

	tmpl, _ = tmpl.Parse("{{.}}: {{add 2 3}}, {{multiply 2 3}}, {{subtract 9 2}}, {{divide 24 3}}\n")
	tmpl.Execute(os.Stdout, "the go-programming-language")

	// Output:
	// the go-programming-language: 5, 6, 7, 8
}

// for standalone test, change package to `main` and the next func def to,
// func main() {
func Example_main() {
	ExampleGenerate()
	ExampleGenerate2()
	ExampleGenerate0()
	fmt.Println()

	// API v2 Usage:
	ExampleAPI2()
	ExampleAPI2_egCal()
}

// To show the full code in GoDoc
type dummy struct {
}
