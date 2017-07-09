////////////////////////////////////////////////////////////////////////////
// Package: easygen
// Purpose: Easy to use universal code/text generator
// Authors: Tong Sun (c) 2015-17, All rights reserved
////////////////////////////////////////////////////////////////////////////

/*

Package easygen is an easy to use universal code/text generator library.

It can be used as a text or html generator for arbitrary purposes with arbitrary data and templates.

It can be used as a code generator, or anything that is structurally repetitive. Some command line parameter handling code generator are provided as examples, including the Go's built-in flag package, and the viper & cobra package.

Many examples have been provided to showcase its functionality, and different ways to use it.

*/
package easygen

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	tt "text/template"

	"gopkg.in/yaml.v2"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

// EgData -- EasyGen driven Data
type EgData map[interface{}]interface{}

// Opts holds the actual values from the command line parameters
var Opts = Options{ExtYaml: ".yaml", ExtTmpl: ".tmpl"}

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Generate2 will produce output according to the given template and driving data files,
// specified via fileNameTempl and fileNames (for data) respectively.
func Generate2(HTML bool, fileNameTempl string, fileNames ...string) (ret string) {
	Opts.TemplateFile = fileNameTempl

	for _, fileName := range fileNames {
		ret += Generate(HTML, fileName)
	}
	Opts.TemplateFile = ""
	return ret
}

// Generate0 will produce output according from driving data without a template file, using the string from strTempl as the template
func Generate0(HTML bool, strTempl string, fileName string) string {
	Opts.TemplateStr = strTempl
	ret := Generate(HTML, fileName)
	Opts.TemplateStr = ""
	return ret
}

// Generate will produce output from the template according to the corresponding driving data, fileName is for both template and data file name
func Generate(HTML bool, fileName string) string {
	var templates []string

	// Allow to use fileName with and without the @Opts.ExtYaml suffix.
	// If both <file>.yaml and <file> exist, prefer <file>.yaml.
	if _, err := os.Stat(fileName + Opts.ExtYaml); err == nil {
		fileName += Opts.ExtYaml
	} else if _, err = os.Stat(fileName); os.IsNotExist(err) {
		fileName += Opts.ExtYaml
	}

	// Allow to use @Opts.TemplateFile without the @Opts.ExtTmpl suffix.
	if len(Opts.TemplateFile) > 0 {
		for _, template := range strings.Split(Opts.TemplateFile, ",") {
			if _, err := os.Stat(template); os.IsNotExist(err) {
				template += Opts.ExtTmpl
			}
			templates = append(templates, template)
		}
	} else if idx := strings.LastIndex(fileName, "."); idx > 0 {
		templates = []string{fileName[:idx] + Opts.ExtTmpl}
	} else {
		templates = []string{fileName + Opts.ExtTmpl}
	}

	m := ReadYamlFile(fileName)

	env := make(map[string]string)
	for _, e := range os.Environ() {
		sep := strings.Index(e, "=")
		env[e[0:sep]] = e[sep+1:]
	}
	m["ENV"] = env
	//fmt.Printf("] %+v\n", m)

	buf := new(bytes.Buffer)
	for _, template := range templates {
		t, err := ParseFiles(HTML, template)
		checkError(err)

		err = t.Execute(buf, m)
		checkError(err)
	}
	return buf.String()
}

// ParseFiles wraps parsing text or HTML template files into a single
// function, dictated by the first parameter "HTML".
// By Matt Harden @gmail.com
func ParseFiles(HTML bool, filenames ...string) (Template, error) {
	var tname string

	if len(Opts.TemplateStr) > 0 {
		tname = "TT"
	} else if len(filenames) == 0 {
		return nil, fmt.Errorf("ParseFiles called without template filename")
	} else {
		tname = filepath.Base(filenames[0])
	}

	textTemplate := tt.New(tname).Funcs(tt.FuncMap{
		"eqf":      strings.EqualFold,
		"split":    strings.Fields,
		"minus1":   minus1,
		"dateI":    dateI,
		"year4":    year4,
		"replace":  replace,
		"replacec": replacec,
	})

	if len(Opts.TemplateStr) > 0 {
		return textTemplate.Parse(Opts.TemplateStr)
	}
	return textTemplate.ParseFiles(filenames...)
}

