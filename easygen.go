////////////////////////////////////////////////////////////////////////////
// Package: easygen
// Purpose: Easy to use universal code/text generator
// Authors: Tong Sun (c) 2015-2021, All rights reserved
////////////////////////////////////////////////////////////////////////////

/*

Package easygen is an easy to use universal code/text generator library.

It can be used as a text or html generator for arbitrary purposes with arbitrary data and templates.

It can be used as a code generator, or anything that is structurally repetitive. Some command line parameter handling code generator are provided as examples, including the Go's built-in flag package, and the viper & cobra package.

Many examples have been provided to showcase its functionality, and different ways to use it.

*/
package easygen

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

// EgData, EasyGen key type
type EgKey = string

// EgData, EasyGen driven Data
type EgData map[EgKey]interface{}

// Opts holds the actual values from the command line parameters
var Opts = Options{ExtYaml: ".yaml", ExtJson: ".json", ExtTmpl: ".tmpl"}

////////////////////////////////////////////////////////////////////////////
// Function definitions

// ReadDataFiles reads in the driving data from the given file, which can
// be optionally without the defined extension, and can be a comma-separated
// string for multiple data files.
func ReadDataFiles(fileName string) EgData {
	var m EgData
	for _, dataFn := range strings.Split(fileName, ",") {
		m = ReadDataFile(dataFn, m)
		if Opts.Debug >= 1 {
			fmt.Fprintf(os.Stderr, "[%s] After reading file %s:\n  %+v\n", progname, dataFn, m)
		}
	}
	return m
}

// ReadDataFile reads in the driving data from the given file, which can
// be optionally without the defined extension
func ReadDataFile(fileName string, ms ...EgData) EgData {
	if IsExist(fileName + Opts.ExtYaml) {
		return ReadYamlFile(fileName+Opts.ExtYaml, ms...)
	} else if IsExist(fileName + Opts.ExtJson) {
		return ReadJsonFile(fileName+Opts.ExtJson, ms...)
	} else if IsExist(fileName) {
		verbose("Reading exist Data File", 3)
		fext := filepath.Ext(fileName)
		fext = fext[1:] // ignore the leading "."
		if regexp.MustCompile(`(?i)^y`).MatchString(fext) {
			verbose("Reading YAML file", 3)
			return ReadYamlFile(fileName, ms...)
		} else if regexp.MustCompile(`(?i)^j`).MatchString(fext) {
			return ReadJsonFile(fileName, ms...)
		} else {
			checkError(fmt.Errorf("Unsupported file extension for DataFile '%s'", fileName))
		}
	} else if fileName == "-" {
		// from stdin
		// Yaml format is a superset of JSON, it read Json file just as fine
		return ReadYamlFile(fileName)
	}
	checkError(fmt.Errorf("DataFile '%s' cannot be found", fileName))
	return nil
}

// ReadYamlFile reads given YAML file as EgData
func ReadYamlFile(fileName string, ms ...EgData) EgData {
	var source []byte
	var err error
	if fileName == "-" {
		source, err = ioutil.ReadAll(os.Stdin)
		checkError(err)
	} else {
		source, err = ioutil.ReadFile(fileName)
		checkError(err)
	}

	m := EgData{}
	if len(ms) > 0 {
		m = ms[0]
	}

	err = yaml.Unmarshal(source, &m)
	checkError(err)

	return m
}

// ReadJsonFile reads given JSON file as EgData
func ReadJsonFile(fileName string, ms ...EgData) EgData {
	source, err := ioutil.ReadFile(fileName)
	checkError(err)

	m := EgData{}
	if len(ms) > 0 {
		m = ms[0]
	}

	err = json.Unmarshal(source, &m)
	checkError(err)

	//fmt.Printf("] Input %v\n", m)
	return m
}

// IsExist checks if the given file exist
func IsExist(fileName string) bool {
	//fmt.Printf("] Checking %s\n", fileName)
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
	// CAUTION! os.IsExist(err) != !os.IsNotExist(err)
	// https://gist.github.com/mastef/05f46d3ab2f5ed6a6787#file-isexist_vs_isnotexist-go-L35-L56
}

// Process will process the standard easygen input: the `fileName` is for both template and data file name, and produce output from the template according to the corresponding driving data.
// Process() is using the V3's calling convention and *only* works properly in V4+ in the case that there is only one fileName passed to it. If need to pass more files, use Process2() instead.
func Process(t Template, wr io.Writer, fileNames ...string) error {
	return Process2(t, wr, fileNames[0], fileNames[:1]...)
}

