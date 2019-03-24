// !!! !!!
// WARNING: Code automatically generated. Editing discouraged.
// !!! !!!

package easygen

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const progname = "easygen" // os.Args[0]

// The Options struct defines the structure to hold the commandline values
type Options struct {
	TemplateStr string // template string (in text)
	ExtYaml     string // `extension` of yaml file
	ExtJson     string // `extension` of json file
	ExtTmpl     string // `extension` of template file
	Debug       int    // debugging `level`
}
