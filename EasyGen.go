////////////////////////////////////////////////////////////////////////////
// Porgram: EasyGen
// Purpose: Universal code/text generator that is easy to use
// Authors: Tong Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////
// Program start

package main

import (
	"bytes"
	"flag"
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

const progname = "EasyGen" // os.Args[0]

// The Options structure holds the values for/from commandline
type Options struct {
	HTML bool
}

// common type for a *(text|html).Template value
type template interface {
	Execute(wr io.Writer, data interface{}) error
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	Name() string
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var opts Options

////////////////////////////////////////////////////////////////////////////
// Commandline definitions

func init() {
	flag.BoolVar(&opts.HTML, "html", false, "Use html template instead of text")
}

func usage() {
	fmt.Fprintf(os.Stderr, "\nUsage:\n %s [flags] TemplateFileName\n\nFlags:\n\n",
		progname)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nTemplateFileName: The name for the template and yaml file\n\tOnly the name part, without extension. Can include the path as well.\n")
	os.Exit(0)
}

////////////////////////////////////////////////////////////////////////////
// Main

func main() {
	flag.Usage = usage
	flag.Parse()

	// One mandatory non-flag arguments
	if len(flag.Args()) < 1 {
		usage()
	}
	fileName := flag.Args()[0]

	fmt.Print(Generate(opts.HTML, fileName))
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Generate will produce output from the template according to driving data
func Generate(HTML bool, fileName string) string {
	source, err := ioutil.ReadFile(fileName + ".yaml")
	checkError(err)

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal(source, &m)
	checkError(err)

	t, err := parseFiles(HTML, fileName+".tmpl")
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
		"join":   join,
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

//==========================================================================
// template custom functions

// Get input less 1
func minus1(n int) int { return n - 1 }

// Custom template function to join string array
func join(s []string) string { return strings.Join(s, ", ") }
