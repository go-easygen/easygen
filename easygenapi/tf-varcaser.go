package easygenapi

import (
	"github.com/danverbraganza/varcaser/varcaser"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// pre-config some varcaser transformers
var (
	Ck2lc = varcaser.Caser{From: varcaser.KebabCase, To: varcaser.LowerCamelCase}
	Ck2uc = varcaser.Caser{From: varcaser.KebabCase, To: varcaser.UpperCamelCase}
)
