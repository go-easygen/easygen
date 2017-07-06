////////////////////////////////////////////////////////////////////////////
// Package: easygen
// Purpose: Easy to use universal code/text generator
// Authors: Tong Sun (c) 2015-17, All rights reserved
////////////////////////////////////////////////////////////////////////////

package easygen

import "io"

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// common type for a *(text|html).Template value
type Template interface {
	Execute(wr io.Writer, data interface{}) error
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	Name() string
}
