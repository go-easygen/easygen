////////////////////////////////////////////////////////////////////////////
// Package: easygen
// Purpose: Easy to use universal code/text generator
// Authors: Tong Sun (c) 2015-2019, All rights reserved
////////////////////////////////////////////////////////////////////////////

/*

Package easygen is an easy to use universal code/text generator library.

It can be used as a text or html generator for arbitrary purposes with arbitrary data and templates.

It can be used as a code generator, or anything that is structurally repetitive. Some command line parameter handling code generator are provided as examples, including the Go's built-in flag package, and the viper & cobra package.

Many examples have been provided to showcase its functionality, and different ways to use it.

*/
package easygen

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

// EgData -- EasyGen driven Data
type EgData interface{}

// Opts holds the actual values from the command line parameters
var Opts = Options{ExtYaml: ".yaml", ExtTmpl: ".tmpl"}

////////////////////////////////////////////////////////////////////////////
// Function definitions

// ReadDataFile reads in the driving data from the given file, which can
// be optionally without the defined extension
func ReadDataFile(fileName string) EgData {
	if IsExist(fileName + Opts.ExtYaml) {
		return ReadYamlFile(fileName + Opts.ExtYaml)
	} else if IsExist(fileName) {
		return ReadYamlFile(fileName)
	}
	checkError(fmt.Errorf("DataFile '%s' cannot be found", fileName))
	return nil
}

// ReadYamlFile reads given YAML file as EgData
func ReadYamlFile(fileName string) EgData {
	source, err := ioutil.ReadFile(fileName)
	checkError(err)

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal(source, &m)
	checkError(err)

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

// Process will process the standard easygen input: the `fileName` is for both template and data file names, and produce output from the template according to the corresponding driving data.
func Process(t Template, wr io.Writer, fileNames ...string) error {
	return Process2(t, wr, fileNames[0], fileNames...)
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
func Process1(t Template, wr io.Writer, fileNameTempl string, fileName string) error {
	m := ReadDataFile(fileName)
	//fmt.Printf("] %+v\n", m)

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

// Execute will execute the Template on the given data map `m`.
func Execute(t Template, wr io.Writer, fileNameT string, m EgData) error {
	if !IsExist(fileNameT) {
		ex, e := os.Executable()
		if e != nil {
			return e
		}
		fileNameT = filepath.Dir(ex) + string(filepath.Separator) + fileNameT
		if !IsExist(fileNameT) {
			checkError(fmt.Errorf("Template file '%s' cannot be found", fileNameT))
		}
	}

	tn, err := t.ParseFiles(fileNameT)
	checkError(err)

	return tn.ExecuteTemplate(wr, filepath.Base(fileNameT), m)
}

// Process0 will produce output according to the driving data *without* a template file, using the string from strTempl as the template
func Process0(t Template, wr io.Writer, strTempl string, fileNames ...string) error {
	fileName := fileNames[0]
	m := ReadDataFile(fileName)

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
