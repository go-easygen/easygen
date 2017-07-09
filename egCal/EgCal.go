////////////////////////////////////////////////////////////////////////////
// Package: egCal
// Purpose: easygen generic calculation functionalities
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package egCal

import (
	"text/template"

	"github.com/go-easygen/easygen"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// EgCal -- EasyGen Calculation
type EgCal struct {
	*easygen.EgBase
}

var egFuncMap = easygen.FuncMap{
	"add":      add,
	"subtract": subtract,
	"multiply": multiply,
	"divide":   divide,
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

// FuncDefs returns the custom definition mapping for this specific package.
func FuncDefs() template.FuncMap {
	return template.FuncMap(egFuncMap)
}
