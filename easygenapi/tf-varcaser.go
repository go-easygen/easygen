package easygenapi

import (
	"github.com/danverbraganza/varcaser/varcaser"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// pre-configed varcaser caser converters
// the names are self-explanatory from their definitions
var (
	cls2lc = varcaser.Caser{From: varcaser.LowerSnakeCase, To: varcaser.LowerCamelCase}
	cls2uc = varcaser.Caser{From: varcaser.LowerSnakeCase, To: varcaser.UpperCamelCase}
	cls2ss = varcaser.Caser{From: varcaser.LowerSnakeCase, To: varcaser.ScreamingSnakeCase}
	ck2lc  = varcaser.Caser{From: varcaser.KebabCase, To: varcaser.LowerCamelCase}
	ck2uc  = varcaser.Caser{From: varcaser.KebabCase, To: varcaser.UpperCamelCase}
	ck2ss  = varcaser.Caser{From: varcaser.KebabCase, To: varcaser.ScreamingSnakeCase}
)

//==========================================================================
// template functions
