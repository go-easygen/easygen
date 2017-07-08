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
	"fmt"
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

	source, err := ioutil.ReadFile(fileName)
	checkError(err)

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal(source, &m)
	checkError(err)

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
		"cls2lc":   cls2lc.String,
		"cls2uc":   cls2uc.String,
		"cls2ss":   cls2ss.String,
		"ck2lc":    ck2lc.String,
		"ck2uc":    ck2uc.String,
		"ck2ls":    ck2ls.String,
		"ck2ss":    ck2ss.String,
		"clc2ss":   clc2ss.String,
		"cuc2ss":   cuc2ss.String,
	})

	if len(Opts.TemplateStr) > 0 {
		return textTemplate.Parse(Opts.TemplateStr)
	}
	return textTemplate.ParseFiles(filenames...)
}

// Exit if error occurs
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[%s] Fatal error - %s\n", progname, err)
		os.Exit(1)
	}
}
