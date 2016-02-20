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
	debug        int    // debugging `level`
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

// Opts holds the actual values from the command line parameters
var Opts Options

////////////////////////////////////////////////////////////////////////////
// Commandline definitions

func init() {

	// set default values for command line parameters
	flag.BoolVar(&Opts.HTML, "html", false,
		"treat the template file as html instead of text")
	flag.StringVar(&Opts.TemplateStr, "ts", "",
		"template string (in text)")
	flag.StringVar(&Opts.TemplateFile, "tf", "",
		".tmpl template file `name` (default: same as .yaml file)")
	flag.StringVar(&Opts.ExtYaml, "ey", ".yaml",
		"`extension` of yaml file")
	flag.StringVar(&Opts.ExtTmpl, "et", ".tmpl",
		"`extension` of template file")
	flag.StringVar(&Opts.StrFrom, "rf", "",
		"replace from, the from string used for the replace function")
	flag.StringVar(&Opts.StrTo, "rt", "",
		"replace to, the to string used for the replace function")
	flag.IntVar(&Opts.debug, "debug", 0,
		"debugging `level`")

	// Now override those default values from environment variables
	if len(Opts.TemplateStr) == 0 ||
		len(os.Getenv("EASYGEN_TS")) != 0 {
		Opts.TemplateStr = os.Getenv("EASYGEN_TS")
	}
	if len(Opts.TemplateFile) == 0 ||
		len(os.Getenv("EASYGEN_TF")) != 0 {
		Opts.TemplateFile = os.Getenv("EASYGEN_TF")
	}
	if len(Opts.ExtYaml) == 0 ||
		len(os.Getenv("EASYGEN_EY")) != 0 {
		Opts.ExtYaml = os.Getenv("EASYGEN_EY")
	}
	if len(Opts.ExtTmpl) == 0 ||
		len(os.Getenv("EASYGEN_ET")) != 0 {
		Opts.ExtTmpl = os.Getenv("EASYGEN_ET")
	}
	if len(Opts.StrFrom) == 0 ||
		len(os.Getenv("EASYGEN_RF")) != 0 {
		Opts.StrFrom = os.Getenv("EASYGEN_RF")
	}
	if len(Opts.StrTo) == 0 ||
		len(os.Getenv("EASYGEN_RT")) != 0 {
		Opts.StrTo = os.Getenv("EASYGEN_RT")
	}

}

// Usage function shows help on commandline usage
func Usage() {
	fmt.Fprintf(os.Stderr,
		"\nUsage:\n %s [flags] YamlFileName\n\nFlags:\n\n",
		progname)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr,
		"\nYamlFileName: The name for the .yaml data and .tmpl template file\n\tOnly the name part, without extension. Can include the path as well.\n")
	os.Exit(0)
}
