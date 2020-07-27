////////////////////////////////////////////////////////////////////////////
// Package: egFilePath
// Purpose: EasyGen Generic FilePath Functionalities
// Authors: Tong Sun (c) 2019, All rights reserved
////////////////////////////////////////////////////////////////////////////

/*

Package egFilePath provides filepath manupilation functionalities.

egFilePath provides filepath manupilation manipulation, provided by "path/filepath".

*/
package egFilePath

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/go-easygen/easygen"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// EgFilePath -- EasyGen FilePath manupilation functionalities
/*

  fpAbs is wrapper for filepath.Abs
  fpBase is wrapper for filepath.Base
  fpClean is wrapper for filepath.Clean
  fpDir is wrapper for filepath.Dir
  fpEvalSymlinks is wrapper for filepath.EvalSymlinks
  fpExt is wrapper for filepath.Ext
  fpFromSlash is wrapper for filepath.FromSlash
  fpGlob is wrapper for filepath.Glob
  fpHasPrefix is wrapper for filepath.HasPrefix
  fpIsAbs is wrapper for filepath.IsAbs
  fpJoin is wrapper for filepath.Join
  fpMatch is wrapper for filepath.Match
  fpRel is wrapper for filepath.Rel
  fpSplitList is wrapper for filepath.SplitList
  fpToSlash is wrapper for filepath.ToSlash
  fpVolumeName is wrapper for filepath.VolumeName

  isDir is template function for IsDir
  basename is template function for Basename

*/
type EgFilePath struct {
	*easygen.EgBase
}

var egFuncMap = easygen.FuncMap{
	"fpAbs":          filepath.Abs,
	"fpBase":         filepath.Base,
	"fpClean":        filepath.Clean,
	"fpDir":          filepath.Dir,
	"fpEvalSymlinks": filepath.EvalSymlinks,
	"fpExt":          filepath.Ext,
	"fpFromSlash":    filepath.FromSlash,
	"fpGlob":         filepath.Glob,
	"fpHasPrefix":    filepath.HasPrefix,
	"fpIsAbs":        filepath.IsAbs,
	"fpJoin":         filepath.Join,
	"fpMatch":        filepath.Match,
	"fpRel":          filepath.Rel,
	"fpSplitList":    filepath.SplitList,
	"fpToSlash":      filepath.ToSlash,
	"fpVolumeName":   filepath.VolumeName,
	"isDir":          IsDir,
	"basename":       Basename,
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

// FuncDefs returns the custom definition mapping for this specific package.
func FuncDefs() template.FuncMap {
	return template.FuncMap(egFuncMap)
}

//==========================================================================
// support functions

// IsDir returns true if path is dir
func IsDir(path string) bool {
	info, _ := os.Stat(path)
	return info.IsDir()
}

// Basename returns basename(path)
func Basename(s string) string {
	s = filepath.Base(s)
	n := strings.LastIndexByte(s, '.')
	if n > 0 {
		return s[:n]
	}
	return s
}
