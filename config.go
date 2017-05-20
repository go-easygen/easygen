// !!! !!!
// WARNING: Code automatically generated. Editing discouraged.
// !!! !!!

package easygen

import (
	"flag"
	"fmt"
	"os"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const progname = "easygen" // os.Args[0]

// The Options struct defines the structure to hold the commandline values
type Options struct {
	HTML         bool   // treat the template file as html instead of text
	TemplateStr  string // template string (in text)
	TemplateFile string // .tmpl template file `name` (default: same as .yaml file)
	ExtYaml      string // `extension` of yaml file
	ExtTmpl      string // `extension` of template file
	StrFrom      string // replace from, the from string used for the replace function
	StrTo        string // replace to, the to string used for the replace function
	Debug        int    // debugging `level`
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

// Opts holds the actual values from the command line parameters
var Opts = Options{ExtYaml: ".yaml", ExtTmpl: ".tmpl"}

////////////////////////////////////////////////////////////////////////////
// Commandline definitions

// Usage function shows help on commandline usage
func Usage() {
	fmt.Fprintf(os.Stderr,
		"\nUsage:\n %s [flags] YamlFileName [YamlFileName...]\n\nFlags:\n\n",
		progname)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr,
		"\nYamlFileName: The name for the .yaml data and .tmpl template file\n\tOnly the name part, without extension. Can include the path as well.\n")
	os.Exit(0)
}
