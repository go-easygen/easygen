// !!! !!!
// WARNING: Code automatically generated. Editing discouraged.
// !!! !!!

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-easygen/easygen"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const progname = "easygen" // os.Args[0]

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
		".tmpl (comma-separated) template file `name(s)` (default: same as .yaml file)")
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

// Usage function shows help on commandline usage
func Usage() {
	fmt.Fprintf(os.Stderr,
		"%s version %s\n\nUsage:\n %s [flags] YamlFileName [YamlFileName...]\n\nFlags:\n\n",
		progname, version, progname)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr,
		"\nYamlFileName(s): The name for the .yaml data and .tmpl template file\n\tCan have the extension or without it. Can include the path as well.\n\nFlag defaults can be overridden by corresponding environment variable, e.g.:\n  EASYGEN_EY=.yml EASYGEN_ET=.tpl easygen ...\n")
	os.Exit(0)
}
