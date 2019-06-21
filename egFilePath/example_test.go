package egFilePath_test

import (
	"os"

	"github.com/go-easygen/easygen"
	"github.com/go-easygen/easygen/egFilePath"
)

// for standalone test, change package to `main` and the next func def to,
// func main() {
func Example() {
	tmpl0 := easygen.NewTemplate().Customize()
	tmpl := tmpl0.Funcs(easygen.FuncDefs()).Funcs(egFilePath.FuncDefs())
	err := easygen.Process0(tmpl, os.Stdout,
		// https://godoc.org/path/filepath
		"Base: {{fpBase `/a/b.c`}}, {{fpBase `/b.c/`}}, {{fpBase `./b.c`}}, {{fpBase `b.c`}}\n"+
			"Clean: {{fpClean `/a//b.c/./..`}}, {{fpClean `//b.c///`}}, {{fpClean `/../b.c`}}\n"+
			"Dir: {{fpDir `/a/b/c`}}, {{fpDir `/b/c/`}}, {{fpDir `./b/c`}}, {{fpDir `b.c`}}\n"+
			"Ext: {{fpExt `index`}}, {{fpExt `index.js`}}, {{fpExt `main.test.js`}}\n"+
			"Join: {{fpJoin `/a` `b` `c` `a//` `//b////c`}}, {{fpJoin `a` `b/c`}}, {{fpJoin `a/b` `c`}}\n"+
			"Rel: {{fpRel `/a` `/a/b/c`}}, {{fpRel `/a` `/b/c`}}\n"+

			// Two error cases, uncomment to verify
			// "{{fpRel `/a` `./b/c`}}, {{isDir `not-exist`}}.\n"+

			"IsDir: {{isDir `.`}}.\n"+
			"Basename: {{basename `/a/b.c`}}, {{basename `/b.c/`}}, {{basename `b.c`}}, {{basename `bc`}}",
		"../test/var0")

	// Output:
	// Base: b.c, b.c, b.c, b.c
	// Clean: /a, /b.c, /b.c
	// Dir: /a/b, /b/c, b, .
	// Ext: , .js, .js
	// Join: /a/b/c/a/b/c, a/b/c, a/b/c
	// Rel: b/c, ../b/c
	// IsDir: true.
	// Basename: b, b, b, bc

	if err != nil {
		panic(err)
	}
}

// To show the full code in GoDoc
type dummy struct {
}
