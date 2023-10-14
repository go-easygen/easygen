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
	flag.StringVar(&easygen.Opts.ExtJson, "ej", ".json",
		"`extension` of json file")
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
	if len(easygen.Opts.ExtJson) == 0 ||
		len(os.Getenv("EASYGEN_EJ")) != 0 {
		easygen.Opts.ExtJson = os.Getenv("EASYGEN_EJ")
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
		`
template_name: The name for the template file.
 - Can have the extension (specified by -et) or without it.
 - Can include the path as well.
 - Can be a comma-separated list giving many template files, in which case
   at least one data_filename must be given.

data_filename(s): The name for the .yaml or .json data.
 - If omitted derive from the template_name.
 - Can have the extension or without it. If withot extension,
   will try the .yaml file first then .json
 - Can include the path as well.
 - Can have more than one data files given on cli, separated by spaces,
   in which case multiple outputs will be produced based on the inputs.
 - Can be a comma-separated list giving many data files, in which case
   data will be merged together as if provided from a single file.
 - If there is a single data_filename, and it is "-",
   then read the data (.yaml or .json format) from stdin.

Flag defaults can be overridden by corresponding environment variable, e.g.:
  EASYGEN_EY=.yml EASYGEN_ET=.tpl easygen ...
`)
	os.Exit(0)
}
