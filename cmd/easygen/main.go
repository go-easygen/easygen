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
	"fmt"

	"github.com/go-easygen/easygen"
)

//go:generate sh -v easygen.gen.sh

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

////////////////////////////////////////////////////////////////////////////
// Main

func main() {
	flag.Usage = Usage
	flag.Parse()

	// One mandatory non-flag arguments
	if flag.NArg() < 1 {
		Usage()
	}
	easygen.TFStringsInit()

	for _, fileName := range flag.Args() {
		fmt.Print(easygen.Generate(easygen.Opts.HTML, fileName))
	}
}
