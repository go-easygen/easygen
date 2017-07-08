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

// EgtBase -- EasyGen Template - Consul Template
type EgtCT struct {
	*easygen.EgtBase
}

var egFuncMap = easygen.FuncMap{
	"add":      add,
	"subtract": subtract,
	"multiply": multiply,
	"divide":   divide,
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

func FuncDefs() template.FuncMap {
	return template.FuncMap(egFuncMap)
}
