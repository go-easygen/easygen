package easygen

import "time"

//==========================================================================
// template function for date & time

// https://godoc.org/time#pkg-constants

// dateI returns the output of `date -I`
func dateI() string { return time.Now().Format("2006-01-02") }
