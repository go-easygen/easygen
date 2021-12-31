////////////////////////////////////////////////////////////////////////////
// Package: easygen
// Purpose: Easy to use universal code/text generator
// Authors: Tong Sun (c) 2015-2021, All rights reserved
////////////////////////////////////////////////////////////////////////////

package easygen

import (
	"fmt"
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
/*

  stringsCompare is wrapper for strings.Compare
  stringsContains is wrapper for strings.Contains
  stringsContainsAny is wrapper for strings.ContainsAny
  stringsContainsRune is wrapper for strings.ContainsRune
  stringsCount is wrapper for strings.Count
  stringsEqualFold is wrapper for strings.EqualFold
  stringsFields is wrapper for strings.Fields
  stringsFieldsFunc is wrapper for strings.FieldsFunc
  stringsHasPrefix is wrapper for strings.HasPrefix
  stringsHasSuffix is wrapper for strings.HasSuffix
  stringsIndex is wrapper for strings.Index
  stringsIndexAny is wrapper for strings.IndexAny
  stringsIndexByte is wrapper for strings.IndexByte
  stringsIndexFunc is wrapper for strings.IndexFunc
  stringsIndexRune is wrapper for strings.IndexRune
  stringsJoin is wrapper for strings.Join
  stringsLastIndex is wrapper for strings.LastIndex
  stringsLastIndexAny is wrapper for strings.LastIndexAny
  stringsLastIndexByte is wrapper for strings.LastIndexByte
  stringsLastIndexFunc is wrapper for strings.LastIndexFunc
  stringsMap is wrapper for strings.Map
  stringsRepeat is wrapper for strings.Repeat
  stringsReplace is wrapper for strings.Replace
  stringsSplit is wrapper for strings.Split
  stringsSplitAfter is wrapper for strings.SplitAfter
  stringsSplitAfterN is wrapper for strings.SplitAfterN
  stringsSplitN is wrapper for strings.SplitN
  stringsTitle is wrapper for strings.Title
  stringsToLower is wrapper for strings.ToLower
  stringsToLowerSpecial is wrapper for strings.ToLowerSpecial
  stringsToTitle is wrapper for strings.ToTitle
  stringsToTitleSpecial is wrapper for strings.ToTitleSpecial
  stringsToUpper is wrapper for strings.ToUpper
  stringsToUpperSpecial is wrapper for strings.ToUpperSpecial
  stringsTrim is wrapper for strings.Trim
  stringsTrimFunc is wrapper for strings.TrimFunc
  stringsTrimLeft is wrapper for strings.TrimLeft
  stringsTrimLeftFunc is wrapper for strings.TrimLeftFunc
  stringsTrimPrefix is wrapper for strings.TrimPrefix
  stringsTrimRight is wrapper for strings.TrimRight
  stringsTrimRightFunc is wrapper for strings.TrimRightFunc
  stringsTrimSpace is wrapper for strings.TrimSpace
  stringsTrimSuffix is wrapper for strings.TrimSuffix

  eqf is wrapper for strings.EqualFold
  split is wrapper for strings.Fields
  sprintf is wrapper for fmt.Sprintf

  regexpFindAllString is template function for RegexpFindAllString
  regexpFindAllStringIndex is template function for RegexpFindAllStringIndex
  regexpFindAllStringSubmatch is template function for RegexpFindAllStringSubmatch
  regexpFindAllStringSubmatchIndex is template function for RegexpFindAllStringSubmatchIndex
  regexpFindString is template function for RegexpFindString
  regexpFindStringIndex is template function for RegexpFindStringIndex
  regexpFindStringSubmatch is template function for RegexpFindStringSubmatch
  regexpFindStringSubmatchIndex is template function for RegexpFindStringSubmatchIndex
  regexpMatchString is template function for RegexpMatchString
  regexpReplaceAllLiteralString is template function for RegexpReplaceAllLiteralString
  regexpReplaceAllString is template function for RegexpReplaceAllString
  regexpReplaceAllStringFunc is template function for RegexpReplaceAllStringFunc
  regexpSplit is template function for RegexpSplit

  ENV is template function for os.Getenv
  substr is template function for Substr
  coalesce is template function for Coalesce
  quote4shell is template function for Quote4shell

  minus1 is template function for Minus1
  date is template function for Date
  timestamp is template function for Timestamp

*/
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
	"sprintf": fmt.Sprintf,

	// == standard regexp function definitions
	"regexpFindAllString":              RegexpFindAllString,
	"regexpFindAllStringIndex":         RegexpFindAllStringIndex,
	"regexpFindAllStringSubmatch":      RegexpFindAllStringSubmatch,
	"regexpFindAllStringSubmatchIndex": RegexpFindAllStringSubmatchIndex,
	"regexpFindString":                 RegexpFindString,
	"regexpFindStringIndex":            RegexpFindStringIndex,
	"regexpFindStringSubmatch":         RegexpFindStringSubmatch,
	"regexpFindStringSubmatchIndex":    RegexpFindStringSubmatchIndex,
	"regexpMatchString":                RegexpMatchString,
	"regexpReplaceAllLiteralString":    RegexpReplaceAllLiteralString,
	"regexpReplaceAllString":           RegexpReplaceAllString,
	"regexpReplaceAllStringFunc":       RegexpReplaceAllStringFunc,
	"regexpSplit":                      RegexpSplit,

	// == my added functions
	"ENV":         os.Getenv,
	"substr":      Substr,
	"coalesce":    Coalesce,
	"quote4shell": Quote4shell,

	"iterate":   Iterate,
	"argsa":     ArgsA,
	"argsm":     ArgsM,
	"minus1":    Minus1,
	"date":      Date,
	"timestamp": Timestamp,
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