// Process2 will process the case that *both* template and data file names are given, and produce output according to the given template and driving data files,
// specified via fileNameTempl and fileNames respectively.
// fileNameTempl can be a comma-separated string giving many template files
func Process2(t Template, wr io.Writer, fileNameTempl string, fileNames ...string) error {
	for _, dataFn := range fileNames {
		for _, templateFn := range strings.Split(fileNameTempl, ",") {
			err := Process1(t, wr, templateFn, dataFn)
			checkError(err)
		}
	}
	return nil
}

// Process1 will process a *single* case where both template and data file names are given, and produce output according to the given template and driving data files,
// specified via fileNameTempl and fileName respectively.
// fileNameTempl is not a comma-separated string, but for a single template file.
// However, the fileName can be a comma-separated string for multiple data files.
func Process1(t Template, wr io.Writer, fileNameTempl string, fileName string) error {
	m := ReadDataFiles(fileName)

	// template file
	fileName = fileNameTempl
	fileNameT := fileNameTempl
	if IsExist(fileName + Opts.ExtTmpl) {
		fileNameT = fileName + Opts.ExtTmpl
	} else {
		// guard against that fileNameTempl passed with Opts.ExtYaml extension
		if fileName[len(fileName)-len(Opts.ExtYaml):] == Opts.ExtYaml {
			idx := strings.LastIndex(fileName, ".")
			fileName = fileName[:idx]
			if IsExist(fileName + Opts.ExtTmpl) {
				fileNameT = fileName + Opts.ExtTmpl
			}
		} else if IsExist(fileName) {
			// fileNameTempl passed with Opts.ExtTmpl already
			fileNameT = fileName
		}
	}

	return Execute(t, wr, fileNameT, m)
}

// Execute0 will execute the Template given as strTempl with the given data map `m` (i.e., no template file and no data file).
// It parses text template strTempl then applies it to to the specified data
// object m, and writes the output to wr. If an error occurs executing the
// template or writing its output, execution stops, but partial results may
// already have been written to the output writer. A template may be
// executed safely in parallel, although if parallel executions share a
// Writer the output may be interleaved.
func Execute0(t Template, wr io.Writer, strTempl string, m EgData) error {
	verbose("Execute with template string: "+strTempl, 2)
	tmpl, err := t.Parse(strTempl)
	checkError(err)
	return tmpl.Execute(wr, m)
}

// Execute will execute the Template from fileNameT on the given data map `m`.
func Execute(t Template, wr io.Writer, fileNameT string, m EgData) error {
	// 1. Check locally
	verbose("Checking for template locally: "+fileNameT, 2)
	if !IsExist(fileNameT) {
		// 2. Check under /etc/
		command := filepath.Base(os.Args[0])
		templateFile := fmt.Sprintf("/etc/%s/%s", command, fileNameT)
		verbose("Checking at "+templateFile, 2)
		if IsExist(templateFile) {
			fileNameT = templateFile
		} else {
			// 3. Check where executable is
			ex, e := os.Executable()
			if e != nil {
				return e
			}
			fileNameT = filepath.Dir(ex) + string(filepath.Separator) + fileNameT
			verbose("Checking at "+fileNameT, 2)
			if !IsExist(fileNameT) {
				checkError(fmt.Errorf("Template file '%s' cannot be found", fileNameT))
			}
		}
	}

	tn, err := t.ParseFiles(fileNameT)
	checkError(err)

	return tn.ExecuteTemplate(wr, filepath.Base(fileNameT), m)
}

// Process0 will produce output according to the driving data *without* a template file, using the string from strTempl as the template
func Process0(t Template, wr io.Writer, strTempl string, fileNames ...string) error {
	fileName := fileNames[0]
	m := ReadDataFiles(fileName)

	tmpl, err := t.Parse(strTempl)
	checkError(err)
	return tmpl.Execute(wr, m)
}

////////////////////////////////////////////////////////////////////////////
// Support Function definitions

// Exit if error occurs
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[%s] Fatal error - %s\n", progname, err)
		os.Exit(1)
	}
}

// verbose will print info to stderr according to the verbose level setting
func verbose(step string, level int) {
	if Opts.Debug >= level {
		print("[", progname, "] ", step, "\n")
	}
}
