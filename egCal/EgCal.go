////////////////////////////////////////////////////////////////////////////
// Package: egCal
// Purpose: easygen generic calculation functionalities
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

/*

Package egCal provides generic calculation functionalities.

egCal provides the generic calculation functionalities from

consul template functions
https://github.com/hashicorp/consul-template / template_functions.go

*/
package egCal

import (
	"text/template"

	"github.com/go-easygen/easygen"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// EgCal -- EasyGen Calculation
/*

  add is template function for Add
  subtract is template function for Subtract
  multiply is template function for Multiply
  divide is template function for Divide

*/
type EgCal struct {
	*easygen.EgBase
}

var egFuncMap = easygen.FuncMap{
	"add":      Add,
	"subtract": Subtract,
	"multiply": Multiply,
	"divide":   Divide,
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

// FuncDefs returns the custom definition mapping for this specific package.
func FuncDefs() template.FuncMap {
	return template.FuncMap(egFuncMap)
}
