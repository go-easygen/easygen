////////////////////////////////////////////////////////////////////////////
// Package: easygen
// Purpose: Easy to use universal code/text generator
// Authors: Tong Sun (c) 2015-17, All rights reserved
////////////////////////////////////////////////////////////////////////////

package easygen

import (
	"io"
	"os"
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
	// == standard strings function definitions
	"stringsCompare":        strings.Compare,
	"stringsContains":       strings.Contains,
	"stringsContainsAny":    strings.ContainsAny,
	"stringsContainsRune":   strings.ContainsRune,
	"stringsCount":          strings.Count,
	"stringsEqualFold":      strings.EqualFold,
	"stringsFields":         strings.Fields,
	"stringsFieldsFunc":     strings.FieldsFunc,
	"stringsHasPrefix":      strings.HasPrefix,
	"stringsHasSuffix":      strings.HasSuffix,
	"stringsIndex":          strings.Index,
	"stringsIndexAny":       strings.IndexAny,
	"stringsIndexByte":      strings.IndexByte,
	"stringsIndexFunc":      strings.IndexFunc,
	"stringsIndexRune":      strings.IndexRune,
	"stringsJoin":           strings.Join,
	"stringsLastIndex":      strings.LastIndex,
	"stringsLastIndexAny":   strings.LastIndexAny,
	"stringsLastIndexByte":  strings.LastIndexByte,
	"stringsLastIndexFunc":  strings.LastIndexFunc,
	"stringsMap":            strings.Map,
	"stringsRepeat":         strings.Repeat,
	"stringsReplace":        strings.Replace,
	"stringsSplit":          strings.Split,
	"stringsSplitAfter":     strings.SplitAfter,
	"stringsSplitAfterN":    strings.SplitAfterN,
	"stringsSplitN":         strings.SplitN,
	"stringsTitle":          strings.Title,
	"stringsToLower":        strings.ToLower,
	"stringsToLowerSpecial": strings.ToLowerSpecial,
	"stringsToTitle":        strings.ToTitle,
	"stringsToTitleSpecial": strings.ToTitleSpecial,
	"stringsToUpper":        strings.ToUpper,
	"stringsToUpperSpecial": strings.ToUpperSpecial,
	"stringsTrim":           strings.Trim,
	"stringsTrimFunc":       strings.TrimFunc,
	"stringsTrimLeft":       strings.TrimLeft,
	"stringsTrimLeftFunc":   strings.TrimLeftFunc,
	"stringsTrimPrefix":     strings.TrimPrefix,
	"stringsTrimRight":      strings.TrimRight,
	"stringsTrimRightFunc":  strings.TrimRightFunc,
	"stringsTrimSpace":      strings.TrimSpace,
	"stringsTrimSuffix":     strings.TrimSuffix,
	// aliases
	"eqf":   strings.EqualFold,
	"split": strings.Fields,

	// == standard regexp function definitions
	"regexpFindAllString":              regexpFindAllString,
	"regexpFindAllStringIndex":         regexpFindAllStringIndex,
	"regexpFindAllStringSubmatch":      regexpFindAllStringSubmatch,
	"regexpFindAllStringSubmatchIndex": regexpFindAllStringSubmatchIndex,
	"regexpFindString":                 regexpFindString,
	"regexpFindStringIndex":            regexpFindStringIndex,
	"regexpFindStringSubmatch":         regexpFindStringSubmatch,
	"regexpFindStringSubmatchIndex":    regexpFindStringSubmatchIndex,
	"regexpMatchString":                regexpMatchString,
	"regexpReplaceAllLiteralString":    regexpReplaceAllLiteralString,
	"regexpReplaceAllString":           regexpReplaceAllString,
	"regexpReplaceAllStringFunc":       regexpReplaceAllStringFunc,
	"regexpSplit":                      regexpSplit,

	// == my added functions
	"ENV":         os.Getenv,
	"coalesce":    coalesce,
	"quote4shell": quote4shell,

	"minus1":    minus1,
	"date":      date,
	"timestamp": timestamp,
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
