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

//const progname = "easygen" // os.Args[0]

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

////////////////////////////////////////////////////////////////////////////
// Commandline definitions

func init() {

	// set default values for command line parameters
	flag.StringVar(&easygen.Opts.TemplateStr, "ts", "",
		"template string (in text)")
	flag.StringVar(&easygen.Opts.ExtYaml, "ey", ".yaml",
		"`extension` of yaml file")
	flag.StringVar(&easygen.Opts.ExtTmpl, "et", ".tmpl",
		"`extension` of template file")
	flag.IntVar(&easygen.Opts.Debug, "debug", 0,
		"debugging `level`")

	// Now override those default values from environment variables
	if len(easygen.Opts.TemplateStr) == 0 ||
		len(os.Getenv("EASYGEN_TS")) != 0 {
		easygen.Opts.TemplateStr = os.Getenv("EASYGEN_TS")
	}
	if len(easygen.Opts.ExtYaml) == 0 ||
		len(os.Getenv("EASYGEN_EY")) != 0 {
		easygen.Opts.ExtYaml = os.Getenv("EASYGEN_EY")
	}
	if len(easygen.Opts.ExtTmpl) == 0 ||
		len(os.Getenv("EASYGEN_ET")) != 0 {
		easygen.Opts.ExtTmpl = os.Getenv("EASYGEN_ET")
	}

}

// Usage function shows help on commandline usage
func Usage() {
	fmt.Fprintf(os.Stderr,
		"%s version %s\n\nUsage:\n %s [flags] template_name [data_filename [data_filename...]]\n\nFlags:\n\n",
		progname, version, progname)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr,
		"\ndata_filename(s): The name for the .yaml or .json data.\n - If omitted derive from the template_name.\n - Can have the extension or without it. If withot extension,\n   will try .yaml first then .json\n - Can include the path as well.\n\ntemplate_name: The name for the template file.\n - Can have the extension or without it.\n - Can include the path as well.\n - Can be a comma-separated list giving many template files, in which case\n   at least one data_filename must be given.\n\nFlag defaults can be overridden by corresponding environment variable, e.g.:\n  EASYGEN_EY=.yml EASYGEN_ET=.tpl easygen ...\n")
	os.Exit(0)
}
