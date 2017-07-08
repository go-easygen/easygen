////////////////////////////////////////////////////////////////////////////
// Package: easygen
// Purpose: Easy to use universal code/text generator
// Authors: Tong Sun (c) 2015-17, All rights reserved
////////////////////////////////////////////////////////////////////////////

package easygen

import (
	"io"
	"strings"
	"text/template"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

type Template interface {
	Execute(wr io.Writer, data interface{}) error
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	ParseFiles(filenames ...string) (*template.Template, error)
	Name() string
}

// EgBase -- EasyGen Template Base
type EgBase struct {
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

func NewTemplate() *EgBase {
	return &EgBase{template.New("EgBase")}
}

func (t *EgBase) Customize() *EgBase {
	return t
}
