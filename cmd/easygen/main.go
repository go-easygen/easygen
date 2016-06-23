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
	"os"
)

import (
	"github.com/go-easygen/easygen"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

////////////////////////////////////////////////////////////////////////////
// Commandline definitions

func init() {

	// set default values for command line parameters
	flag.BoolVar(&easygen.Opts.HTML, "html", false,
		"treat the template file as html instead of text")
	flag.StringVar(&easygen.Opts.TemplateStr, "ts", "",
		"template string (in text)")
	flag.StringVar(&easygen.Opts.TemplateFile, "tf", "",
		".tmpl template file `name` (default: same as .yaml file)")
	flag.StringVar(&easygen.Opts.ExtYaml, "ey", ".yaml",
		"`extension` of yaml file")
	flag.StringVar(&easygen.Opts.ExtTmpl, "et", ".tmpl",
		"`extension` of template file")
	flag.StringVar(&easygen.Opts.StrFrom, "rf", "",
		"replace from, the from string used for the replace function")
	flag.StringVar(&easygen.Opts.StrTo, "rt", "",
		"replace to, the to string used for the replace function")
	flag.IntVar(&easygen.Opts.Debug, "debug", 0,
		"debugging `level`")

	// Now override those default values from environment variables
	if len(easygen.Opts.TemplateStr) == 0 ||
		len(os.Getenv("EASYGEN_TS")) != 0 {
		easygen.Opts.TemplateStr = os.Getenv("EASYGEN_TS")
	}
	if len(easygen.Opts.TemplateFile) == 0 ||
		len(os.Getenv("EASYGEN_TF")) != 0 {
		easygen.Opts.TemplateFile = os.Getenv("EASYGEN_TF")
	}
	if len(easygen.Opts.ExtYaml) == 0 ||
		len(os.Getenv("EASYGEN_EY")) != 0 {
		easygen.Opts.ExtYaml = os.Getenv("EASYGEN_EY")
	}
	if len(easygen.Opts.ExtTmpl) == 0 ||
		len(os.Getenv("EASYGEN_ET")) != 0 {
		easygen.Opts.ExtTmpl = os.Getenv("EASYGEN_ET")
	}
	if len(easygen.Opts.StrFrom) == 0 ||
		len(os.Getenv("EASYGEN_RF")) != 0 {
		easygen.Opts.StrFrom = os.Getenv("EASYGEN_RF")
	}
	if len(easygen.Opts.StrTo) == 0 ||
		len(os.Getenv("EASYGEN_RT")) != 0 {
		easygen.Opts.StrTo = os.Getenv("EASYGEN_RT")
	}

}

////////////////////////////////////////////////////////////////////////////
// Main

func main() {
	flag.Usage = easygen.Usage
	flag.Parse()

	// One mandatory non-flag arguments
	if len(flag.Args()) < 1 {
		easygen.Usage()
	}
	fileName := flag.Args()[0]
	easygen.TFStringsInit()

	fmt.Print(easygen.Generate(easygen.Opts.HTML, fileName))
}
