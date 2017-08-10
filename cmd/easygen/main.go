////////////////////////////////////////////////////////////////////////////
// Porgram: easygen
// Purpose: Easy to use universal code/text generator
// Authors: Tong Sun (c) 2015-16, All rights reserved
////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////
// Program start

/*

Command easygen is an easy to use universal code/text generator.

It can be used as a text or html generator for arbitrary purposes with arbitrary data and templates.

It can be used as a code generator, or anything that is structurally repetitive. Some command line parameter handling code generator are provided as examples, including the Go's built-in flag package, and the viper & cobra package.

You can even use easygen as a generic Go template testing tool using the -ts commandline option.

*/
package main

import (
	"flag"
	"os"

	"github.com/go-easygen/easygen"
	"github.com/go-easygen/easygen/egCal"
	"github.com/go-easygen/easygen/egVar"
)

//go:generate sh -v easygen.gen.sh

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var version = "master"

////////////////////////////////////////////////////////////////////////////
// Main

func main() {
	flag.Usage = Usage
	flag.Parse()

	// One mandatory non-flag arguments
	if flag.NArg() < 1 {
		Usage()
	}
	easygen.TFStringsInit() // Done *after* flag parsing!

	tmpl0 := easygen.NewTemplate().Customize()
	tmpl := tmpl0.Funcs(easygen.FuncDefs()).
		Funcs(egVar.FuncDefs()).Funcs(egCal.FuncDefs())

	args := flag.Args()
	if len(easygen.Opts.TemplateStr) > 0 {
		easygen.Process0(tmpl, os.Stdout, easygen.Opts.TemplateStr, args...)
	} else {
		easygen.Process(tmpl, os.Stdout, args...)
	}
}