////////////////////////////////////////////////////////////////////////////
// Version 2 Function definitions

// ReadDataFile reads in the driving data from the given file, which can
// optinally without the defined extension
func ReadDataFile(fileName string) EgData {
	if IsExist(fileName + Opts.ExtYaml) {
		return ReadYamlFile(fileName + Opts.ExtYaml)
	} else if IsExist(fileName) {
		return ReadYamlFile(fileName)
	} else {
		checkError(errors.
			New(fmt.Sprintf("DataFile '%s' cannot be found", fileName)))
	}
	return make(EgData)
}

// ReadYamlFile reads given YAML file as EgData
func ReadYamlFile(fileName string) EgData {
	source, err := ioutil.ReadFile(fileName)
	checkError(err)

	m := make(EgData)

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

// GetEnv returns the Environment variables in a map
func GetEnv() map[string]string {
	env := make(map[string]string)
	for _, e := range os.Environ() {
		sep := strings.Index(e, "=")
		env[e[0:sep]] = e[sep+1:]
	}
	return env
}

// Process will process the standard easygen input: the `fileName` is for both template and data file names, and produce output from the template according to the corresponding driving data.
func Process(t Template, wr io.Writer, fileName string) error {
	return Process2(t, wr, fileName, fileName)
}

// Process2 will process the case that *both* template and data file names are given, and produce output according to the given template and driving data files,
// specified via fileNameTempl and fileNames respectively.
// fileNameTempl can be a comma-separated string giving many template files
func Process2(t Template, wr io.Writer, fileNameTempl string, fileNames ...string) error {
	if len(Opts.TemplateFile) > 0 {
		fileNameTempl = Opts.TemplateFile
	}
	for _, dataFn := range fileNames {
		for _, templateFn := range strings.Split(fileNameTempl, ",") {
			err := Process1(t, wr, templateFn, dataFn)
			checkError(err)
		}
	}
}

// Process1 will process a *single* case where both template and data file names are given, and produce output according to the given template and driving data files,
// specified via fileNameTempl and fileName respectively.
// fileNameTempl is not a comma-separated string, but for a single template file.
func Process1(t Template, wr io.Writer, fileNameTempl string, fileName string) error {
	m := ReadDataFile(fileName)
	m["ENV"] = GetEnv()
	//fmt.Printf("] %+v\n", m)

	// template file
	fileName = fileNameTempl
	fileNameT := fileNameTempl
	if IsExist(fileName + Opts.ExtTmpl) {
		fileNameT = fileName + Opts.ExtTmpl
	} else if IsExist(fileName) {
		fileNameT = fileName
	} else if idx := strings.LastIndex(fileName, "."); idx > 0 {
		// check for the case that fileNameTempl ends with Opts.ExtYaml
		fileName = fileName[:idx]
		if IsExist(fileName + Opts.ExtTmpl) {
			fileNameT = fileName + Opts.ExtTmpl
		}
	}
	// catch all
	if !IsExist(fileNameT) {
		checkError(errors.
			New(fmt.Sprintf("Template file '%s' cannot be found", fileNameTempl)))
	}

	tn, err := t.ParseFiles(fileNameT)
	checkError(err)

	return tn.ExecuteTemplate(wr, filepath.Base(fileNameT), m)
}

// Process0 will produce output according to the driving data *without* a template file, using the string from strTempl as the template
func Process0(t Template, wr io.Writer, strTempl string, fileNames ...string) error {
	fileName := fileNames[0]
	m := ReadDataFile(fileName)
	m["ENV"] = GetEnv()

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
