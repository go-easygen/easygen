////////////////////////////////////////////////////////////////////////////
// Porgram: EasyGen
// Purpose: Easy to use universal code/text generator
// Authors: Tong Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////
// Program start

package easygenapi

import (
	"bytes"
	"fmt"
	ht "html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
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

// Generate will produce output from the template according to driving data
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

	t, err := parseFiles(HTML, fileNameT+Opts.ExtTmpl)
	checkError(err)

	buf := new(bytes.Buffer)
	err = t.Execute(buf, m)
	checkError(err)

	return buf.String()
}

// parseFiles, intialization. By Matt Harden @gmail.com
func parseFiles(HTML bool, filenames ...string) (template, error) {
	tname := filepath.Base(filenames[0])

	if HTML {
		// use html template
		t, err := ht.ParseFiles(filenames...)
		return t, err
	}

	// use text template
	funcMap := tt.FuncMap{
		"minus1": minus1,
		"cls2lc": cls2lc.String,
		"cls2uc": cls2uc.String,
		"cls2ss": cls2ss.String,
		"ck2lc":  ck2lc.String,
		"ck2uc":  ck2uc.String,
		"ck2ss":  ck2ss.String,
		"clc2ss": clc2ss.String,
		"cuc2ss": cuc2ss.String,
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
