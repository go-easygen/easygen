////////////////////////////////////////////////////////////////////////////
// Porgram: EasyGen
// Purpose: Universal code/text generator that is easy to use
// Authors: Tong Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////
// Program start

package main

import (
	"flag"
	"fmt"
	//ht "html/template"
	"io/ioutil"
	"os"
	"strings"
	tt "text/template"

	"gopkg.in/yaml.v2"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const progname = "EasyGen" // os.Args[0]

type Options struct {
	Html bool
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var opts Options

////////////////////////////////////////////////////////////////////////////
// Commandline definitions

func init() {
	flag.BoolVar(&opts.Html, "html", false, "Use html template instead of text")
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

	source, err := ioutil.ReadFile(fileName + ".yaml")
	checkError(err)

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal(source, &m)
	checkError(err)

	// var t interface{}
	// if opts.Html {
	// 	t := ht.New("H")
	// } else {
	// 	t := tt.New("T")
	// }

	t := tt.New("T")
	// add functions
	t = t.Funcs(tt.FuncMap{"join": Join})
	//fmt.Fprintf(os.Stderr, "] Ok\n")
	t, err = t.Parse("{{range .Colors}}{{.}}, {{end}}.")

	t, err = tt.ParseFiles(fileName + ".tmpl")
	checkError(err)

	err = t.Execute(os.Stdout, m)
	checkError(err)

}

////////////////////////////////////////////////////////////////////////////
// Function definitions

func checkError(err error) {
	if err != nil {
		fmt.Printf("[%s] Fatal error - %v", progname, err.Error())
		os.Exit(1)
	}
}

//==========================================================================
// template custom functions

var funcs = tt.FuncMap{"minus1": func(n int) int { return n - 1 }}

func Join(s []string) string { return strings.Join(s, ", ") }
