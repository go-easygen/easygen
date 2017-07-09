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

// The Template defines the common ground for both text and html Template
type Template interface {
	Execute(wr io.Writer, data interface{}) error
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	Parse(text string) (*template.Template, error)
	ParseFiles(filenames ...string) (*template.Template, error)
	Name() string
}

// EgBase -- EasyGen Template Base
type EgBase struct {
	*template.Template
}

// The FuncMap defined in easygen will shield the dependency of either
// text or html template, giving an implementation agnostic abstraction
// that will works for both cases.
type FuncMap map[string]interface{}

var egFuncMap = FuncMap{
	"eqf":      strings.EqualFold,
	"split":    strings.Fields,
	"replace":  replace,
	"replacec": replacec,

	"minus1": minus1,
	"dateI":  dateI,
	"year4":  year4,
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

// FuncDefs returns the custom definition mapping for this specific package.
func FuncDefs() template.FuncMap {
	return template.FuncMap(egFuncMap)
}

// NewTemplate returns a new Template for this specific package.
func NewTemplate() *EgBase {
	return &EgBase{template.New("EgBase")}
}

// Customize allows customization for this specific package.
func (t *EgBase) Customize() *EgBase {
	return t
}
