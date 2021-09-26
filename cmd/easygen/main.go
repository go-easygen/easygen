////////////////////////////////////////////////////////////////////////////
// Porgram: easygen
// Purpose: Easy to use universal code/text generator
// Authors: Tong Sun (c) 2015-2021, All rights reserved
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
	"strings"

	"github.com/go-easygen/easygen"
	"github.com/go-easygen/easygen/egCal"
	"github.com/go-easygen/easygen/egFilePath"
	"github.com/go-easygen/easygen/egVar"
)

//go:generate sh -v easygen.gen.sh

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname = "easygen"
	version  = "5.1.7"
	date     = "2021-09-25"
)

////////////////////////////////////////////////////////////////////////////
// Main

func main() {
	flag.Usage = Usage
	flag.Parse()

	// One mandatory non-flag arguments
	if flag.NArg() < 1 {
		Usage()
	}

	args := flag.Args()
	// There is only one command line argument
	if len(args) == 1 {
		// when template_name is comma-separated list, data_filename must be given
		if strings.Contains(args[0], ",") {
			Usage()
		}
		// else, dup template_name as data_filename
		args = append(args, args[0])
	}

	tmpl0 := easygen.NewTemplate().Customize()
	tmpl := tmpl0.Funcs(easygen.FuncDefs()).
		Funcs(egVar.FuncDefs()).Funcs(egCal.FuncDefs()).Funcs(egFilePath.FuncDefs())

	var err error
	if len(easygen.Opts.TemplateStr) > 0 {
		err = easygen.Process0(tmpl, os.Stdout, easygen.Opts.TemplateStr, args...)
	} else {
		err = easygen.Process2(tmpl, os.Stdout, args[0], args[1:]...)
	}
	if err != nil {
		panic(err)
	}
}
