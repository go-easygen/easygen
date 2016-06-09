////////////////////////////////////////////////////////////////////////////
// Package: easygen
// Purpose: Easy to use universal code/text generator
// Authors: Tong Sun (c) 2015-16, All rights reserved
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
	ht "html/template"
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

// common type for a *(text|html).Template value
type template interface {
	Execute(wr io.Writer, data interface{}) error
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	Name() string
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Generate2 will produce output according to the given template and driving data files, specified via fileNameTempl and fileName (for data) respectively
func Generate2(HTML bool, fileNameTempl string, fileName string) string {
	Opts.TemplateFile = fileNameTempl
	ret := Generate(HTML, fileName)
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
	source, err := ioutil.ReadFile(fileName + Opts.ExtYaml)
	checkError(err)

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal(source, &m)
	checkError(err)

	// template file name
	fileNameT := fileName
	if len(Opts.TemplateFile) > 0 {
		fileNameT = Opts.TemplateFile
	}

	t, err := ParseFiles(HTML, fileNameT+Opts.ExtTmpl)
	checkError(err)

	buf := new(bytes.Buffer)
	err = t.Execute(buf, m)
	checkError(err)

	return buf.String()
}

// ParseFiles wraps parsing text or HTML template files into a single
// function, dictated by the first parameter "HTML".
// By Matt Harden @gmail.com
func ParseFiles(HTML bool, filenames ...string) (template, error) {
	tname := filepath.Base(filenames[0])

	// use text template
	funcMapHT := ht.FuncMap{
		"minus1": minus1,
		"cls2lc": cls2lc.String,
		"cls2uc": cls2uc.String,
		"cls2ss": cls2ss.String,
		"ck2lc":  ck2lc.String,
		"ck2uc":  ck2uc.String,
		"ck2ls":  ck2ls.String,
		"ck2ss":  ck2ss.String,
		"clc2ss": clc2ss.String,
		"cuc2ss": cuc2ss.String,
	}

	_ = funcMapHT

	if HTML {
		// use html template
		t, err := ht.ParseFiles(filenames...)
		//t, err := ht.New("HT").Funcs(funcMapHT).ParseFiles(filenames...)
		return t, err
	}

	// use text template
	funcMap := tt.FuncMap{
		"eqf":      strings.EqualFold,
		"split":    strings.Fields,
		"minus1":   minus1,
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
	}

	if len(Opts.TemplateStr) > 0 {
		t, err := tt.New("TT").Funcs(funcMap).Parse(Opts.TemplateStr)
		return t, err
	}

	t, err := tt.New(tname).Funcs(funcMap).ParseFiles(filenames...)
	return t, err
}

// Exit if error occurs
func checkError(err error) {
	if err != nil {
		fmt.Printf("[%s] Fatal error - %v", progname, err.Error())
		os.Exit(1)
	}
}
