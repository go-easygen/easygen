package easygen

import (
	"fmt"
	"strconv"
	"time"
)

// now is function that represents the current time in UTC. This is here
// primarily for the tests to override times.
var now = func() time.Time { return time.Now() }

//==========================================================================
// template function for date & time

// https://godoc.org/time#pkg-constants

// date returns the date or year
func date(fmt string) string {
	switch fmt {
	case "Y4":
		// returns the year string of length 4
		return time.Now().Format("2006")
	case "I":
		// returns the output of `date -I`
		return time.Now().Format("2006-01-02")
		// for the output of `date -Iseconds`
		// use `timestamp`
	default:
		return time.Now().Format("2006" + fmt + "01" + fmt + "02")
	}
}

// https://github.com/hashicorp/consul-template/blob/de2ebf4/template_functions.go#L666-L682

// timestamp returns the current UNIX timestamp in UTC. If an argument is
// specified, it will be used to format the timestamp.
func timestamp(s ...string) (string, error) {
	switch len(s) {
	case 0:
		return now().Format(time.RFC3339), nil
	case 1:
		if s[0] == "unix" {
			return strconv.FormatInt(now().Unix(), 10), nil
		}
		return now().Format(s[0]), nil
	default:
		return "", fmt.Errorf("timestamp: wrong number of arguments, expected 0 or 1"+
			", but got %d", len(s))
	}
}
