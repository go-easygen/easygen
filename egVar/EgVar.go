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
// https://github.com/danverbraganza/varcaser/
/*
ls - LowerSnakeCase: lower_snake_case
ss - ScreamingSnakeCase: SCREAMING_SNAKE_CASE
lk - KebabCase: kebab-case
ScreamingKebabCase: SCREAMING-KEBAB-CASE
hh - HttpHeaderCase: HTTP-Header-Case
UpperCamelCase: UpperCamelCase (renders HTTP as Http)
LowerCamelCase: lowerCamelCase (renders HTTP as Http)
uc - UpperCamelCaseKeepCaps: UpperCamelCaseKeepCaps (renders HTTP as HTTP)
lc - LowerCamelCaseKeepCaps: lowerCamelCaseKeepCaps (renders HTTP as HTTP)
*/
var (
	cls2lc = varcaser.Caser{
		From: varcaser.LowerSnakeCase, To: varcaser.LowerCamelCaseKeepCaps}
	cls2uc = varcaser.Caser{
		From: varcaser.LowerSnakeCase, To: varcaser.UpperCamelCaseKeepCaps}
	cls2ss = varcaser.Caser{
		From: varcaser.LowerSnakeCase, To: varcaser.ScreamingSnakeCase}
	cls2lk = varcaser.Caser{
		From: varcaser.LowerSnakeCase, To: varcaser.KebabCase}
	cls2hh = varcaser.Caser{
		From: varcaser.LowerSnakeCase, To: varcaser.HttpHeaderCase}

	css2lc = varcaser.Caser{
		From: varcaser.ScreamingSnakeCase, To: varcaser.LowerCamelCaseKeepCaps}
	css2uc = varcaser.Caser{
		From: varcaser.ScreamingSnakeCase, To: varcaser.UpperCamelCaseKeepCaps}
	css2ss = varcaser.Caser{
		From: varcaser.ScreamingSnakeCase, To: varcaser.ScreamingSnakeCase}
	css2lk = varcaser.Caser{
		From: varcaser.ScreamingSnakeCase, To: varcaser.KebabCase}
	css2hh = varcaser.Caser{
		From: varcaser.ScreamingSnakeCase, To: varcaser.HttpHeaderCase}

	clk2lc = varcaser.Caser{From: varcaser.KebabCase, To: varcaser.LowerCamelCaseKeepCaps}
	clk2uc = varcaser.Caser{From: varcaser.KebabCase, To: varcaser.UpperCamelCaseKeepCaps}
	clk2ls = varcaser.Caser{From: varcaser.KebabCase, To: varcaser.LowerSnakeCase}
	clk2ss = varcaser.Caser{
		From: varcaser.KebabCase, To: varcaser.ScreamingSnakeCase}

	clc2ss = varcaser.Caser{
		From: varcaser.LowerCamelCase, To: varcaser.ScreamingSnakeCase}
	cuc2ss = varcaser.Caser{
		From: varcaser.UpperCamelCase, To: varcaser.ScreamingSnakeCase}
)

var egFuncMap = easygen.FuncMap{
	"cls2lc": cls2lc.String,
	"cls2uc": cls2uc.String,
	"cls2ss": cls2ss.String,
	"cls2lk": cls2lk.String,
	"cls2hh": cls2hh.String,

	"css2lc": css2lc.String,
	"css2uc": css2uc.String,
	"css2ss": css2ss.String,
	"css2lk": css2lk.String,
	"css2hh": css2hh.String,

	"clk2lc": clk2lc.String,
	"clk2uc": clk2uc.String,
	"clk2ls": clk2ls.String,
	"clk2ss": clk2ss.String,
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
