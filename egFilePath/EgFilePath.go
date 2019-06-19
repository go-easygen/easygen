////////////////////////////////////////////////////////////////////////////
// Package: egFilePath
// Purpose: EasyGen Generic FilePath Functionalities
// Authors: Tong Sun (c) 2019, All rights reserved
////////////////////////////////////////////////////////////////////////////

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

// EgFilePath -- EasyGen Calculation
type EgFilePath struct {
	*easygen.EgBase
}

var egFuncMap = easygen.FuncMap{
	//"fpAbs":          filepath.Abs,
	"fpBase":  filepath.Base,
	"fpClean": filepath.Clean,
	"fpDir":   filepath.Dir,
	//"fpEvalSymlinks": filepath.EvalSymlinks,
	"fpExt":       filepath.Ext,
	"fpFromSlash": filepath.FromSlash,
	//"fpGlob":         filepath.Glob,
	"fpHasPrefix": filepath.HasPrefix,
	"fpIsAbs":     filepath.IsAbs,
	"fpJoin":      filepath.Join,
	//"fpMatch":     filepath.Match,
	"fpRel": filepath.Rel,
	//"fpSplit":      filepath.Split,
	//"fpSplitList":  filepath.SplitList,
	"fpToSlash":    filepath.ToSlash,
	"fpVolumeName": filepath.VolumeName,
	"isDir":        IsDir,
	"basename":     Basename,
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

// FuncDefs returns the custom definition mapping for this specific package.
func FuncDefs() template.FuncMap {
	return template.FuncMap(egFuncMap)
}

func IsDir(path string) bool {
	info, _ := os.Stat(path)
	return info.IsDir()
}

func Basename(s string) string {
	s = filepath.Base(s)
	n := strings.LastIndexByte(s, '.')
	if n > 0 {
		return s[:n]
	}
	return s
}
