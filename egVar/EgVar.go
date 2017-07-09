////////////////////////////////////////////////////////////////////////////
// Package: egVar
// Purpose: easygen variable naming functionalities
// Authors: Tong Sun (c) 2016-17, All rights reserved
////////////////////////////////////////////////////////////////////////////

/*

Package egVar provides variable naming functionalities.

egVar provides variable naming manipulation, available from danverbraganza/varcaser.

*/
package egVar

import (
	"text/template"

	"github.com/danverbraganza/varcaser/varcaser"
	"github.com/go-easygen/easygen"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// EgVar -- EasyGen variable naming
type EgVar struct {
	*easygen.EgBase
}

// pre-configed varcaser caser converters
// the names are self-explanatory from their definitions
var (
	cls2lc = varcaser.Caser{From: varcaser.LowerSnakeCase, To: varcaser.LowerCamelCase}
	cls2uc = varcaser.Caser{From: varcaser.LowerSnakeCase, To: varcaser.UpperCamelCase}
	cls2ss = varcaser.Caser{From: varcaser.LowerSnakeCase, To: varcaser.ScreamingSnakeCase}

	ck2lc = varcaser.Caser{From: varcaser.KebabCase, To: varcaser.LowerCamelCase}
	ck2uc = varcaser.Caser{From: varcaser.KebabCase, To: varcaser.UpperCamelCase}
	ck2ls = varcaser.Caser{From: varcaser.KebabCase, To: varcaser.LowerSnakeCase}
	ck2ss = varcaser.Caser{From: varcaser.KebabCase, To: varcaser.ScreamingSnakeCase}

	clc2ss = varcaser.Caser{From: varcaser.LowerCamelCase, To: varcaser.ScreamingSnakeCase}
	cuc2ss = varcaser.Caser{From: varcaser.UpperCamelCase, To: varcaser.ScreamingSnakeCase}
)

var egFuncMap = easygen.FuncMap{
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

// FuncDefs returns the custom definition mapping for this specific package.
func FuncDefs() template.FuncMap {
	return template.FuncMap(egFuncMap)
}

// NewTemplate returns a new Template for this specific package.
func NewTemplate() *EgVar {
	return &EgVar{&easygen.EgBase{template.New("EgVar")}}
}
