////////////////////////////////////////////////////////////////////////////
// Package: easygen
// Purpose: Easy to use universal code/text generator
// Authors: Tong Sun (c) 2015-17, All rights reserved
////////////////////////////////////////////////////////////////////////////

package easygen

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	yaml "gopkg.in/yaml.v2"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

type Template interface {
	Execute(wr io.Writer, data interface{}) error
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	ParseFiles(filenames ...string) (*template.Template, error)
	Name() string
}

// EgtBase -- EasyGen Template Base
type EgtBase struct {
	*template.Template
}

type FuncMap map[string]interface{}

var egFuncMap = FuncMap{
	"eqf":      strings.EqualFold,
	"split":    strings.Fields,
	"replace":  replace,
	"replacec": replacec,

	"minus1": minus1,
	"dateI":  dateI,
	"year4":  year4,
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

////////////////////////////////////////////////////////////////////////////
// Function definitions

func FuncDefs() template.FuncMap {
	return template.FuncMap(egFuncMap)
}

func NewTemplate() *EgtBase {
	return &EgtBase{template.New("EgtBase")}
}

func (t *EgtBase) Customize() *EgtBase {
	return t
}

func Process(t Template, fileName string) string {
	return Process2(t, fileName, fileName)
}

func Process2(t Template, fileNameTempl string, fileNames ...string) string {
	fileName := fileNames[0]

	source, err := ioutil.ReadFile(fileName + Opts.ExtYaml)
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

	// template file name
	fileNameT := fileNameTempl
	if len(Opts.TemplateFile) > 0 {
		fileNameT = Opts.TemplateFile
	}
	fileNameT = fileNameT + Opts.ExtTmpl
	tn, err := t.ParseFiles(fileNameT)
	checkError(err)

	buf := new(bytes.Buffer)
	err = tn.ExecuteTemplate(buf, filepath.Base(fileNameT), m)
	checkError(err)

	return buf.String()
}
